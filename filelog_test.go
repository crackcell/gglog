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
 * @file filelog_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed Jul  1 17:27:24 2015
 *
 **/

package gglog

import (
	"testing"
)

func TestFileLoggerALl(t *testing.T) {
	f, err := NewFileLogger("./file.log", "filelog", NewStandardFormatter(),
		LOGLEVEL_ALL)
	if err != nil {
		t.Error(err)
		return
	}
	fl := f.(*FileLogger)
	fl.Debug("test debug")
	fl.Debugf("test debug %d", 1)
	fl.Info("test info")
	fl.Infof("test info %d", 1)
	fl.Warn("test warn")
	fl.Warnf("test warn %d", 1)

	err = fl.Rename("./file.rename.log")
	if err != nil {
		t.Error(err)
	}

	fl.Warn("file changed")
}
