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
 * @file baidulog_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed Jul  1 16:10:26 2015
 *
 **/

package gglog

import (
	"testing"
)

var blog Logger

func TestNewBaiduLogger(t *testing.T) {
	blog = NewBaiduLogger("testmodule", "./", LOGLEVEL_ALL)
}

func TestBaiduLoggerAll(t *testing.T) {
	blog.Debug("test debug")
	blog.Debugf("test debug %d", 1)
	blog.Info("test info")
	blog.Infof("test info %d", 1)
	blog.Warn("test warn")
	blog.Warnf("test warn %d", 1)
}
