package glogger

var _default = Create(INFO)

func Default() Logger {
	return _default
}

func SetLevel(logLevel LogLevel) {
	_default = Create(logLevel)
}

func Trace(format string, objs ...interface{}) {
	_default.Trace(format, objs...)
}

func Debug(format string, objs ...interface{}) {
	_default.Debug(format, objs...)
}

func Info(format string, objs ...interface{}) {
	_default.Info(format, objs...)
}

func Warn(format string, objs ...interface{}) {
	_default.Warn(format, objs)
}

func Error(format string, objs ...interface{}) {
	_default.Error(format, objs)
}

func Panic(format string, objs ...interface{}) {
	_default.Panic(format, objs...)
}

func Fatal(format string, objs ...interface{}) {
	_default.Fatal(format, objs...)
}

func Log(level LogLevel, format string, objs ...interface{}) {
	_default.Log(level, format, objs...)
}

func OutputLevel(level LogLevel) Output {
	return _default.GetOutput(level)
}
