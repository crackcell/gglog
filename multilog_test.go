/***************************************************************
 *
 * Copyright (c) 2014, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 * Test for multi.
 *
 * @file multi_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed Oct 15 12:38:33 2014
 *
 **/

package gglog

import (
	"os"
	"testing"
)

var m Logger

func TestNewMulti(t *testing.T) {
	m = NewMultiLogger(
		NewLogger(os.Stdout, "log", NewStandardFormatter(),
			LOGLEVEL_DEBUG|LOGLEVEL_INFO|LOGLEVEL_WARN|LOGLEVEL_FATAL),
		NewLogger(os.Stderr, "log.wf", NewStandardFormatter(),
			LOGLEVEL_WARN|LOGLEVEL_FATAL))
}

func TestMultiDebug(t *testing.T) {
	m.Debug("test debug")
}

func TestMultiDebugf(t *testing.T) {
	m.Debugf("test debug %d", 1)
}

func TestMultiInfo(t *testing.T) {
	m.Info("test info")
}
func TestMultiInfof(t *testing.T) {
	m.Infof("test info %d", 1)
}

func TestMultiWarn(t *testing.T) {
	m.Warn("test warn")
}
func TestMultiWarnf(t *testing.T) {
	m.Warnf("test warn %d", 1)
}

/*
func TestMultiFatal(t *testing.T) {
	m.Fatal("test fatal")
}
*/
