package sdkcore

import (
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

type coreWorker struct{}

func NewWorker(client client.Client, taskQueue string, options worker.Options) worker.Worker {
	// TODO
	return &coreWorker{}
}

func (*coreWorker) Start() error {
	panic("TODO")
}

func (*coreWorker) Run(interruptCh <-chan interface{}) error {
	panic("TODO")
}

func (*coreWorker) Stop() {
	panic("TODO")
}

func (*coreWorker) RegisterWorkflow(w interface{}) {
	panic("TODO")
}

func (*coreWorker) RegisterWorkflowWithOptions(w interface{}, options workflow.RegisterOptions) {
	panic("TODO")
}

func (*coreWorker) RegisterActivity(a interface{}) {
	panic("TODO")
}

func (*coreWorker) RegisterActivityWithOptions(a interface{}, options activity.RegisterOptions) {
	panic("TODO")
}
