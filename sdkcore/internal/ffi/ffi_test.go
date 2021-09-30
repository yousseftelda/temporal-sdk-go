package ffi_test

import (
	"testing"

	"go.temporal.io/sdk/sdkcore/internal/ffi"
)

func TestFFI(t *testing.T) {
	ffi.HelloRust()
}
