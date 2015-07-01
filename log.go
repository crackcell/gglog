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
	"log"
	"math"
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
	logLevelMask int
	loggers      map[int]*log.Logger
	prefix       string
	writer       io.Writer
	format       Formatter
}

func NewLogger(writer io.Writer, prefix string, format Formatter,
	mask int) Logger {

	l := new(logger)
	l.format = format
	l.loggers = make(map[int]*log.Logger)
	if writer == nil {
		l.writer = &nullWriter{}
	} else {
		l.writer = writer
	}
	l.prefix = prefix
	l.SetLogLevel(mask)
	return l
}

//===================================================================
// Private
//===================================================================

func (l *logger) SetLogLevel(mask int) {
	l.logLevelMask = mask
	nullwriter := &nullWriter{}
	for i := 0; i < LOGLEVEL_COUNT; i++ {
		m := int(math.Exp2(float64(i)))
		flag := log.LstdFlags
		if mask&m != 0 {
			l.loggers[m] = log.New(l.writer, l.format.GetPrefix(l.prefix, m),
				flag)
		} else {
			l.loggers[m] = log.New(nullwriter, l.format.GetPrefix(l.prefix, m),
				flag)
		}
	}
}

func (l *logger) Debug(v ...interface{}) {
	l.loggers[LOGLEVEL_DEBUG].Print(append([]interface{}{l.format.GetPreFormat()}, v...)...)
}

func (l *logger) Debugf(format string, v ...interface{}) {
	l.loggers[LOGLEVEL_DEBUG].Printf(l.format.GetPreFormat()+format, v...)
}

func (l *logger) Info(v ...interface{}) {
	l.loggers[LOGLEVEL_INFO].Print(append([]interface{}{l.format.GetPreFormat()}, v...)...)
}

func (l *logger) Infof(format string, v ...interface{}) {
	l.loggers[LOGLEVEL_INFO].Printf(l.format.GetPreFormat()+format, v...)
}

func (l *logger) Warn(v ...interface{}) {
	l.loggers[LOGLEVEL_WARN].Print(append([]interface{}{l.format.GetPreFormat()}, v...)...)
}

func (l *logger) Warnf(format string, v ...interface{}) {
	l.loggers[LOGLEVEL_WARN].Printf(l.format.GetPreFormat()+format, v...)
}

func (l *logger) Fatal(v ...interface{}) {
	l.loggers[LOGLEVEL_FATAL].Fatal(append([]interface{}{l.format.GetPreFormat()}, v...)...)
}

func (l *logger) Fatalf(format string, v ...interface{}) {
	l.loggers[LOGLEVEL_FATAL].Fatalf(l.format.GetPreFormat()+format, v...)
}

type nullWriter struct{}

func (f *nullWriter) Write(p []byte) (int, error) {
	return len(p), nil
}
