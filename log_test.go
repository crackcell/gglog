/***************************************************************
 *
 * Copyright (c) 2014, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 * Test for log.
 *
 * @file log_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed Oct 15 12:25:48 2014
 *
 **/

package gglog

import (
	"os"
	"testing"
)

var l Logger

func TestNewLogger(t *testing.T) {
	l = NewLogger(os.Stderr, "TestLogger", LOGLEVEL_DEBUG)
}

func TestDebug(t *testing.T) {
	l.Debug("test debug")
	l.Debugf("test debug %d", 1)
}

func TestInfo(t *testing.T) {
	l.Info("test info")
	l.Infof("test info %d", 1)
}

func TestWarn(t *testing.T) {
	l.Warn("test warn")
	l.Warnf("test warn %d", 1)
}

/*
func TestFatal(t *testing.T) {
	l.Fatal("test fatal")
	l.Fatalf("test fatal %d", 1)
}
*/
