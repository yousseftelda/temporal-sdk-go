package sdkcore

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/interceptors"
	"go.temporal.io/sdk/internal"
	"go.temporal.io/sdk/sdkcore/internal/ffi"
	workflowactivationpb "go.temporal.io/sdk/sdkcore/internal/ffi/corepb/workflowactivationpb"
	"go.temporal.io/sdk/worker"
)

type coreWorker struct {
	client    *coreClient
	taskQueue string
	options   *worker.Options

	runCtx      context.Context
	runCancel   context.CancelFunc
	runDoneCh   chan struct{}
	runDoneErr  error
	runDoneLock sync.Mutex

	*internal.Registry

	runningWorkflows map[string]*workflowEnv
}

func NewWorker(client client.Client, taskQueue string, options worker.Options) worker.Worker {
	c, ok := client.(*coreClient)
	if !ok {
		panic("client must be created in this package")
	}
	cw := &coreWorker{
		client:    c,
		taskQueue: taskQueue,
		options:   &options,
		Registry:  internal.NewRegistry(),
	}
	// Set interceptor
	cw.Registry.SetWorkflowInterceptors([]interceptors.WorkflowInterceptor{cw})
	// Create context
	cw.runCtx, cw.runCancel = context.WithCancel(context.Background())
	return cw
}

func (c *coreWorker) Run(interruptCh <-chan interface{}) error {
	if err := c.Start(); err != nil {
		return err
	}
	c.runDoneLock.Lock()
	runDoneCh := c.runDoneCh
	c.runDoneLock.Unlock()
	select {
	case <-interruptCh:
		c.Stop()
	case <-runDoneCh:
	}
	c.runDoneLock.Lock()
	err := c.runDoneErr
	c.runDoneLock.Unlock()
	return err
}

// TODO(cretz): Document callers should use Run if they ever care what the error
// is (which they should)
func (c *coreWorker) Start() error {
	runDoneCh, err := c.registerWorker()
	if err != nil {
		return err
	}

	// Run in background, this is the entry point for all runs
	go func() {
		defer close(runDoneCh)
		if err := c.run(); err != nil {
			c.runDoneLock.Lock()
			c.runDoneErr = err
			c.runDoneLock.Unlock()
		}
	}()
	return nil
}

func (c *coreWorker) registerWorker() (runDoneCh chan struct{}, err error) {
	c.runDoneLock.Lock()
	defer c.runDoneLock.Unlock()
	if c.runDoneCh != nil {
		return nil, fmt.Errorf("already started")
	}
	// Register the worker
	// TODO(cretz): More options
	if err := c.client.core.RegisterWorker(&ffi.WorkerConfig{TaskQueue: c.taskQueue}); err != nil {
		return nil, fmt.Errorf("registering worker: %w", err)
	}
	c.runDoneCh = make(chan struct{})
	return c.runDoneCh, nil
}

func (c *coreWorker) Stop() {
	c.runCancel()
	c.runDoneLock.Lock()
	runDoneCh := c.runDoneCh
	c.runDoneLock.Unlock()
	if runDoneCh != nil {
		<-c.runDoneCh
	}
}

func (c *coreWorker) run() error {
	// Run both the workflow poller and activity poller and wait until done
	workflowErrCh := make(chan error, 1)
	go func() { workflowErrCh <- c.runWorkflowPoller() }()
	activityErrCh := make(chan error, 1)
	go func() { activityErrCh <- c.runActivityPoller() }()
	// Just fail on first error
	for workflowErrCh != nil || activityErrCh != nil {
		select {
		case <-c.runCtx.Done():
			return nil
		case err := <-workflowErrCh:
			if err != nil {
				return fmt.Errorf("polling workflow: %w", err)
			}
			workflowErrCh = nil
		case err := <-activityErrCh:
			if err != nil {
				return fmt.Errorf("polling activity: %w", err)
			}
			activityErrCh = nil
		}
	}
	return nil
}

func (c *coreWorker) runWorkflowPoller() error {
	for {
		act, err := c.client.core.PollWorkflowActivation(c.taskQueue)
		if err != nil {
			return err
		}
		c.client.options.Logger.Info("Got workflow activation", "Activation", act)
		// Keep track of all workflows that had things run so we can run the "tick"
		workflows := map[*workflowEnv]struct{}{}
		for _, job := range act.Jobs {
			var w *workflowEnv
			var err error
			switch job := job.Variant.(type) {
			case *workflowactivationpb.WFActivationJob_StartWorkflow:
				w, err = c.startWorkflow(act, job.StartWorkflow)
			default:
				err = fmt.Errorf("unrecognized job")
			}
			if err != nil {
				c.client.options.Logger.Error("Failed processing job", "Error", err, "Job", job)
			} else if w != nil {
				workflows[w] = struct{}{}
			}
		}
		// Run OnWorkflowTaskStarted on each affected workflow
		for workflow := range workflows {
			// TODO(cretz): Customizable deadlock timeout
			workflow.def.OnWorkflowTaskStarted(time.Minute)
		}
	}
}

func (c *coreWorker) runActivityPoller() error {
	for {
		task, err := c.client.core.PollActivityTask(c.taskQueue)
		if err != nil {
			return err
		}
		c.client.options.Logger.Info("Got activity task", "Task", task)
	}
}
