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
	out      *log.Logger
	err      *log.Logger
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

type prefixOutput struct {
	prefix string
	out    Output
}

func (p prefixOutput) Printf(format string, objs ...interface{}) {
	p.out.Printf(p.prefix+" "+format, objs...)
}

func (l logger) Logger(logLevel LogLevel) Output {
	var out Output
	if logLevel.weight < l.logLevel.weight {
		out = dumbLoggerInstance
	} else if logLevel.weight >= WARN.weight {
		out = l.err
	} else {
		out = l.out
	}
	return prefixOutput{
		prefix: logLevel.prefix,
		out:    out,
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
