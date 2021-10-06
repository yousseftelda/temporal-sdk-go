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
	"unsafe"

	"go.temporal.io/sdk/client"
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
	taskQueue := []byte(config.TaskQueue)
	conf := C.struct_tmprl_worker_config_t{
		task_queue:     (*C.char)(unsafe.Pointer(&taskQueue[0])),
		task_queue_len: C.size_t(len(taskQueue)),
	}
	return errFFIFromError(C.tmprl_register_worker(c.core, &conf))
}

// TODO(cretz): More functions
