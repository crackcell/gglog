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

	l := new(FileLogger)
	var err error
	if l.file, err = os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666); os.IsNotExist(err) {
		l.file, err = os.Create(path)
	}
	if err != nil {
		return nil, err
	}
	l.log = NewLogger(l.file, prefix, format, mask)
	return l, nil
}

//===================================================================
// Private
//===================================================================

type FileLogger struct {
	file *os.File
	log  Logger
}

func (l *FileLogger) SetLogLevel(mask int) {
	l.log.SetLogLevel(mask)
}

func (l *FileLogger) Debug(v ...interface{}) {
	l.log.Debug(v...)
}

func (l *FileLogger) Debugf(format string, v ...interface{}) {
	l.log.Debugf(format, v...)
}

func (l *FileLogger) Info(v ...interface{}) {
	l.log.Info(v...)
}

func (l *FileLogger) Infof(format string, v ...interface{}) {
	l.log.Infof(format, v...)
}

func (l *FileLogger) Warn(v ...interface{}) {
	l.log.Warn(v...)
}

func (l *FileLogger) Warnf(format string, v ...interface{}) {
	l.log.Warnf(format, v...)
}

func (l *FileLogger) Fatal(v ...interface{}) {
	l.log.Fatal(v...)
}

func (l *FileLogger) Fatalf(format string, v ...interface{}) {
	l.log.Fatalf(format, v...)
}
