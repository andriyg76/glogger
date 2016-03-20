package glogger

var _default = Create(WARN)

func SetLevel(logLevel LogLevel) {
	_default.SetLevel(logLevel)
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
