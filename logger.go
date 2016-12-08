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
	INFO LogLevel = 0
	WARN LogLevel = 1
	ERROR LogLevel = 2
	PANIC LogLevel = 3
	FATAL LogLevel = 4
)

type Logger interface {

	Log(LogLevel LogLevel, format string, objs ...interface{})
	Logger(LogLevel LogLevel) *log.Logger
	Trace(format string, objs ...interface{})
	TraceLogger() *log.Logger
	Debug(format string, objs ...interface{})
//	DebugLogger() *log.Logger
	Info(format string, objs ...interface{})
//	InfoLogger() *log.Logger
	Warn(format string, objs ...interface{})
//	WarnLogger() *log.Logger
	Error(format string, objs ...interface{})
//	ErrorLogger() *log.Logger
	Panic(format string, objs ...interface{})
	Fatal(format string, objs ...interface{})
	SetLevel(LogLevel)
}

type logger struct {
	logLevel LogLevel
}

//go:generate command stringer -type LogLevel


func Create(logLevel LogLevel) Logger {
	return &logger{logLevel}
}

var stdout = log.New(os.Stdout, "", log.LstdFlags)
var stderr = log.New(os.Stderr, "", log.LstdFlags)

func (l logger) Logger(logLevel LogLevel) *log.Logger {
	if logLevel > l.logLevel {
		return nil
	}

	if logLevel >= WARN {
		return stderr
	} else {
		return stdout
	}
}

func (l logger) TraceLogger() *log.Logger {
	return l.Logger(TRACE)
}

func (l logger) Log(logLevel LogLevel, format string, objs ...interface{}) {
//	fmt.Fprintf(os.Stderr, "Logger level is %s print level %s\n", l.logLevel, logLevel)
	if logLevel >= l.logLevel {
		var out *log.Logger
//		out_name := ""
		if logLevel >= WARN {
			out = stderr
//			out_name = "stderr"
		} else {
			out = stdout
//			out_name = "stdout"
		}
		// fmt.Fprintf(os.Stderr, "Will log to %s \n", out_name)
		out.Print(logLevel, " ", fmt.Sprintf(format, objs...))
	} else {
		// fmt.Fprintf(os.Stderr, "Will not write log\n")
	}

	if logLevel == PANIC {
		panic(fmt.Sprintf(format, objs...))
	}
	if logLevel == FATAL {
		os.Exit(1)
	}
}

func (l *logger) Debug(format string, objs ...interface{}) {
	l.Log(DEBUG, format, objs...)
}

func (l *logger) Trace(format string, objs ...interface{}) {
	l.Log(TRACE, format, objs...)
}

func (l *logger) Info(format string, objs ...interface{}) {
	l.Log(INFO, format, objs...)
}

func (l *logger) Warn(format string, objs ...interface{}) {
	l.Log(WARN, format, objs...)
}

func (l *logger) Error(format string, objs ...interface{}) {
	l.Log(ERROR, format, objs...)
}

func (l *logger) Panic(format string, objs ...interface{}) {
	l.Log(PANIC, format, objs...)
}

func (l *logger) Fatal(format string, objs ...interface{}) {
	l.Log(FATAL, format, objs...)
}

func (l *logger) SetLevel(logLevel LogLevel) {
	// if l.logLevel != logLevel { fmt.Fprintf(os.Stderr, "Set logger from %s level to %s\n", l.logLevel, logLevel) }
	l.logLevel = logLevel
}