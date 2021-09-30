package ffi

// #cgo CFLAGS:-I${SRCDIR}/sdk-core-bridge/lib/include
// #cgo !windows LDFLAGS:-ltemporal_sdk_core_bridge -lm -ldl -pthread
// #cgo windows LDFLAGS:-ltemporal_sdk_core_bridge -luserenv -lole32 -lntdll -lws2_32 -lkernel32 -lbcrypt
// #cgo linux,amd64 LDFLAGS:-L${SRCDIR}/sdk-core-bridge/lib/linux-x86_64
// #cgo linux,arm64 LDFLAGS:-L${SRCDIR}/sdk-core-bridge/lib/linux-aarch64
// #cgo darwin,amd64 LDFLAGS:-L${SRCDIR}/sdk-core-bridge/lib/macos-x86_64
// #cgo windows,amd64 LDFLAGS:-L${SRCDIR}/sdk-core-bridge/lib/windows-x86_64
// #include <sdk-core-bridge.h>
import "C"

func HelloRust() {
	C.hello_rust()
}
