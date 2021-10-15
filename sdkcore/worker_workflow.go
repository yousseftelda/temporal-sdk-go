package sdkcore

import (
	"github.com/gogo/protobuf/types"
	commonpb "go.temporal.io/api/common/v1"
	"go.temporal.io/sdk/converter"
	"go.temporal.io/sdk/interceptors"
	"go.temporal.io/sdk/internal"
	"go.temporal.io/sdk/log"
	workflowactivationpb "go.temporal.io/sdk/sdkcore/internal/ffi/corepb/workflowactivationpb"
	"go.temporal.io/sdk/workflow"
)

type workflowIntercept struct {
	// Intentionally nil interfaces at the moment
	interceptors.WorkflowInboundCallsInterceptor
	interceptors.WorkflowOutboundCallsInterceptor

	nextInbound  interceptors.WorkflowInboundCallsInterceptor
	nextOutbound interceptors.WorkflowOutboundCallsInterceptor
}

type workflowEnv struct {
	// Intentionally nil interface at the moment
	internal.WorkflowEnvironment
	def    internal.WorkflowDefinition
	info   *workflow.Info
	worker *coreWorker
}

func (c *coreWorker) InterceptWorkflow(
	info *workflow.Info,
	next interceptors.WorkflowInboundCallsInterceptor,
) interceptors.WorkflowInboundCallsInterceptor {
	return &workflowIntercept{nextInbound: next}
}

func (c *coreWorker) startWorkflow(
	act *workflowactivationpb.WFActivation,
	job *workflowactivationpb.StartWorkflow,
) (*workflowEnv, error) {
	// Get definition and execute
	def, err := c.GetWorkflowDefinition(internal.WorkflowType{Name: job.WorkflowType})
	if err != nil {
		return nil, err
	}
	// Convert headers and input
	header := &commonpb.Header{Fields: make(map[string]*commonpb.Payload, len(job.Headers))}
	for k, v := range job.Headers {
		header.Fields[k] = &commonpb.Payload{Metadata: v.Metadata, Data: v.Data}
	}
	input := &commonpb.Payloads{Payloads: make([]*commonpb.Payload, len(job.Arguments))}
	for i, arg := range job.Arguments {
		input.Payloads[i] = &commonpb.Payload{Metadata: arg.Metadata, Data: arg.Data}
	}

	// Build workflow
	w := &workflowEnv{
		def: def,
		info: &workflow.Info{
			WorkflowExecution: workflow.Execution{
				ID:    job.WorkflowId,
				RunID: act.RunId,
			},
			WorkflowType:  workflow.Type{Name: job.WorkflowType},
			TaskQueueName: c.taskQueue,
			// TODO(cretz):
			// WorkflowExecutionTimeout:
			// WorkflowRunTimeout:
			// WorkflowTaskTimeout:
			Namespace: c.client.options.Namespace,
			// TODO(cretz):
			// Attempt:
			// WorkflowStartTime:
			// lastCompletionResult:
			// lastFailure:
			// CronSchedule:
			// ContinuedExecutionRunID:
			// ParentWorkflowNamespace:
			// ParentWorkflowExecution:
			// Memo:
			// SearchAttributes:
		},
		worker: c,
	}
	w.info.WorkflowStartTime, _ = types.TimestampFromProto(act.Timestamp)

	// Execute
	def.Execute(w, header, input)
	return w, nil
}

func (w *workflowIntercept) Init(next interceptors.WorkflowOutboundCallsInterceptor) error {
	// Use ourself as the outbound interceptor
	w.nextOutbound = next
	return w.nextInbound.Init(w)
}

func (w *workflowIntercept) ExecuteWorkflow(
	ctx workflow.Context,
	workflowType string,
	args ...interface{},
) []interface{} {
	return w.nextInbound.ExecuteWorkflow(ctx, workflowType, args...)
}

func (w *workflowIntercept) Go(ctx workflow.Context, name string, f func(ctx workflow.Context)) workflow.Context {
	return w.nextOutbound.Go(ctx, name, f)
}

func (w *workflowIntercept) GetLogger(ctx workflow.Context) log.Logger {
	return w.nextOutbound.GetLogger(ctx)
}

func (w *workflowEnv) WorkflowInfo() *workflow.Info { return w.info }

func (w *workflowEnv) Complete(result *commonpb.Payloads, err error) {
	// TODO(cretz): This
	w.worker.client.options.Logger.Info("Workflow done", "Result", result, "Error", err)
	if withStackTrace, _ := err.(interface{ StackTrace() string }); withStackTrace != nil {
		w.worker.client.options.Logger.Debug("Error stack: " + withStackTrace.StackTrace())
	}
	panic("TODO")
}

func (w *workflowEnv) RegisterCancelHandler(handler func()) {
	// TODO(cretz): this
}

func (w *workflowEnv) GetLogger() log.Logger {
	return w.worker.client.options.Logger
}

func (w *workflowEnv) RegisterSignalHandler(handler func(name string, input *commonpb.Payloads)) {
	// TODO(cretz): this
}

func (w *workflowEnv) RegisterQueryHandler(
	handler func(queryType string, queryArgs *commonpb.Payloads) (*commonpb.Payloads, error),
) {
	// TODO(cretz): this
}

func (w *workflowEnv) GetDataConverter() converter.DataConverter {
	return w.worker.client.options.DataConverter
}

func (w *workflowEnv) GetContextPropagators() []workflow.ContextPropagator {
	return w.worker.client.options.ContextPropagators
}

func (w *workflowEnv) GetRegistry() *internal.Registry { return w.worker.Registry }
