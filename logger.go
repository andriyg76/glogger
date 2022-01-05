package glogger

import (
	"fmt"
	"log"
	"os"
)

type LogLevel int8

const (
	TRACE LogLevel = -2
	DEBUG LogLevel = -1
	INFO  LogLevel = 0
	WARN  LogLevel = 1
	ERROR LogLevel = 2
	PANIC LogLevel = 3
	FATAL LogLevel = 4
)

//go:generate stringer -type LogLevel

type Output interface {
	Printf(format string, objs ...interface{})
}

type TraceLogger interface {
	Trace(format string, objs ...interface{})
	TraceLogger() Output
}

type DebugLogger interface {
	Debug(format string, objs ...interface{})
	DebugLogger() Output
}

type InfoLogger interface {
	Info(format string, objs ...interface{})
}

type WarnLogger interface {
	Warn(format string, objs ...interface{})
}

type ErrorLogger interface {
	Error(format string, objs ...interface{})
}

type Logger interface {
	DebugLogger
	TraceLogger
	WarnLogger
	InfoLogger
	ErrorLogger

	Log(LogLevel LogLevel, format string, objs ...interface{})
	Logger(LogLevel LogLevel) Output

	Panic(format string, objs ...interface{})
	Fatal(format string, objs ...interface{})
}

type logger struct {
	logLevel LogLevel
}

//go:generate command stringer -type LogLevel

func Create(logLevel LogLevel) Logger {
	return logger{logLevel}
}

var stdout = log.New(os.Stdout, "", log.LstdFlags)
var stderr = log.New(os.Stderr, "", log.LstdFlags)

type dumbLogger struct{}

func (d dumbLogger) Printf(format string, objs ...interface{}) {}

var dumbLoggerInstance = dumbLogger{}

func (l logger) Logger(logLevel LogLevel) Output {
	if logLevel < l.logLevel {
		return dumbLoggerInstance
	}

	if logLevel >= WARN {
		return stderr
	} else {
		return stdout
	}
}

func (l logger) TraceLogger() Output {
	return l.Logger(TRACE)
}

func (l logger) Log(logLevel LogLevel, format string, objs ...interface{}) {
	l.Logger(logLevel).Printf(format, objs...)

	if logLevel == PANIC {
		panic(fmt.Sprintf(format, objs...))
	}
	if logLevel == FATAL {
		os.Exit(1)
	}
}

func (l logger) Debug(format string, objs ...interface{}) {
	l.Log(DEBUG, format, objs...)
}

func (l logger) DebugLogger() Output {
	return l.Logger(DEBUG)
}

func (l logger) Trace(format string, objs ...interface{}) {
	l.Log(TRACE, format, objs...)
}

func (l logger) Info(format string, objs ...interface{}) {
	l.Log(INFO, format, objs...)
}

func (l logger) Warn(format string, objs ...interface{}) {
	l.Log(WARN, format, objs...)
}

func (l logger) Error(format string, objs ...interface{}) {
	l.Log(ERROR, format, objs...)
}

func (l logger) Panic(format string, objs ...interface{}) {
	l.Log(PANIC, format, objs...)
}

func (l logger) Fatal(format string, objs ...interface{}) {
	l.Log(FATAL, format, objs...)
}
