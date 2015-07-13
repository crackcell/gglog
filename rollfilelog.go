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

import ()

//===================================================================
// Public APIs
//===================================================================

func NewRollFileLogger(path string, prefix string, formatter Formatter,
	mask int) (Logger, error) {
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
}

func (l *RollFileLogger) SetLogLevel(mask int) {
}

func (l *RollFileLogger) Debug(v ...interface{}) {
}

func (l *RollFileLogger) Debugf(fmt string, v ...interface{}) {
}

func (l *RollFileLogger) Info(v ...interface{}) {
}

func (l *RollFileLogger) Infof(fmt string, v ...interface{}) {
}

func (l *RollFileLogger) Warn(v ...interface{}) {
}

func (l *RollFileLogger) Warnf(fmt string, v ...interface{}) {
}

func (l *RollFileLogger) Fatal(v ...interface{}) {
}

func (l *RollFileLogger) Fatalf(fmt string, v ...interface{}) {
}
