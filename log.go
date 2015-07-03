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
	"fmt"
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
	l.output(LOGLEVEL_DEBUG, fmt.Sprint(v...))
}

func (l *logger) Debugf(format string, v ...interface{}) {
	l.output(LOGLEVEL_DEBUG, fmt.Sprintf(format, v...))
}

func (l *logger) Info(v ...interface{}) {
	l.output(LOGLEVEL_INFO, fmt.Sprint(v...))
}

func (l *logger) Infof(format string, v ...interface{}) {
	l.output(LOGLEVEL_INFO, fmt.Sprintf(format, v...))
}

func (l *logger) Warn(v ...interface{}) {
	l.output(LOGLEVEL_WARN, fmt.Sprint(v...))
}

func (l *logger) Warnf(format string, v ...interface{}) {
	l.output(LOGLEVEL_WARN, fmt.Sprintf(format, v...))
}

func (l *logger) Fatal(v ...interface{}) {
	l.output(LOGLEVEL_FATAL, fmt.Sprint(v...))
}

func (l *logger) Fatalf(format string, v ...interface{}) {
	l.output(LOGLEVEL_FATAL, fmt.Sprintf(format, v...))
}

func (l *logger) output(level int, s string) {
	if l.logLevelMask&level != 0 {
		l.log.Print(l.formatter.GetHeader(level) + s)
	}
}
