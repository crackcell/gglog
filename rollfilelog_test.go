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
 * @file rollfilelog_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Mon Jul 13 14:59:41 2015
 *
 **/

package gglog

import (
	"testing"
)

func TestRollFileLoggerAll(t *testing.T) {
	l, err := NewRollFileLogger("./rollfile.log", "TestRollFileLog",
		NewStandardFormatter(), LOGLEVEL_ALL)
	l.SetRollSize(1 * KB)
	if err != nil {
		t.Error(err)
		return
	}
	l.Debug("test debug")
	l.Debugf("test debug %d", 1)
	l.Info("test info")
	l.Infof("test info %d", 1)
	l.Warn("test warn")
	l.Warnf("test warn %d", 1)
	l.Fatal("test fatal")
	l.Fatalf("test fatal %d", 1)
}
