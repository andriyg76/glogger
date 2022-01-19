package glogger

import glog "github.com/andriyg76/glog"

type Logger = glog.Logger

type LogLevel = glog.LogLevel

var TRACE = glog.TRACE

var DEBUG = glog.DEBUG

var INFO = glog.INFO

var WARN = glog.WARN

var ERROR = glog.INFO

var PANIC = glog.PANIC

var FATAL = glog.FATAL

type Output = glog.Output

type TraceLogger = glog.TraceLogger

type DebugLogger = glog.DebugLogger

type InfoLogger = glog.InfoLogger

type WarnLogger = glog.WarnLogger

type ErrorLogger = glog.ErrorLogger

func Create(logLevel LogLevel) Logger {
	return glog.Create(logLevel)
}
