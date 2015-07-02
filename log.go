/***************************************************************
 *
 * Copyright (c) 2014, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 * Logger.
 *
 * @file log.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed Oct 15 17:43:58 2014
 *
 **/

package gglog

import (
	"io"
	stdlog "log"
	"os"
)

//===================================================================
// Public APIs
//===================================================================

const (
	LOGLEVEL_FATAL = 1 << iota
	LOGLEVEL_WARN
	LOGLEVEL_INFO
	LOGLEVEL_DEBUG
	LOGLEVEL_ALL   = LOGLEVEL_DEBUG | LOGLEVEL_INFO | LOGLEVEL_WARN | LOGLEVEL_FATAL
	LOGLEVEL_COUNT = 4
)

type Logger interface {
	SetLogLevel(logLevelMask int)

	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
}

type logger struct {
	log          *stdlog.Logger
	logLevelMask int
	prefix       string
	writer       io.Writer
	formatter    Formatter
}

func NewLogger(writer io.Writer, prefix string, formatter Formatter,
	mask int) Logger {

	l := new(logger)
	l.formatter = formatter
	if writer == nil {
		l.writer = os.Stdout
	} else {
		l.writer = writer
	}
	l.log = stdlog.New(l.writer, "", 0)
	l.prefix = prefix
	l.SetLogLevel(mask)
	return l
}

//===================================================================
// Private
//===================================================================

func (l *logger) SetLogLevel(mask int) {
	l.logLevelMask = mask
}

func (l *logger) Debug(v ...interface{}) {
	l.writelog(LOGLEVEL_DEBUG, v...)
}

func (l *logger) Debugf(fmt string, v ...interface{}) {
	l.writelogf(LOGLEVEL_DEBUG, fmt, v...)
}

func (l *logger) Info(v ...interface{}) {
	l.writelog(LOGLEVEL_INFO, v...)
}

func (l *logger) Infof(fmt string, v ...interface{}) {
	l.writelogf(LOGLEVEL_INFO, fmt, v...)
}

func (l *logger) Warn(v ...interface{}) {
	l.writelog(LOGLEVEL_WARN, v...)
}

func (l *logger) Warnf(fmt string, v ...interface{}) {
	l.writelogf(LOGLEVEL_WARN, fmt, v...)
}

func (l *logger) Fatal(v ...interface{}) {
	l.writelog(LOGLEVEL_FATAL, v...)
}

func (l *logger) Fatalf(fmt string, v ...interface{}) {
	l.writelogf(LOGLEVEL_FATAL, fmt, v...)
}

func (l *logger) writelog(level int, v ...interface{}) {
	if l.logLevelMask&level != 0 {
		l.log.Print(l.formatter.GetMessage(level, v...)...)
	}
}

func (l *logger) writelogf(level int, fmt string, v ...interface{}) {
	if l.logLevelMask&level != 0 {
		l.log.Printf(l.formatter.GetFormat(level, fmt), v...)
	}
}
