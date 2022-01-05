package glogger

var _default = Create(WARN)

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
	_default.Panic(format, objs)
}

func Fatal(format string, objs ...interface{}) {
	_default.Fatal(format, objs)
}
