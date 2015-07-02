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

func TestLoggerAll(t *testing.T) {
	l := NewLogger(os.Stderr, "TestLogger", NewStandardFormatter(), LOGLEVEL_ALL)
	l.Debug("test debug")
	l.Debugf("test debug %d", 1)
	l.Info("test info")
	l.Infof("test info %d", 1)
	l.Warn("test warn")
	l.Warnf("test warn %d", 1)
	l.Fatal("test fatal")
	l.Fatalf("test fatal %d", 1)
}
