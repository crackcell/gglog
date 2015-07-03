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
	"os"
	"sync"
)

//===================================================================
// Public APIs
//===================================================================

type FileWithLock struct {
	file *os.File
}

type FileQueue struct {
}

type LogRoller interface {
	Roll() error
}

func NewRollFileLogger(path string, prefix string, format Formatter,
	mask int) (Logger, error) {

	l := new(rollFileLogger)
	l.mu = new(sync.Mutex)

	return l, nil
}

//===================================================================
// Private
//===================================================================

type rollFileLogger struct {
	mu    *sync.Mutex
	files []*os.File
}

func (l *rollFileLogger) SetLogLevel(mask int) {
}

func (l *rollFileLogger) Debug(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
}

func (l *rollFileLogger) Debugf(fmt string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
}

func (l *rollFileLogger) Info(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
}

func (l *rollFileLogger) Infof(fmt string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
}

func (l *rollFileLogger) Warn(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
}

func (l *rollFileLogger) Warnf(fmt string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
}

func (l *rollFileLogger) Fatal(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
}

func (l *rollFileLogger) Fatalf(fmt string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
}
