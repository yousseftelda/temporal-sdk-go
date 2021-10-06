// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: child_workflow.proto

package childworkflowpb

import (
	v1 "go.temporal.io/api/failure/v1"
	commonpb "go.temporal.io/sdk/sdkcore/internal/ffi/corepb/commonpb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//*
// Used by the service to determine the fate of a child workflow
// in case its parent is closed.
type ParentClosePolicy int32

const (
	//* Let's the server set the default.
	ParentClosePolicy_PARENT_CLOSE_POLICY_UNSPECIFIED ParentClosePolicy = 0
	//* Terminate means terminating the child workflow.
	ParentClosePolicy_PARENT_CLOSE_POLICY_TERMINATE ParentClosePolicy = 1
	//* Abandon means not doing anything on the child workflow.
	ParentClosePolicy_PARENT_CLOSE_POLICY_ABANDON ParentClosePolicy = 2
	//* Cancel means requesting cancellation on the child workflow.
	ParentClosePolicy_PARENT_CLOSE_POLICY_REQUEST_CANCEL ParentClosePolicy = 3
)

// Enum value maps for ParentClosePolicy.
var (
	ParentClosePolicy_name = map[int32]string{
		0: "PARENT_CLOSE_POLICY_UNSPECIFIED",
		1: "PARENT_CLOSE_POLICY_TERMINATE",
		2: "PARENT_CLOSE_POLICY_ABANDON",
		3: "PARENT_CLOSE_POLICY_REQUEST_CANCEL",
	}
	ParentClosePolicy_value = map[string]int32{
		"PARENT_CLOSE_POLICY_UNSPECIFIED":    0,
		"PARENT_CLOSE_POLICY_TERMINATE":      1,
		"PARENT_CLOSE_POLICY_ABANDON":        2,
		"PARENT_CLOSE_POLICY_REQUEST_CANCEL": 3,
	}
)

func (x ParentClosePolicy) Enum() *ParentClosePolicy {
	p := new(ParentClosePolicy)
	*p = x
	return p
}

func (x ParentClosePolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ParentClosePolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_child_workflow_proto_enumTypes[0].Descriptor()
}

func (ParentClosePolicy) Type() protoreflect.EnumType {
	return &file_child_workflow_proto_enumTypes[0]
}

func (x ParentClosePolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ParentClosePolicy.Descriptor instead.
func (ParentClosePolicy) EnumDescriptor() ([]byte, []int) {
	return file_child_workflow_proto_rawDescGZIP(), []int{0}
}

//* Possible causes of failure to start a child workflow
type StartChildWorkflowExecutionFailedCause int32

const (
	StartChildWorkflowExecutionFailedCause_START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED             StartChildWorkflowExecutionFailedCause = 0
	StartChildWorkflowExecutionFailedCause_START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_WORKFLOW_ALREADY_EXISTS StartChildWorkflowExecutionFailedCause = 1
)

// Enum value maps for StartChildWorkflowExecutionFailedCause.
var (
	StartChildWorkflowExecutionFailedCause_name = map[int32]string{
		0: "START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED",
		1: "START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_WORKFLOW_ALREADY_EXISTS",
	}
	StartChildWorkflowExecutionFailedCause_value = map[string]int32{
		"START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED":             0,
		"START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_WORKFLOW_ALREADY_EXISTS": 1,
	}
)

func (x StartChildWorkflowExecutionFailedCause) Enum() *StartChildWorkflowExecutionFailedCause {
	p := new(StartChildWorkflowExecutionFailedCause)
	*p = x
	return p
}

func (x StartChildWorkflowExecutionFailedCause) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StartChildWorkflowExecutionFailedCause) Descriptor() protoreflect.EnumDescriptor {
	return file_child_workflow_proto_enumTypes[1].Descriptor()
}

func (StartChildWorkflowExecutionFailedCause) Type() protoreflect.EnumType {
	return &file_child_workflow_proto_enumTypes[1]
}

func (x StartChildWorkflowExecutionFailedCause) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StartChildWorkflowExecutionFailedCause.Descriptor instead.
func (StartChildWorkflowExecutionFailedCause) EnumDescriptor() ([]byte, []int) {
	return file_child_workflow_proto_rawDescGZIP(), []int{1}
}

//*
// Controls at which point to report back to lang when a child workflow is cancelled
type ChildWorkflowCancellationType int32

const (
	//* Do not request cancellation of the child workflow if already scheduled
	ChildWorkflowCancellationType_ABANDON ChildWorkflowCancellationType = 0
	//* Initiate a cancellation request and immediately report cancellation to the parent.
	ChildWorkflowCancellationType_TRY_CANCEL ChildWorkflowCancellationType = 1
	//* Wait for child cancellation completion.
	ChildWorkflowCancellationType_WAIT_CANCELLATION_COMPLETED ChildWorkflowCancellationType = 2
	//* Request cancellation of the child and wait for confirmation that the request was received.
	ChildWorkflowCancellationType_WAIT_CANCELLATION_REQUESTED ChildWorkflowCancellationType = 3
)

// Enum value maps for ChildWorkflowCancellationType.
var (
	ChildWorkflowCancellationType_name = map[int32]string{
		0: "ABANDON",
		1: "TRY_CANCEL",
		2: "WAIT_CANCELLATION_COMPLETED",
		3: "WAIT_CANCELLATION_REQUESTED",
	}
	ChildWorkflowCancellationType_value = map[string]int32{
		"ABANDON":                     0,
		"TRY_CANCEL":                  1,
		"WAIT_CANCELLATION_COMPLETED": 2,
		"WAIT_CANCELLATION_REQUESTED": 3,
	}
)

func (x ChildWorkflowCancellationType) Enum() *ChildWorkflowCancellationType {
	p := new(ChildWorkflowCancellationType)
	*p = x
	return p
}

func (x ChildWorkflowCancellationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChildWorkflowCancellationType) Descriptor() protoreflect.EnumDescriptor {
	return file_child_workflow_proto_enumTypes[2].Descriptor()
}

func (ChildWorkflowCancellationType) Type() protoreflect.EnumType {
	return &file_child_workflow_proto_enumTypes[2]
}

func (x ChildWorkflowCancellationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChildWorkflowCancellationType.Descriptor instead.
func (ChildWorkflowCancellationType) EnumDescriptor() ([]byte, []int) {
	return file_child_workflow_proto_rawDescGZIP(), []int{2}
}

//*
// Used by core to resolve child workflow executions.
type ChildWorkflowResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Status:
	//	*ChildWorkflowResult_Completed
	//	*ChildWorkflowResult_Failed
	//	*ChildWorkflowResult_Cancelled
	Status isChildWorkflowResult_Status `protobuf_oneof:"status"`
}

func (x *ChildWorkflowResult) Reset() {
	*x = ChildWorkflowResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_child_workflow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChildWorkflowResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChildWorkflowResult) ProtoMessage() {}

func (x *ChildWorkflowResult) ProtoReflect() protoreflect.Message {
	mi := &file_child_workflow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChildWorkflowResult.ProtoReflect.Descriptor instead.
func (*ChildWorkflowResult) Descriptor() ([]byte, []int) {
	return file_child_workflow_proto_rawDescGZIP(), []int{0}
}

func (m *ChildWorkflowResult) GetStatus() isChildWorkflowResult_Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (x *ChildWorkflowResult) GetCompleted() *Success {
	if x, ok := x.GetStatus().(*ChildWorkflowResult_Completed); ok {
		return x.Completed
	}
	return nil
}

func (x *ChildWorkflowResult) GetFailed() *Failure {
	if x, ok := x.GetStatus().(*ChildWorkflowResult_Failed); ok {
		return x.Failed
	}
	return nil
}

func (x *ChildWorkflowResult) GetCancelled() *Cancellation {
	if x, ok := x.GetStatus().(*ChildWorkflowResult_Cancelled); ok {
		return x.Cancelled
	}
	return nil
}

type isChildWorkflowResult_Status interface {
	isChildWorkflowResult_Status()
}

type ChildWorkflowResult_Completed struct {
	Completed *Success `protobuf:"bytes,1,opt,name=completed,proto3,oneof"`
}

type ChildWorkflowResult_Failed struct {
	Failed *Failure `protobuf:"bytes,2,opt,name=failed,proto3,oneof"`
}

type ChildWorkflowResult_Cancelled struct {
	Cancelled *Cancellation `protobuf:"bytes,3,opt,name=cancelled,proto3,oneof"`
}

func (*ChildWorkflowResult_Completed) isChildWorkflowResult_Status() {}

func (*ChildWorkflowResult_Failed) isChildWorkflowResult_Status() {}

func (*ChildWorkflowResult_Cancelled) isChildWorkflowResult_Status() {}

//*
// Used in ChildWorkflowResult to report successful completion.
type Success struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *commonpb.Payload `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *Success) Reset() {
	*x = Success{}
	if protoimpl.UnsafeEnabled {
		mi := &file_child_workflow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Success) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Success) ProtoMessage() {}

func (x *Success) ProtoReflect() protoreflect.Message {
	mi := &file_child_workflow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Success.ProtoReflect.Descriptor instead.
func (*Success) Descriptor() ([]byte, []int) {
	return file_child_workflow_proto_rawDescGZIP(), []int{1}
}

func (x *Success) GetResult() *commonpb.Payload {
	if x != nil {
		return x.Result
	}
	return nil
}

//*
// Used in ChildWorkflowResult to report non successful outcomes such as
// application failures, timeouts, terminations, and cancellations.
type Failure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Failure *v1.Failure `protobuf:"bytes,1,opt,name=failure,proto3" json:"failure,omitempty"`
}

func (x *Failure) Reset() {
	*x = Failure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_child_workflow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Failure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Failure) ProtoMessage() {}

func (x *Failure) ProtoReflect() protoreflect.Message {
	mi := &file_child_workflow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Failure.ProtoReflect.Descriptor instead.
func (*Failure) Descriptor() ([]byte, []int) {
	return file_child_workflow_proto_rawDescGZIP(), []int{2}
}

func (x *Failure) GetFailure() *v1.Failure {
	if x != nil {
		return x.Failure
	}
	return nil
}

//*
// Used in ChildWorkflowResult to report cancellation.
// Failure should be ChildWorkflowFailure with a CanceledFailure cause.
type Cancellation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Failure *v1.Failure `protobuf:"bytes,1,opt,name=failure,proto3" json:"failure,omitempty"`
}

func (x *Cancellation) Reset() {
	*x = Cancellation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_child_workflow_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cancellation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cancellation) ProtoMessage() {}

func (x *Cancellation) ProtoReflect() protoreflect.Message {
	mi := &file_child_workflow_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cancellation.ProtoReflect.Descriptor instead.
func (*Cancellation) Descriptor() ([]byte, []int) {
	return file_child_workflow_proto_rawDescGZIP(), []int{3}
}

func (x *Cancellation) GetFailure() *v1.Failure {
	if x != nil {
		return x.Failure
	}
	return nil
}

var File_child_workflow_proto protoreflect.FileDescriptor

var file_child_workflow_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x64, 0x6b, 0x2e,
	0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x0c,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x61, 0x69, 0x6c, 0x75,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x01, 0x0a, 0x13, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3f, 0x0a, 0x09, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x64, 0x6b, 0x2e, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x48,
	0x00, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x06,
	0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x73, 0x64, 0x6b, 0x2e, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x48, 0x00, 0x52,
	0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x44, 0x0a, 0x09, 0x63, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x64, 0x6b, 0x2e, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x09, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x42, 0x08, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3a, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x64, 0x6b, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x45, 0x0a, 0x07, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x3a,
	0x0a, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x52, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x22, 0x4a, 0x0a, 0x0c, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x07, 0x66, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x61, 0x69, 0x6c, 0x75,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x07, 0x66,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x2a, 0xa4, 0x01, 0x0a, 0x11, 0x50, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x23, 0x0a, 0x1f,
	0x50, 0x41, 0x52, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x5f, 0x50, 0x4f, 0x4c,
	0x49, 0x43, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x21, 0x0a, 0x1d, 0x50, 0x41, 0x52, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4c, 0x4f, 0x53,
	0x45, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x54, 0x45, 0x52, 0x4d, 0x49, 0x4e, 0x41,
	0x54, 0x45, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x50, 0x41, 0x52, 0x45, 0x4e, 0x54, 0x5f, 0x43,
	0x4c, 0x4f, 0x53, 0x45, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x41, 0x42, 0x41, 0x4e,
	0x44, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x26, 0x0a, 0x22, 0x50, 0x41, 0x52, 0x45, 0x4e, 0x54, 0x5f,
	0x43, 0x4c, 0x4f, 0x53, 0x45, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x10, 0x03, 0x2a, 0xae, 0x01,
	0x0a, 0x26, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x43, 0x61, 0x75, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x37, 0x53, 0x54, 0x41, 0x52,
	0x54, 0x5f, 0x43, 0x48, 0x49, 0x4c, 0x44, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57,
	0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45,
	0x44, 0x5f, 0x43, 0x41, 0x55, 0x53, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x47, 0x0a, 0x43, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x43,
	0x48, 0x49, 0x4c, 0x44, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58,
	0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x43,
	0x41, 0x55, 0x53, 0x45, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x41, 0x4c,
	0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x01, 0x2a, 0x7e,
	0x0a, 0x1d, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x41, 0x42, 0x41, 0x4e, 0x44, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a,
	0x54, 0x52, 0x59, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b,
	0x57, 0x41, 0x49, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1f, 0x0a,
	0x1b, 0x57, 0x41, 0x49, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x44, 0x10, 0x03, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_child_workflow_proto_rawDescOnce sync.Once
	file_child_workflow_proto_rawDescData = file_child_workflow_proto_rawDesc
)

func file_child_workflow_proto_rawDescGZIP() []byte {
	file_child_workflow_proto_rawDescOnce.Do(func() {
		file_child_workflow_proto_rawDescData = protoimpl.X.CompressGZIP(file_child_workflow_proto_rawDescData)
	})
	return file_child_workflow_proto_rawDescData
}

var file_child_workflow_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_child_workflow_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_child_workflow_proto_goTypes = []interface{}{
	(ParentClosePolicy)(0),                      // 0: coresdk.child_workflow.ParentClosePolicy
	(StartChildWorkflowExecutionFailedCause)(0), // 1: coresdk.child_workflow.StartChildWorkflowExecutionFailedCause
	(ChildWorkflowCancellationType)(0),          // 2: coresdk.child_workflow.ChildWorkflowCancellationType
	(*ChildWorkflowResult)(nil),                 // 3: coresdk.child_workflow.ChildWorkflowResult
	(*Success)(nil),                             // 4: coresdk.child_workflow.Success
	(*Failure)(nil),                             // 5: coresdk.child_workflow.Failure
	(*Cancellation)(nil),                        // 6: coresdk.child_workflow.Cancellation
	(*commonpb.Payload)(nil),                    // 7: coresdk.common.Payload
	(*v1.Failure)(nil),                          // 8: temporal.api.failure.v1.Failure
}
var file_child_workflow_proto_depIdxs = []int32{
	4, // 0: coresdk.child_workflow.ChildWorkflowResult.completed:type_name -> coresdk.child_workflow.Success
	5, // 1: coresdk.child_workflow.ChildWorkflowResult.failed:type_name -> coresdk.child_workflow.Failure
	6, // 2: coresdk.child_workflow.ChildWorkflowResult.cancelled:type_name -> coresdk.child_workflow.Cancellation
	7, // 3: coresdk.child_workflow.Success.result:type_name -> coresdk.common.Payload
	8, // 4: coresdk.child_workflow.Failure.failure:type_name -> temporal.api.failure.v1.Failure
	8, // 5: coresdk.child_workflow.Cancellation.failure:type_name -> temporal.api.failure.v1.Failure
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_child_workflow_proto_init() }
func file_child_workflow_proto_init() {
	if File_child_workflow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_child_workflow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChildWorkflowResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_child_workflow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Success); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_child_workflow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Failure); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_child_workflow_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cancellation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_child_workflow_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ChildWorkflowResult_Completed)(nil),
		(*ChildWorkflowResult_Failed)(nil),
		(*ChildWorkflowResult_Cancelled)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_child_workflow_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_child_workflow_proto_goTypes,
		DependencyIndexes: file_child_workflow_proto_depIdxs,
		EnumInfos:         file_child_workflow_proto_enumTypes,
		MessageInfos:      file_child_workflow_proto_msgTypes,
	}.Build()
	File_child_workflow_proto = out.File
	file_child_workflow_proto_rawDesc = nil
	file_child_workflow_proto_goTypes = nil
	file_child_workflow_proto_depIdxs = nil
}
