package sdkcore_test

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/sdkcore"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

// Use to flip between core and not
const useSDKCore = true

func newClientAndWorker(
	clientOpts client.Options,
	workerOpts worker.Options,
	taskQueue string,
) (client.Client, worker.Worker, error) {
	newClient, newWorker := client.NewClient, worker.New
	if useSDKCore {
		newClient, newWorker = sdkcore.NewClient, sdkcore.NewWorker
	}
	c, err := newClient(clientOpts)
	if err != nil {
		return nil, nil, fmt.Errorf("failed creating client: %w", err)
	}
	return c, newWorker(c, taskQueue, workerOpts), nil
}

func Example() {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	// Create client and worker
	const taskQueue = "greeting-task-queue"
	c, w, err := newClientAndWorker(client.Options{}, worker.Options{}, taskQueue)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	log.Printf("Created client and worker")

	// Register workflow/activity and start worker in background
	w.RegisterWorkflow(GreetingWorkflow)
	w.RegisterActivity(ComposeGreeting)
	if err := w.Start(); err != nil {
		log.Fatalf("Failed starting worker: %v", err)
	}
	defer w.Stop()
	log.Printf("Started worker")

	// Run workflow
	wrk, err := c.ExecuteWorkflow(
		ctx,
		client.StartWorkflowOptions{TaskQueue: taskQueue},
		GreetingWorkflow,
		"World",
	)
	if err != nil {
		log.Fatalf("Failed starting workflow: %v", err)
	}
	log.Printf("Started workflow, %v - %v", wrk.GetID(), wrk.GetRunID())
	var greeting string
	if err = wrk.Get(ctx, &greeting); err != nil {
		log.Fatalf("Failed getting workflow result: %v", err)
	}
	fmt.Printf("Result: %v", greeting)
	// Output: Result: Hello, World!
}

func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	workflow.GetLogger(ctx).Debug("Started greeting workflow")
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{StartToCloseTimeout: time.Second * 5})
	var result string
	err := workflow.ExecuteActivity(ctx, ComposeGreeting, name).Get(ctx, &result)
	return result, err
}

func ComposeGreeting(name string) (string, error) {
	greeting := fmt.Sprintf("Hello, %s!", name)
	return greeting, nil
}
