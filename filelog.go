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
 * @file filelog.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed Jul  1 17:02:45 2015
 *
 **/

package gglog

import (
	"os"
)

//===================================================================
// Public APIs
//===================================================================

func NewFileLogger(path string, prefix string, format Formatter,
	mask int) (Logger, error) {

	l := new(fileLogger)
	var f *os.File
	var err error
	if f, err = os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666); os.IsNotExist(err) {
		f, err = os.Create(path)
	}
	if err != nil {
		return nil, err
	}
	l.log = NewLogger(f, prefix, format, mask)
	return l, nil
}

//===================================================================
// Private
//===================================================================

type fileLogger struct {
	log Logger
}

func (l *fileLogger) SetLogLevel(mask int) {
	l.log.SetLogLevel(mask)
}

func (l *fileLogger) Debug(v ...interface{}) {
	l.log.Debug(v...)
}

func (l *fileLogger) Debugf(format string, v ...interface{}) {
	l.log.Debugf(format, v...)
}

func (l *fileLogger) Info(v ...interface{}) {
	l.log.Info(v...)
}

func (l *fileLogger) Infof(format string, v ...interface{}) {
	l.log.Infof(format, v...)
}

func (l *fileLogger) Warn(v ...interface{}) {
	l.log.Warn(v...)
}

func (l *fileLogger) Warnf(format string, v ...interface{}) {
	l.log.Warnf(format, v...)
}

func (l *fileLogger) Fatal(v ...interface{}) {
	l.log.Fatal(v...)
}

func (l *fileLogger) Fatalf(format string, v ...interface{}) {
	l.log.Fatalf(format, v...)
}
