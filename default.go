package glogger

var _default = Create(INFO)

func Default() Logger {
	return _default
}

func SetLevel(logLevel LogLevel) {
	_default = Create(logLevel)
}

func Trace(format string, a ...interface{}) {
	_default.Trace(format, a...)
}

func IsTrace() bool {
	return _default.IsTrace()
}

func Debug(format string, a ...interface{}) {
	_default.Debug(format, a...)
}

func IsDebug() bool {
	return _default.IsDebug()
}

func Info(format string, a ...interface{}) {
	_default.Info(format, a...)
}

func IsInfo() bool {
	return _default.IsInfo()
}

func Warn(format string, a ...interface{}) {
	_default.Warn(format, a...)
}

func IsWarn() bool {
	return _default.IsWarn()
}

func Error(format string, a ...interface{}) {
	_default.Error(format, a...)
}

func IsError() bool {
	return _default.IsError()
}

func IsEnabled(logLevel LogLevel) bool {
	return _default.IsEnabled(logLevel)
}
func Panic(format string, a ...interface{}) {
	_default.Panic(format, a...)
}

func Fatal(format string, a ...interface{}) {
	_default.Fatal(format, a...)
}

func Log(level LogLevel, a string, objs ...interface{}) {
	_default.Log(level, a, objs...)
}

func OutputLevel(level LogLevel) Output {
	return _default.GetOutput(level)
}
