/***************************************************************
 *
 * Copyright (c) 2015, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 *
 *
 * @file rollfilelog.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Thu Jul  2 15:53:02 2015
 *
 **/

package gglog

import (
	"fmt"
)

//===================================================================
// Public APIs
//===================================================================

const (
	KB = 1024
	MB = 1024 * 1024
	GB = 1024 * 1024 * 1024
)

func NewRollFileLogger(path string, prefix string, formatter Formatter,
	mask int) (*RollFileLogger, error) {
	l := new(RollFileLogger)
	if fl, err := NewFileLogger(path, prefix, formatter, mask); err != nil {
		return nil, err
	} else {
		l.loggers = []*FileLogger{fl}
	}
	return l, nil
}

//===================================================================
// Private
//===================================================================

type RollFileLogger struct {
	loggers []*FileLogger
	mask    int
	size    int
	maxSize int
}

func (l *RollFileLogger) SetRollSize(size int) {
	l.maxSize = size
}

func (l *RollFileLogger) SetLogLevel(mask int) {
	l.mask = mask
}

func (l *RollFileLogger) Debug(v ...interface{}) {
	l.loggers[0].Output(LOGLEVEL_DEBUG, fmt.Sprint(v...))
}

func (l *RollFileLogger) Debugf(format string, v ...interface{}) {
	l.loggers[0].Output(LOGLEVEL_DEBUG, fmt.Sprintf(format, v...))
}

func (l *RollFileLogger) Info(v ...interface{}) {
	l.loggers[0].Output(LOGLEVEL_INFO, fmt.Sprint(v...))
}

func (l *RollFileLogger) Infof(format string, v ...interface{}) {
	l.loggers[0].Output(LOGLEVEL_INFO, fmt.Sprintf(format, v...))
}

func (l *RollFileLogger) Warn(v ...interface{}) {
	l.loggers[0].Output(LOGLEVEL_WARN, fmt.Sprint(v...))
}

func (l *RollFileLogger) Warnf(format string, v ...interface{}) {
	l.loggers[0].Output(LOGLEVEL_WARN, fmt.Sprintf(format, v...))
}

func (l *RollFileLogger) Fatal(v ...interface{}) {
	l.loggers[0].Output(LOGLEVEL_FATAL, fmt.Sprint(v...))
}

func (l *RollFileLogger) Fatalf(format string, v ...interface{}) {
	l.loggers[0].Output(LOGLEVEL_FATAL, fmt.Sprintf(format, v...))
}

func (l *RollFileLogger) Output(level int, s string) {
	if l.mask&level != 0 {
		l.loggers[0].Output(level, s)
	}
}
