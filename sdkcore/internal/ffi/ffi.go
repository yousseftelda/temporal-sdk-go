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
	"context"
	"time"
	"unsafe"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/sdkcore/internal/ffi/corepb"
	"go.temporal.io/sdk/sdkcore/internal/ffi/corepb/activitytaskpb"
	"go.temporal.io/sdk/sdkcore/internal/ffi/corepb/workflowactivationpb"
	"go.temporal.io/sdk/sdkcore/internal/ffi/corepb/workflowcompletionpb"
	"google.golang.org/protobuf/proto"
)

func HelloRust() {
	C.hello_rust()
}

func copyString(ptr *C.char, size C.size_t) string {
	in := (*[1<<30 - 1]byte)(unsafe.Pointer(ptr))
	out := make([]byte, size)
	copy(out, in[:size])
	return string(out)
}

type ErrFFI struct {
	Message string
	Code    int
}

func (e *ErrFFI) Error() string { return e.Message }

// This also frees the error
func errFFIFromError(err *C.struct_tmprl_error_t) *ErrFFI {
	if err == nil {
		return nil
	}
	defer C.tmprl_error_free(err)
	return &ErrFFI{
		Message: copyString(C.tmprl_error_message(err), C.tmprl_error_message_len(err)),
		Code:    int(C.tmprl_error_message_code(err)),
	}
}

type Core struct {
	runtime *C.tmprl_runtime_t
	core    *C.tmprl_core_t
}

func NewCore() (*Core, error) {
	// Create runtime and core
	runtime := C.tmprl_runtime_new()
	targetURL := []byte(client.DefaultHostPort)
	coreOpts := C.struct_tmprl_core_new_options_t{
		// TODO(cretz): More options
		target_url:     (*C.char)(unsafe.Pointer(&targetURL[0])),
		target_url_len: C.size_t(len(targetURL)),
	}
	coreOrErr := C.tmprl_core_new(runtime, &coreOpts)
	if coreOrErr.error != nil {
		// Free the runtime before returning error
		C.tmprl_runtime_free(runtime)
		return nil, errFFIFromError(coreOrErr.error)
	}
	return &Core{runtime: runtime, core: coreOrErr.core}, nil
}

// Caller must make absolutely sure no calls are made after this is closed.
func (c *Core) Close() error {
	if c.core != nil {
		C.tmprl_core_free(c.core)
		c.core = nil
	}
	if c.runtime != nil {
		C.tmprl_runtime_free(c.runtime)
		c.runtime = nil
	}
	return nil
}

type WorkerConfig struct {
	TaskQueue string
}

func (c *Core) RegisterWorker(config *WorkerConfig) error {
	taskQueueBytes := []byte(config.TaskQueue)
	conf := C.struct_tmprl_worker_config_t{
		task_queue:     (*C.char)(unsafe.Pointer(&taskQueueBytes[0])),
		task_queue_len: C.size_t(len(taskQueueBytes)),
	}
	return errFFIFromError(C.tmprl_register_worker(c.core, &conf))
}

func (c *Core) PollWorkflowActivation(taskQueue string) (*workflowactivationpb.WFActivation, error) {
	taskQueueBytes := []byte(taskQueue)
	actOrErr := C.tmprl_poll_workflow_activation(
		c.core,
		(*C.char)(unsafe.Pointer(&taskQueueBytes[0])),
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
		(*C.char)(unsafe.Pointer(&taskQueueBytes[0])),
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
		(*C.char)(unsafe.Pointer(&taskQueueBytes[0])),
		C.size_t(len(taskQueueBytes)),
		(*C.char)(unsafe.Pointer(&runIDBytes[0])),
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
		(*C.char)(unsafe.Pointer(&taskQueueBytes[0])),
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

func (c *Core) RunLogListener(ctx context.Context, cb func(*Log) bool) {
	panic("TODO")
}

func (c *Core) protoFromBytesOrErr(bytesOrErr C.struct_tmprl_bytes_or_error_t, p proto.Message) error {
	if err := errFFIFromError(bytesOrErr.error); err != nil {
		return err
	}
	defer C.tmprl_bytes_free(c.core, bytesOrErr.bytes)
	in := (*[1<<30 - 1]byte)(unsafe.Pointer(bytesOrErr.bytes.bytes))
	return proto.Unmarshal(in[:bytesOrErr.bytes.len], p)
}
