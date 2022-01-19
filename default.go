package glogger

import glog "github.com/andriyg76/glog"

func Default() Logger {
	return glog.Default()
}

func SetLevel(logLevel LogLevel) {
	glog.SetLevel(logLevel)
}

// Trace is deprecated, use github.com/andriyg76/glog.Trace
func Trace(format string, a ...interface{}) {
	glog.Trace(format, a...)
}

func IsTrace() bool {
	return glog.IsTrace()
}

// Debug is deprecated, use github.com/andriyg76/glog.Debug
func Debug(format string, a ...interface{}) {
	glog.Debug(format, a...)
}

func IsDebug() bool {
	return glog.IsDebug()
}

// Info is deprecated, use github.com/andriyg76/glog.Info
func Info(format string, a ...interface{}) {
	glog.Info(format, a...)
}

func IsInfo() bool {
	return glog.IsInfo()
}

// Warn is deprecated, use github.com/andriyg76/glog.Warn
func Warn(format string, a ...interface{}) {
	glog.Warn(format, a...)
}

func IsWarn() bool {
	return glog.IsWarn()
}

// Error is deprecated, use github.com/andriyg76/glog.Error
func Error(format string, a ...interface{}) {
	glog.Error(format, a...)
}

func IsError() bool {
	return glog.IsError()
}

func IsEnabled(logLevel LogLevel) bool {
	return glog.IsEnabled(logLevel)
}

// Panic is deprecated, use github.com/andriyg76/glog.Panic
func Panic(format string, a ...interface{}) {
	glog.Panic(format, a...)
}

// Fatal is deprecated, use github.com/andriyg76/glog.Fatal
func Fatal(format string, a ...interface{}) {
	glog.Fatal(format, a...)
}

// Fatal is deprecated, use github.com/andriyg76/glog.Fatal
func Log(level LogLevel, a string, objs ...interface{}) {
	glog.Log(level, a, objs...)
}

func OutputLevel(level LogLevel) Output {
	return glog.OutputLevel(level)
}
