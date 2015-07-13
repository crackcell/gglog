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
	"sync"
)

//===================================================================
// Public APIs
//===================================================================

func NewFileLogger(path string, prefix string, formatter Formatter,
	mask int) (*FileLogger, error) {

	l := new(FileLogger)
	l.mu = new(sync.Mutex)
	var err error
	l.file, err = l.Open(path)
	if err != nil {
		return nil, err
	}
	l.log = NewLogger(l.file, prefix, formatter, mask)
	l.prefix = prefix
	l.formatter = formatter
	l.mask = mask
	return l, nil
}

//===================================================================
// Private
//===================================================================

type FileLogger struct {
	mu        *sync.Mutex
	file      *os.File
	prefix    string
	formatter Formatter
	mask      int
	path      string
	log       Logger
}

func (l *FileLogger) Open(path string) (*os.File, error) {
	var f *os.File
	var err error
	if f, err = os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666); os.IsNotExist(err) {
		f, err = os.Create(path)
	}
	if err != nil {
		return nil, err
	}
	l.path = path
	return f, nil
}

func (l *FileLogger) GetPath() string {
	return l.path
}

func (l *FileLogger) Close() error {
	return l.file.Close()
}

func (l *FileLogger) Rename(newPath string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	var err error

	err = l.Close()
	if err != nil {
		return err
	}

	err = os.Rename(l.path, newPath)
	if err != nil {
		return err
	}

	l.file, err = l.Open(newPath)
	if err != nil {
		return err
	}

	l.log = NewLogger(l.file, l.prefix, l.formatter, l.mask)
	l.path = newPath
	return nil
}

func (l *FileLogger) SetLogLevel(mask int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.log.SetLogLevel(mask)
}

func (l *FileLogger) Debug(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.log.Debug(v...)
}

func (l *FileLogger) Debugf(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.log.Debugf(format, v...)
}

func (l *FileLogger) Info(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.log.Info(v...)
}

func (l *FileLogger) Infof(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.log.Infof(format, v...)
}

func (l *FileLogger) Warn(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.log.Warn(v...)
}

func (l *FileLogger) Warnf(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.log.Warnf(format, v...)
}

func (l *FileLogger) Fatal(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.log.Fatal(v...)
}

func (l *FileLogger) Fatalf(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.log.Fatalf(format, v...)
}
