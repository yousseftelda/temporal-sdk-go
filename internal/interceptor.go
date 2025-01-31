// The MIT License
//
// Copyright (c) 2021 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package internal

import (
	"context"
	"time"

	"github.com/uber-go/tally/v4"
	commonpb "go.temporal.io/api/common/v1"
	"go.temporal.io/sdk/converter"
	"go.temporal.io/sdk/log"
)

// Interceptor is a common interface for all interceptors. See documentation in
// the interceptor package for more details.
type Interceptor interface {
	ClientInterceptor
	WorkerInterceptor
}

// WorkerInterceptor is a common interface for all interceptors. See
// documentation in the interceptor package for more details.
type WorkerInterceptor interface {
	// InterceptActivity is called before each activity interception needed with
	// the next interceptor in the chain.
	InterceptActivity(ctx context.Context, next ActivityInboundInterceptor) ActivityInboundInterceptor

	// InterceptWorkflow is called before each workflow interception needed with
	// the next interceptor in the chain.
	InterceptWorkflow(ctx Context, next WorkflowInboundInterceptor) WorkflowInboundInterceptor

	mustEmbedWorkerInterceptorBase()
}

// ActivityInboundInterceptor is an interface for all activity calls originating
// from the server. See documentation in the interceptor package for more
// details.
type ActivityInboundInterceptor interface {
	// Init is the first call of this interceptor. Implementations can change/wrap
	// the outbound interceptor before calling Init on the next interceptor.
	Init(outbound ActivityOutboundInterceptor) error

	// ExecuteActivity is called when an activity is to be run on this worker.
	// interceptor.Header will return a non-nil map for this context.
	ExecuteActivity(ctx context.Context, in *ExecuteActivityInput) (interface{}, error)

	mustEmbedActivityInboundInterceptorBase()
}

// ExecuteActivityInput is the input to ActivityInboundInterceptor.ExecuteActivity.
type ExecuteActivityInput struct {
	Args []interface{}
}

// ActivityOutboundInterceptor is an interface for all activity calls
// originating from the SDK. See documentation in the interceptor package for
// more details.
type ActivityOutboundInterceptor interface {
	// GetInfo intercepts activity.GetInfo.
	GetInfo(ctx context.Context) ActivityInfo

	// GetLogger intercepts activity.GetLogger.
	GetLogger(ctx context.Context) log.Logger

	// GetMetricsScope intercepts activity.GetMetricsScope.
	GetMetricsScope(ctx context.Context) tally.Scope

	// RecordHeartbeat intercepts activity.RecordHeartbeat.
	RecordHeartbeat(ctx context.Context, details ...interface{})

	// HasHeartbeatDetails intercepts activity.HasHeartbeatDetails.
	HasHeartbeatDetails(ctx context.Context) bool

	// GetHeartbeatDetails intercepts activity.GetHeartbeatDetails.
	GetHeartbeatDetails(ctx context.Context, d ...interface{}) error

	// GetWorkerStopChannel intercepts activity.GetWorkerStopChannel.
	GetWorkerStopChannel(ctx context.Context) <-chan struct{}

	mustEmbedActivityOutboundInterceptorBase()
}

// WorkflowInboundInterceptor is an interface for all workflow calls originating
// from the server. See documentation in the interceptor package for more
// details.
type WorkflowInboundInterceptor interface {
	// Init is the first call of this interceptor. Implementations can change/wrap
	// the outbound interceptor before calling Init on the next interceptor.
	Init(outbound WorkflowOutboundInterceptor) error

	// ExecuteWorkflow is called when a workflow is to be run on this worker.
	// interceptor.WorkflowHeader will return a non-nil map for this context.
	ExecuteWorkflow(ctx Context, in *ExecuteWorkflowInput) (interface{}, error)

	// HandleSignal is called when a signal is sent to a workflow on this worker.
	HandleSignal(ctx Context, in *HandleSignalInput) error

	// HandleQuery is called when a query is sent to a workflow on this worker.
	HandleQuery(ctx Context, in *HandleQueryInput) (interface{}, error)

	mustEmbedWorkflowInboundInterceptorBase()
}

// ExecuteWorkflowInput is the input to
// WorkflowInboundInterceptor.ExecuteWorkflow.
type ExecuteWorkflowInput struct {
	Args []interface{}
}

// HandleSignalInput is the input to WorkflowInboundInterceptor.HandleSignal.
type HandleSignalInput struct {
	SignalName string
	// Arg is the signal argument. It is presented as a primitive payload since
	// the type needed for decode is not available at the time of interception.
	Arg *commonpb.Payloads
}

// HandleQueryInput is the input to WorkflowInboundInterceptor.HandleQuery.
type HandleQueryInput struct {
	QueryType string
	Args      []interface{}
}

// WorkflowOutboundInterceptor is an interface for all workflow calls
// originating from the SDK. See documentation in the interceptor package for
// more details.
type WorkflowOutboundInterceptor interface {
	// Go intercepts workflow.Go.
	Go(ctx Context, name string, f func(ctx Context)) Context

	// ExecuteActivity intercepts workflow.ExecuteActivity.
	// interceptor.WorkflowHeader will return a non-nil map for this context.
	ExecuteActivity(ctx Context, activityType string, args ...interface{}) Future

	// ExecuteLocalActivity intercepts workflow.ExecuteLocalActivity.
	// interceptor.WorkflowHeader will return a non-nil map for this context.
	ExecuteLocalActivity(ctx Context, activityType string, args ...interface{}) Future

	// ExecuteChildWorkflow intercepts workflow.ExecuteChildWorkflow.
	// interceptor.WorkflowHeader will return a non-nil map for this context.
	ExecuteChildWorkflow(ctx Context, childWorkflowType string, args ...interface{}) ChildWorkflowFuture

	// GetInfo intercepts workflow.GetInfo.
	GetInfo(ctx Context) *WorkflowInfo

	// GetLogger intercepts workflow.GetLogger.
	GetLogger(ctx Context) log.Logger

	// GetMetricsScope intercepts workflow.GetMetricsScope.
	GetMetricsScope(ctx Context) tally.Scope

	// Now intercepts workflow.Now.
	Now(ctx Context) time.Time

	// NewTimer intercepts workflow.NewTimer.
	NewTimer(ctx Context, d time.Duration) Future

	// Sleep intercepts workflow.Sleep.
	Sleep(ctx Context, d time.Duration) (err error)

	// RequestCancelExternalWorkflow intercepts
	// workflow.RequestCancelExternalWorkflow.
	RequestCancelExternalWorkflow(ctx Context, workflowID, runID string) Future

	// SignalExternalWorkflow intercepts workflow.SignalExternalWorkflow.
	SignalExternalWorkflow(ctx Context, workflowID, runID, signalName string, arg interface{}) Future

	// UpsertSearchAttributes intercepts workflow.UpsertSearchAttributes.
	UpsertSearchAttributes(ctx Context, attributes map[string]interface{}) error

	// GetSignalChannel intercepts workflow.GetSignalChannel.
	GetSignalChannel(ctx Context, signalName string) ReceiveChannel

	// SideEffect intercepts workflow.SideEffect.
	SideEffect(ctx Context, f func(ctx Context) interface{}) converter.EncodedValue

	// MutableSideEffect intercepts workflow.MutableSideEffect.
	MutableSideEffect(
		ctx Context,
		id string,
		f func(ctx Context) interface{},
		equals func(a, b interface{}) bool,
	) converter.EncodedValue

	// GetVersion intercepts workflow.GetVersion.
	GetVersion(ctx Context, changeID string, minSupported, maxSupported Version) Version

	// SetQueryHandler intercepts workflow.SetQueryHandler.
	SetQueryHandler(ctx Context, queryType string, handler interface{}) error

	// IsReplaying intercepts workflow.IsReplaying.
	IsReplaying(ctx Context) bool

	// HasLastCompletionResult intercepts workflow.HasLastCompletionResult.
	HasLastCompletionResult(ctx Context) bool

	// GetLastCompletionResult intercepts workflow.GetLastCompletionResult.
	GetLastCompletionResult(ctx Context, d ...interface{}) error

	// GetLastError intercepts workflow.GetLastError.
	GetLastError(ctx Context) error

	// NewContinueAsNewError intercepts workflow.NewContinueAsNewError.
	// interceptor.WorkflowHeader will return a non-nil map for this context.
	NewContinueAsNewError(ctx Context, wfn interface{}, args ...interface{}) error

	mustEmbedWorkflowOutboundInterceptorBase()
}

// ClientInterceptor for providing a ClientOutboundInterceptor to intercept
// certain workflow-specific client calls from the SDK. See documentation in the
// interceptor package for more details.
type ClientInterceptor interface {
	// This is called on client creation if set via client options
	InterceptClient(next ClientOutboundInterceptor) ClientOutboundInterceptor

	mustEmbedClientInterceptorBase()
}

// ClientOutboundInterceptor is an interface for certain workflow-specific calls
// originating from the SDK. See documentation in the interceptor package for
// more details.
type ClientOutboundInterceptor interface {
	// ExecuteWorkflow intercepts client.Client.ExecuteWorkflow.
	// interceptor.Header will return a non-nil map for this context.
	ExecuteWorkflow(context.Context, *ClientExecuteWorkflowInput) (WorkflowRun, error)

	// SignalWorkflow intercepts client.Client.SignalWorkflow.
	SignalWorkflow(context.Context, *ClientSignalWorkflowInput) error

	// SignalWithStartWorkflow intercepts client.Client.SignalWithStartWorkflow.
	// interceptor.Header will return a non-nil map for this context.
	SignalWithStartWorkflow(context.Context, *ClientSignalWithStartWorkflowInput) (WorkflowRun, error)

	// CancelWorkflow intercepts client.Client.CancelWorkflow.
	CancelWorkflow(context.Context, *ClientCancelWorkflowInput) error

	// TerminateWorkflow intercepts client.Client.TerminateWorkflow.
	TerminateWorkflow(context.Context, *ClientTerminateWorkflowInput) error

	// QueryWorkflow intercepts client.Client.QueryWorkflow.
	QueryWorkflow(context.Context, *ClientQueryWorkflowInput) (converter.EncodedValue, error)

	mustEmbedClientOutboundInterceptorBase()
}

// ClientExecuteWorkflowInput is the input to
// ClientOutboundInterceptor.ExecuteWorkflow.
type ClientExecuteWorkflowInput struct {
	Options      *StartWorkflowOptions
	WorkflowType string
	Args         []interface{}
}

// ClientSignalWorkflowInput is the input to
// ClientOutboundInterceptor.SignalWorkflow.
type ClientSignalWorkflowInput struct {
	WorkflowID string
	RunID      string
	SignalName string
	Arg        interface{}
}

// ClientSignalWithStartWorkflowInput is the input to
// ClientOutboundInterceptor.SignalWithStartWorkflow.
type ClientSignalWithStartWorkflowInput struct {
	SignalName   string
	SignalArg    interface{}
	Options      *StartWorkflowOptions
	WorkflowType string
	Args         []interface{}
}

// ClientCancelWorkflowInput is the input to
// ClientOutboundInterceptor.CancelWorkflow.
type ClientCancelWorkflowInput struct {
	WorkflowID string
	RunID      string
}

// ClientTerminateWorkflowInput is the input to
// ClientOutboundInterceptor.TerminateWorkflow.
type ClientTerminateWorkflowInput struct {
	WorkflowID string
	RunID      string
	Reason     string
	Details    []interface{}
}

// ClientQueryWorkflowInput is the input to
// ClientOutboundInterceptor.QueryWorkflow.
type ClientQueryWorkflowInput struct {
	WorkflowID string
	RunID      string
	QueryType  string
	Args       []interface{}
}
