package sdkcore

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"reflect"
	"runtime"
	"strings"
	"sync"
)

func functionName(i interface{}) (name string, isMethod bool) {
	if fullName, ok := i.(string); ok {
		return fullName, false
	}
	fullName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	// Full function name that has a struct pointer receiver has the following format
	// <prefix>.(*<type>).<function>
	isMethod = strings.ContainsAny(fullName, "*")
	elements := strings.Split(fullName, ".")
	shortName := elements[len(elements)-1]
	// This allows to call activities by method pointer
	// Compiler adds -fm suffix to a function name which has a receiver
	// Note that this works even if struct pointer used to get the function is nil
	// It is possible because nil receivers are allowed.
	// For example:
	// var a *Activities
	// ExecuteActivity(ctx, a.Foo)
	// will call this function which is going to return "Foo"
	return strings.TrimSuffix(shortName, "-fm"), isMethod
}

var binaryChecksumOnce sync.Once
var binaryChecksumStr string
var binaryChecksumErr error

func binaryChecksum() (string, error) {
	binaryChecksumOnce.Do(func() { binaryChecksumStr, binaryChecksumErr = buildBinaryChecksum() })
	return binaryChecksumStr, binaryChecksumErr
}

func buildBinaryChecksum() (string, error) {
	exec, err := os.Executable()
	if err != nil {
		return "", err
	}
	f, err := os.Open(exec)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)[:]), nil
}
