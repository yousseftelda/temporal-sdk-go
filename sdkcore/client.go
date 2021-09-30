package sdkcore

import (
	"context"

	enumspb "go.temporal.io/api/enums/v1"
	workflowservicepb "go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/converter"
)

type coreClient struct {
}

func NewClient(options client.Options) (client.Client, error) {
	// TODO
	return &coreClient{}, nil
}

func (*coreClient) ExecuteWorkflow(
	ctx context.Context,
	options client.StartWorkflowOptions,
	workflow interface{},
	args ...interface{},
) (client.WorkflowRun, error) {
	panic("TODO")
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

func (*coreClient) Close() {
	panic("TODO")
}
