package glogger

import (
	"fmt"
	"os"
	"log"
)

type LogLevel byte

const (
	TRACE LogLevel = iota
	DEBUG
	INFO
	WARN
	ERROR
)

type Logger interface {
	Debug(format string, objs ...interface{})
	Trace(format string, objs ...interface{})
	Info(format string, objs ...interface{})
	Warn(format string, objs ...interface{})
	Error(format string, objs ...interface{})
	SetLevel(LogLevel)
}

type logger struct {
	LogLevel LogLevel
}

func Create(logLevel LogLevel) Logger {
	return logger{logLevel}
}

var stdout = log.New(os.Stdout, "", log.LstdFlags)
var stderr = log.New(os.Stderr, "", log.LstdFlags)

func (l logger) log(logLevel LogLevel, out *log.Logger, format string, objs ...interface{}) {
	if logLevel >= l.LogLevel {
		out.Print(logLevel, " ", fmt.Sprintf(format, objs...))
	}
}

func (l logger) Debug(format string, objs ...interface{}) {
	l.log(DEBUG, stdout, format, objs...)
}

func (l logger) Trace(format string, objs ...interface{}) {
	l.log(TRACE, stdout, format, objs...)
}

func (l logger) Info(format string, objs ...interface{}) {
	l.log(INFO, stdout, format, objs...)
}

func (l logger) Warn(format string, objs ...interface{}) {
	l.log(WARN, stderr, format, objs...)
}

func (l logger) Error(format string, objs ...interface{}) {
	l.log(ERROR, stderr, format, objs...)
}

func (l logger) SetLevel(logLevel LogLevel) {
	l.LogLevel = logLevel
}