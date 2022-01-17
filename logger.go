package glogger

import (
	"fmt"
	"log"
	"os"
)

type LogLevel struct {
	prefix string
	weight int
}

var TRACE = LogLevel{
	prefix: "[trace]",
	weight: -2,
}

var DEBUG = LogLevel{
	prefix: "[debug]",
	weight: -1,
}

var INFO = LogLevel{
	prefix: "[info ]",
	weight: 0,
}

var WARN = LogLevel{
	prefix: "[warn ]",
	weight: 1,
}

var ERROR = LogLevel{
	prefix: "[error]",
	weight: 2,
}

var PANIC = LogLevel{
	prefix: "[trace]",
	weight: 2,
}

var FATAL = LogLevel{
	prefix: "[fatal]",
	weight: 2,
}

func (l LogLevel) String() string {
	return l.prefix
}

type Output interface {
	Printf(format string, a ...interface{})
}

type TraceLogger interface {
	Trace(format string, a ...interface{})
	TraceLogger() Output
	IsTrace() bool
}

type DebugLogger interface {
	Debug(format string, a ...interface{})
	IsDebug() bool
	DebugLogger() Output
}

type InfoLogger interface {
	Info(format string, a ...interface{})
	IsInfo() bool
}

type WarnLogger interface {
	Warn(format string, a ...interface{})
	IsWarn() bool
}

type ErrorLogger interface {
	Error(format string, a ...interface{})
	IsError() bool
}

type Logger interface {
	DebugLogger
	TraceLogger
	WarnLogger
	InfoLogger
	ErrorLogger

	// Log formats according to a format specifier
	Log(LogLevel LogLevel, format string, a ...interface{})
	IsEnabled(logLevel LogLevel) bool
	GetOutput(LogLevel LogLevel) Output

	Panic(format string, a ...interface{})
	Fatal(format string, a ...interface{})
}

type logger struct {
	logLevel LogLevel
	out      *log.Logger
	err      *log.Logger
}

func (l logger) IsDebug() bool {
	return l.IsEnabled(DEBUG)
}

func (l logger) IsTrace() bool {
	return l.IsEnabled(TRACE)
}

func (l logger) IsWarn() bool {
	return l.IsEnabled(WARN)
}

func (l logger) IsInfo() bool {
	return l.IsEnabled(INFO)
}

func (l logger) IsError() bool {
	return l.IsEnabled(ERROR)
}

func (l logger) IsEnabled(logLevel LogLevel) bool {
	return logLevel.weight >= l.logLevel.weight
}

//go:generate command stringer -type LogLevel

func Create(logLevel LogLevel) Logger {
	return logger{
		logLevel: logLevel,
		err:      _stderr,
		out:      _stdout,
	}
}

var _stdout = log.New(os.Stdout, "", log.LstdFlags)
var _stderr = log.New(os.Stderr, "", log.LstdFlags)

type dumbLogger struct{}

func (d dumbLogger) Printf(format string, objs ...interface{}) {}

var dumbLoggerInstance = dumbLogger{}

type loggerWithLevel struct {
	logLevel *LogLevel
	logger   *logger
}

func (l loggerWithLevel) Printf(format string, objs ...interface{}) {
	l.logger.Log(*l.logLevel, format, objs...)
}

func (l logger) GetOutput(logLevel LogLevel) Output {
	return loggerWithLevel{
		logLevel: &logLevel,
		logger:   &l,
	}
}

func (l logger) TraceLogger() Output {
	return l.GetOutput(TRACE)
}

func (l logger) Log(logLevel LogLevel, format string, objs ...interface{}) {
	var out Output
	if logLevel.weight < l.logLevel.weight {
		out = dumbLoggerInstance
	} else if logLevel.weight >= WARN.weight {
		out = l.err
	} else {
		out = l.out
	}

	out.Printf(logLevel.prefix+" "+format, objs...)

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
	return l.GetOutput(DEBUG)
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
