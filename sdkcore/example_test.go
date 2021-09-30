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

func Example() {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	// Create client
	c, err := sdkcore.NewClient(client.Options{Logger: logLevelWarn})
	if err != nil {
		log.Fatalf("Failed creating client: %v", err)
	}
	defer c.Close()

	// Start worker in background
	const taskQueue = "greeting-task-queue"
	w := sdkcore.NewWorker(c, taskQueue, worker.Options{})
	w.RegisterWorkflow(GreetingWorkflow)
	w.RegisterActivity(ComposeGreeting)
	if err := w.Start(); err != nil {
		log.Fatalf("Failed starting worker: %v", err)
	}
	defer w.Stop()

	// Run workflow
	wrk, err := c.ExecuteWorkflow(
		ctx,
		client.StartWorkflowOptions{ID: "greeting-workflow", TaskQueue: taskQueue},
		GreetingWorkflow,
		"World",
	)
	if err != nil {
		log.Fatalf("Failed starting workflow: %v", err)
	}
	var greeting string
	if err = wrk.Get(ctx, &greeting); err != nil {
		log.Fatalf("Failed getting workflow result: %v", err)
	}
	fmt.Printf("Result: %v", greeting)
	// Output: Result: Hello, World!
}

func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{StartToCloseTimeout: time.Second * 5})
	var result string
	err := workflow.ExecuteActivity(ctx, ComposeGreeting, name).Get(ctx, &result)
	return result, err
}

func ComposeGreeting(name string) (string, error) {
	greeting := fmt.Sprintf("Hello, %s!", name)
	return greeting, nil
}

type logLevel int

const (
	logLevelNone logLevel = iota
	logLevelError
	logLevelWarn
	logLevelInfo
	logLevelDebug
	logLevelAll
)

func (l logLevel) log(level logLevel, levelStr string, msg string, keyVals ...interface{}) {
	if l >= level {
		log.Println(append([]interface{}{levelStr, msg}, keyVals...))
	}
}

func (l logLevel) Debug(msg string, keyVals ...interface{}) {
	l.log(logLevelDebug, "DEBUG", msg, keyVals...)
}
func (l logLevel) Info(msg string, keyVals ...interface{}) {
	l.log(logLevelInfo, "INFO", msg, keyVals...)
}
func (l logLevel) Warn(msg string, keyVals ...interface{}) {
	l.log(logLevelWarn, "WARN", msg, keyVals...)
}
func (l logLevel) Error(msg string, keyVals ...interface{}) {
	l.log(logLevelError, "ERROR", msg, keyVals...)
}
