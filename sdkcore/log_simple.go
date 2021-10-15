package sdkcore

import (
	"fmt"
	"log"
)

type logLevel int

const (
	logLevelNone logLevel = iota
	logLevelError
	logLevelWarn
	logLevelInfo
	logLevelDebug
	logLevelAll
)

func (l logLevel) log(level logLevel, levelStr string, msg string, keyVals ...interface{}) {
	if l >= level {
		suffix := ""
		for i := 0; i < len(keyVals); i += 2 {
			suffix += fmt.Sprintf("\n    %v: %v", keyVals[i], keyVals[i+1])
		}
		log.Printf("[%v] %v%v", levelStr, msg, suffix)
	}
}

func (l logLevel) Debug(msg string, keyVals ...interface{}) {
	l.log(logLevelDebug, "DEBUG", msg, keyVals...)
}
func (l logLevel) Info(msg string, keyVals ...interface{}) {
	l.log(logLevelInfo, "INFO", msg, keyVals...)
}
func (l logLevel) Warn(msg string, keyVals ...interface{}) {
	l.log(logLevelWarn, "WARN", msg, keyVals...)
}
func (l logLevel) Error(msg string, keyVals ...interface{}) {
	l.log(logLevelError, "ERROR", msg, keyVals...)
}
