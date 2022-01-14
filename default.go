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

func Debug(format string, a ...interface{}) {
	_default.Debug(format, a...)
}

func Info(format string, a ...interface{}) {
	_default.Info(format, a...)
}

func Warn(format string, a ...interface{}) {
	_default.Warn(format, a...)
}

func Error(format string, a ...interface{}) {
	_default.Error(format, a...)
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
