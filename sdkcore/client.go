package sdkcore

import (
	"context"
	"fmt"
	"reflect"
	"runtime"

	"github.com/pborman/uuid"
	enumspb "go.temporal.io/api/enums/v1"
	workflowservicepb "go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/converter"
	"go.temporal.io/sdk/sdkcore/internal/ffi"
)

// Global runtime
var rt *ffi.Runtime

func init() {
	rt = ffi.NewRuntime()
	runtime.SetFinalizer(rt, func(rt *ffi.Runtime) { rt.Close() })
}

type coreClient struct {
	core    *ffi.Core
	options *client.Options
}

func NewClient(options client.Options) (client.Client, error) {
	// TODO(cretz): More options
	if options.HostPort == "" {
		options.HostPort = client.DefaultHostPort
	}
	if options.Namespace == "" {
		options.Namespace = client.DefaultNamespace
	}
	if options.Logger == nil {
		options.Logger = logLevelAll
	}
	if options.DataConverter == nil {
		options.DataConverter = converter.GetDefaultDataConverter()
	}
	workerBinaryID, err := binaryChecksum()
	if err != nil {
		return nil, fmt.Errorf("failed building checksum: %w", err)
	}

	core, err := ffi.NewCore(rt, ffi.CoreConfig{
		// TODO(cretz): https?
		TargetURL:      "http://" + options.HostPort,
		Namespace:      options.Namespace,
		ClientName:     "temporal-go-core",
		ClientVersion:  "0.1.0",
		WorkerBinaryID: workerBinaryID,
	})
	if err != nil {
		return nil, fmt.Errorf("initializing core: %w", err)
	}
	return &coreClient{core: core, options: &options}, nil
}

func (c *coreClient) ExecuteWorkflow(
	ctx context.Context,
	options client.StartWorkflowOptions,
	workflow interface{},
	args ...interface{},
) (client.WorkflowRun, error) {
	// Default options
	// TODO(cretz): More options
	if options.ID == "" {
		options.ID = uuid.New()
	}
	if options.TaskQueue == "" {
		return nil, fmt.Errorf("missing task queue option")
	}

	// Workflow type
	// TODO(cretz): More function types
	if reflect.TypeOf(workflow).Kind() != reflect.Func {
		return nil, fmt.Errorf("only function types currently supported")
	}
	workflowType, isMethod := functionName(workflow)
	if isMethod {
		return nil, fmt.Errorf("methods not currently supported")
	}

	// Convert args
	// TODO(cretz): Different converters
	input, err := converter.GetDefaultDataConverter().ToPayloads(args...)
	if err != nil {
		return nil, fmt.Errorf("converting input: %w", err)
	}

	// Call
	resp, err := c.core.StartWorkflow(&ffi.StartWorkflowRequest{
		Input:        input,
		TaskQueue:    options.TaskQueue,
		WorkflowID:   options.ID,
		WorkflowType: workflowType,
		TaskTimeout:  options.WorkflowTaskTimeout,
	})
	if err != nil {
		return nil, fmt.Errorf("starting workflow: %w", err)
	}
	return &workflowRun{id: options.ID, runID: resp.RunId}, nil
}

type workflowRun struct {
	id    string
	runID string
}

func (w *workflowRun) GetID() string    { return w.id }
func (w *workflowRun) GetRunID() string { return w.runID }

func (w *workflowRun) Get(ctx context.Context, valuePtr interface{}) error {
	// TODO(cretz): Implement this
	<-ctx.Done()
	return ctx.Err()
}

func (*coreClient) GetWorkflow(ctx context.Context, workflowID string, runID string) client.WorkflowRun {
	panic("TODO")
}

func (*coreClient) SignalWorkflow(
	ctx context.Context,
	workflowID string,
	runID string,
	signalName string,
	arg interface{},
) error {
	panic("TODO")
}

func (*coreClient) SignalWithStartWorkflow(
	ctx context.Context,
	workflowID string,
	signalName string,
	signalArg interface{},
	options client.StartWorkflowOptions,
	workflow interface{},
	workflowArgs ...interface{},
) (client.WorkflowRun, error) {
	panic("TODO")
}

func (*coreClient) CancelWorkflow(ctx context.Context, workflowID string, runID string) error {
	panic("TODO")
}

func (*coreClient) TerminateWorkflow(
	ctx context.Context,
	workflowID string,
	runID string,
	reason string,
	details ...interface{},
) error {
	panic("TODO")
}

func (*coreClient) GetWorkflowHistory(
	ctx context.Context,
	workflowID string,
	runID string,
	isLongPoll bool,
	filterType enumspb.HistoryEventFilterType,
) client.HistoryEventIterator {
	panic("TODO")
}

func (*coreClient) CompleteActivity(ctx context.Context, taskToken []byte, result interface{}, err error) error {
	panic("TODO")
}

func (*coreClient) CompleteActivityByID(
	ctx context.Context,
	namespace string,
	workflowID string,
	runID string,
	activityID string,
	result interface{},
	err error,
) error {
	panic("TODO")
}

func (*coreClient) RecordActivityHeartbeat(ctx context.Context, taskToken []byte, details ...interface{}) error {
	panic("TODO")
}

func (*coreClient) RecordActivityHeartbeatByID(
	ctx context.Context,
	namespace string,
	workflowID string,
	runID string,
	activityID string,
	details ...interface{},
) error {
	panic("TODO")
}

func (*coreClient) ListClosedWorkflow(
	ctx context.Context,
	request *workflowservicepb.ListClosedWorkflowExecutionsRequest,
) (*workflowservicepb.ListClosedWorkflowExecutionsResponse, error) {
	panic("TODO")
}

func (*coreClient) ListOpenWorkflow(
	ctx context.Context,
	request *workflowservicepb.ListOpenWorkflowExecutionsRequest,
) (*workflowservicepb.ListOpenWorkflowExecutionsResponse, error) {
	panic("TODO")
}

func (*coreClient) ListWorkflow(
	ctx context.Context,
	request *workflowservicepb.ListWorkflowExecutionsRequest,
) (*workflowservicepb.ListWorkflowExecutionsResponse, error) {
	panic("TODO")
}

func (*coreClient) ListArchivedWorkflow(
	ctx context.Context,
	request *workflowservicepb.ListArchivedWorkflowExecutionsRequest,
) (*workflowservicepb.ListArchivedWorkflowExecutionsResponse, error) {
	panic("TODO")
}

func (*coreClient) ScanWorkflow(
	ctx context.Context,
	request *workflowservicepb.ScanWorkflowExecutionsRequest,
) (*workflowservicepb.ScanWorkflowExecutionsResponse, error) {
	panic("TODO")
}

func (*coreClient) CountWorkflow(
	ctx context.Context,
	request *workflowservicepb.CountWorkflowExecutionsRequest,
) (*workflowservicepb.CountWorkflowExecutionsResponse, error) {
	panic("TODO")
}

func (*coreClient) GetSearchAttributes(ctx context.Context) (*workflowservicepb.GetSearchAttributesResponse, error) {
	panic("TODO")
}

func (*coreClient) QueryWorkflow(
	ctx context.Context,
	workflowID string,
	runID string,
	queryType string,
	args ...interface{},
) (converter.EncodedValue, error) {
	panic("TODO")
}

func (*coreClient) QueryWorkflowWithOptions(
	ctx context.Context,
	request *client.QueryWorkflowWithOptionsRequest,
) (*client.QueryWorkflowWithOptionsResponse, error) {
	panic("TODO")
}

func (*coreClient) DescribeWorkflowExecution(
	ctx context.Context,
	workflowID string,
	runID string,
) (*workflowservicepb.DescribeWorkflowExecutionResponse, error) {
	panic("TODO")
}

func (*coreClient) DescribeTaskQueue(
	ctx context.Context,
	taskqueue string,
	taskqueueType enumspb.TaskQueueType,
) (*workflowservicepb.DescribeTaskQueueResponse, error) {
	panic("TODO")
}

func (*coreClient) ResetWorkflowExecution(
	ctx context.Context,
	request *workflowservicepb.ResetWorkflowExecutionRequest,
) (*workflowservicepb.ResetWorkflowExecutionResponse, error) {
	panic("TODO")
}

func (c *coreClient) Close() {
	if c.core != nil {
		c.core.Shutdown()
		// TODO(cretz): This will segfault stuff! Fix reference counting...
		c.core.Close()
		c.core = nil
	}
}
