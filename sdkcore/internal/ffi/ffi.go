package ffi

/*
#cgo CFLAGS:-I${SRCDIR}/sdk-core-bridge/lib/include
#cgo !windows LDFLAGS:-ltemporal_sdk_core_bridge -lm -ldl -pthread
#cgo windows LDFLAGS:-ltemporal_sdk_core_bridge -luserenv -lole32 -lntdll -lws2_32 -lkernel32 -lsecur32 -lcrypt32 -lbcrypt -lncrypt
#cgo linux,amd64 LDFLAGS:-L${SRCDIR}/sdk-core-bridge/lib/linux-x86_64
#cgo linux,arm64 LDFLAGS:-L${SRCDIR}/sdk-core-bridge/lib/linux-aarch64
#cgo darwin,amd64 LDFLAGS:-L${SRCDIR}/sdk-core-bridge/lib/macos-x86_64
#cgo windows,amd64 LDFLAGS:-L${SRCDIR}/sdk-core-bridge/lib/windows-x86_64
#include <sdk-core-bridge.h>
*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"

	"github.com/gogo/protobuf/proto"
	"go.temporal.io/api/common/v1"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
	corepb "go.temporal.io/sdk/sdkcore/internal/ffi/corepb"
	activitytaskpb "go.temporal.io/sdk/sdkcore/internal/ffi/corepb/activitytaskpb"
	workflowactivationpb "go.temporal.io/sdk/sdkcore/internal/ffi/corepb/workflowactivationpb"
	workflowcompletionpb "go.temporal.io/sdk/sdkcore/internal/ffi/corepb/workflowcompletionpb"
)

func HelloRust() {
	C.hello_rust()
}

func copyString(ptr *C.uint8_t, size C.size_t) string {
	in := (*[1<<30 - 1]byte)(unsafe.Pointer(ptr))
	out := make([]byte, size)
	copy(out, in[:size])
	return string(out)
}

func bytesPtrAndLen(b []byte) (*C.uint8_t, C.size_t) {
	return (*C.uint8_t)(C.CBytes(b)), C.size_t(len(b))
}

func bytesPtrFree(ptrs ...*C.uint8_t) {
	for _, ptr := range ptrs {
		C.free((unsafe.Pointer)(ptr))
	}
}

func utf8PtrAndLen(s string) (*C.uint8_t, C.size_t) { return bytesPtrAndLen([]byte(s)) }

func utf8PtrFree(ptrs ...*C.uint8_t) { bytesPtrFree(ptrs...) }

type ErrFFI struct {
	Message string
	Code    int
}

func (e *ErrFFI) Error() string { return e.Message }

// This also frees the error
func errFFIFromError(err *C.struct_tmprl_error_t) error {
	if err == nil {
		return nil
	}
	defer C.tmprl_error_free(err)
	return &ErrFFI{
		Message: copyString(C.tmprl_error_message(err), C.tmprl_error_message_len(err)),
		Code:    int(C.tmprl_error_message_code(err)),
	}
}

type Runtime struct {
	runtime *C.tmprl_runtime_t
}

func NewRuntime() *Runtime {
	return &Runtime{runtime: C.tmprl_runtime_new()}
}

func (r *Runtime) Close() error {
	if r.runtime != nil {
		C.tmprl_runtime_free(r.runtime)
		r.runtime = nil
	}
	return nil
}

type Core struct {
	core *C.tmprl_core_t
}

type CoreConfig struct {
	TargetURL      string
	Namespace      string
	ClientName     string
	ClientVersion  string
	WorkerBinaryID string
}

func NewCore(runtime *Runtime, config CoreConfig) (*Core, error) {
	// Config defaults
	if config.TargetURL == "" {
		config.TargetURL = client.DefaultHostPort
	}
	var opts C.struct_tmprl_core_new_options_t
	opts.target_url, opts.target_url_len = utf8PtrAndLen(config.TargetURL)
	opts.namespace_, opts.namespace_len = utf8PtrAndLen(config.Namespace)
	opts.client_name, opts.client_name_len = utf8PtrAndLen(config.ClientName)
	opts.client_version, opts.client_version_len = utf8PtrAndLen(config.ClientVersion)
	opts.worker_binary_id, opts.worker_binary_id_len = utf8PtrAndLen(config.WorkerBinaryID)
	defer utf8PtrFree(opts.target_url, opts.namespace_, opts.client_name, opts.client_version, opts.worker_binary_id)
	coreOrErr := C.tmprl_core_new(runtime.runtime, &opts)
	if coreOrErr.error != nil {
		return nil, errFFIFromError(coreOrErr.error)
	}
	return &Core{core: coreOrErr.core}, nil
}

// Caller must make absolutely sure no calls are made after this is closed.
func (c *Core) Close() error {
	if c.core != nil {
		C.tmprl_core_free(c.core)
		c.core = nil
	}
	return nil
}

type WorkerConfig struct {
	TaskQueue string
}

func (c *Core) RegisterWorker(config *WorkerConfig) error {
	taskQueuePtr := C.CBytes([]byte(config.TaskQueue))
	defer C.free(taskQueuePtr)
	conf := C.struct_tmprl_worker_config_t{
		task_queue:     (*C.uint8_t)(taskQueuePtr),
		task_queue_len: C.size_t(len(config.TaskQueue)),
	}
	return errFFIFromError(C.tmprl_register_worker(c.core, &conf))
}

func (c *Core) PollWorkflowActivation(taskQueue string) (*workflowactivationpb.WFActivation, error) {
	taskQueueBytes := []byte(taskQueue)
	actOrErr := C.tmprl_poll_workflow_activation(
		c.core,
		(*C.uint8_t)(unsafe.Pointer(&taskQueueBytes[0])),
		C.size_t(len(taskQueueBytes)),
	)
	if err := errFFIFromError(actOrErr.error); err != nil {
		return nil, err
	}
	defer C.tmprl_wf_activation_free(actOrErr.activation)
	var act workflowactivationpb.WFActivation
	err := c.protoFromBytesOrErr(C.tmprl_wf_activation_to_proto(c.core, actOrErr.activation), &act)
	if err != nil {
		return nil, err
	}
	return &act, err
}

func (c *Core) PollActivityTask(taskQueue string) (*activitytaskpb.ActivityTask, error) {
	taskQueueBytes := []byte(taskQueue)
	taskOrErr := C.tmprl_poll_activity_task(
		c.core,
		(*C.uint8_t)(unsafe.Pointer(&taskQueueBytes[0])),
		C.size_t(len(taskQueueBytes)),
	)
	if err := errFFIFromError(taskOrErr.error); err != nil {
		return nil, err
	}
	defer C.tmprl_activity_task_free(taskOrErr.task)
	var task activitytaskpb.ActivityTask
	err := c.protoFromBytesOrErr(C.tmprl_activity_task_to_proto(c.core, taskOrErr.task), &task)
	if err != nil {
		return nil, err
	}
	return &task, err
}

func (c *Core) CompleteWorkflowActivation(completion *workflowcompletionpb.WFActivationCompletion) error {
	b, err := proto.Marshal(completion)
	if err != nil {
		return err
	}
	compOrErr := C.tmprl_wf_activation_completion_from_proto(
		(*C.uint8_t)(unsafe.Pointer(&b[0])),
		C.size_t(len(b)),
	)
	if err := errFFIFromError(compOrErr.error); err != nil {
		return err
	}
	return errFFIFromError(C.tmprl_complete_workflow_activation(c.core, compOrErr.completion))
}

func (c *Core) CompleteActivityTask(completion *corepb.ActivityTaskCompletion) error {
	b, err := proto.Marshal(completion)
	if err != nil {
		return err
	}
	compOrErr := C.tmprl_activity_task_completion_from_proto(
		(*C.uint8_t)(unsafe.Pointer(&b[0])),
		C.size_t(len(b)),
	)
	if err := errFFIFromError(compOrErr.error); err != nil {
		return err
	}
	return errFFIFromError(C.tmprl_complete_activity_task(c.core, compOrErr.completion))
}

func (c *Core) RecordActivityHeartbeat(heartbeat *corepb.ActivityHeartbeat) error {
	b, err := proto.Marshal(heartbeat)
	if err != nil {
		return err
	}
	heartOrErr := C.tmprl_activity_heartbeat_from_proto(
		(*C.uint8_t)(unsafe.Pointer(&b[0])),
		C.size_t(len(b)),
	)
	if err := errFFIFromError(heartOrErr.error); err != nil {
		return err
	}
	C.tmprl_record_activity_heartbeat(c.core, heartOrErr.heartbeat)
	return nil
}

func (c *Core) RequestWorkflowEviction(taskQueue, runID string) {
	taskQueueBytes := []byte(taskQueue)
	runIDBytes := []byte(runID)
	C.tmprl_request_workflow_eviction(
		c.core,
		(*C.uint8_t)(unsafe.Pointer(&taskQueueBytes[0])),
		C.size_t(len(taskQueueBytes)),
		(*C.uint8_t)(unsafe.Pointer(&runIDBytes[0])),
		C.size_t(len(runIDBytes)),
	)
}

func (c *Core) Shutdown() {
	C.tmprl_shutdown(c.core)
}

func (c *Core) ShutdownWorker(taskQueue string) {
	taskQueueBytes := []byte(taskQueue)
	C.tmprl_shutdown_worker(
		c.core,
		(*C.uint8_t)(unsafe.Pointer(&taskQueueBytes[0])),
		C.size_t(len(taskQueueBytes)),
	)
}

type Log struct {
	Message string
	Time    time.Time
	Level   LogLevel
}

type LogLevel uint

const (
	LogLevelError LogLevel = iota + 1
	LogLevelWarn
	LogLevelInfo
	LogLevelDebug
	LogLevelTrace
)

// This can drop messages if not receiving fast enough. Log parameter to
// callback is not reused. This returns when Core is closed or the callback
// returns false.
func (c *Core) RunLogListener(cb func(*Log) bool) {
	l := C.tmprl_log_listener_new(c.core)
	defer C.tmprl_log_listener_free(l)
	// TODO(cretz): When core is closed this may segfault!
	ok := true
	for ok && bool(C.tmprl_log_listener_next(c.core, l)) {
		ok = cb(&Log{
			Message: copyString(C.tmprl_log_listener_curr_message(l), C.tmprl_log_listener_curr_message_len(l)),
			Time:    time.Unix(0, int64(C.tmprl_log_listener_curr_unix_nanos(l))),
			Level:   LogLevel(C.tmprl_log_listener_curr_level(l)),
		})
	}
}

type StartWorkflowRequest struct {
	Input        *common.Payloads
	TaskQueue    string
	WorkflowID   string
	WorkflowType string
	TaskTimeout  time.Duration
}

func (c *Core) StartWorkflow(req *StartWorkflowRequest) (*workflowservice.StartWorkflowExecutionResponse, error) {
	// Create request
	var coreReq C.struct_tmprl_start_workflow_request_t
	coreReq.input_payloads_proto_len = 0
	coreReq.task_queue, coreReq.task_queue_len = utf8PtrAndLen(req.TaskQueue)
	coreReq.workflow_id, coreReq.workflow_id_len = utf8PtrAndLen(req.WorkflowID)
	coreReq.workflow_type, coreReq.workflow_type_len = utf8PtrAndLen(req.WorkflowType)
	defer utf8PtrFree(coreReq.task_queue, coreReq.workflow_id, coreReq.workflow_type)
	// Marshal input
	if req.Input != nil && len(req.Input.Payloads) > 0 {
		b, err := proto.Marshal(req.Input)
		if err != nil {
			return nil, fmt.Errorf("marshaling input: %w", err)
		}
		coreReq.input_payloads_proto, coreReq.input_payloads_proto_len = bytesPtrAndLen(b)
		defer bytesPtrFree(coreReq.input_payloads_proto)
	}

	// Call
	respOrErr := C.tmprl_start_workflow(c.core, &coreReq)
	if err := errFFIFromError(respOrErr.error); err != nil {
		return nil, err
	}
	defer C.tmprl_start_workflow_response_free(respOrErr.response)
	var resp workflowservice.StartWorkflowExecutionResponse
	err := c.protoFromBytesOrErr(C.tmprl_start_workflow_response_to_proto(c.core, respOrErr.response), &resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (c *Core) protoFromBytesOrErr(bytesOrErr C.struct_tmprl_bytes_or_error_t, p proto.Message) error {
	if err := errFFIFromError(bytesOrErr.error); err != nil {
		return err
	}
	defer C.tmprl_bytes_free(c.core, bytesOrErr.bytes)
	in := (*[1<<30 - 1]byte)(unsafe.Pointer(bytesOrErr.bytes.bytes))
	return proto.Unmarshal(in[:bytesOrErr.bytes.len], p)
}
