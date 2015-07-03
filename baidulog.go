/***************************************************************
 *
 * Copyright (c) 2015, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 * Baidu style logger
 *
 * @file baidulog.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed Jul  1 14:48:23 2015
 *
 **/

package gglog

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

//===================================================================
// Public APIs
//===================================================================

func NewBaiduLogger(module string, path string, logLevelMask int) (Logger, error) {
	l := new(baiduLogger)
	l.module = module
	var err error
	l.outlog, err = NewFileLogger(path+".log", "", NewBaiduFormatter(), logLevelMask)
	if err != nil {
		return nil, err
	}
	l.errlog, err = NewFileLogger(path+".log.wf", "", NewBaiduFormatter(), logLevelMask)
	if err != nil {
		return nil, err
	}
	return l, nil
}

//===================================================================
// Private
//===================================================================

type baiduLogger struct {
	module string
	outlog Logger
	errlog Logger
}

func (l *baiduLogger) SetLogLevel(mask int) {
	l.outlog.SetLogLevel(mask)
	l.errlog.SetLogLevel(mask)
}

func (l *baiduLogger) Debug(v ...interface{}) {
	l.outlog.Debug(append(
		[]interface{}{l.module, " * ", getCallerInfo(), " "}, v...)...)
}

func (l *baiduLogger) Debugf(format string, v ...interface{}) {
	l.outlog.Debugf(l.module+" * "+getCallerInfo()+" "+format, v...)
}

func (l *baiduLogger) Info(v ...interface{}) {
	l.outlog.Info(append(
		[]interface{}{l.module, " * "}, v...)...)
}

func (l *baiduLogger) Infof(format string, v ...interface{}) {
	l.outlog.Infof(l.module+" * "+format, v...)
}

func (l *baiduLogger) Warn(v ...interface{}) {
	l.errlog.Warn(append(
		[]interface{}{l.module, " * ", getCallerInfo(), " "}, v...)...)
}

func (l *baiduLogger) Warnf(format string, v ...interface{}) {
	l.errlog.Warnf(l.module+" * "+getCallerInfo()+" "+format, v...)
}

func (l *baiduLogger) Fatal(v ...interface{}) {
	l.errlog.Fatal(append(
		[]interface{}{l.module, " * ", getCallerInfo(), " "}, v...)...)
}

func (l *baiduLogger) Fatalf(format string, v ...interface{}) {
	l.errlog.Fatalf(l.module+" * "+getCallerInfo()+" "+format, v...)
}

type BaiduFormatter struct {
	logLevelToName map[int]string
}

func NewBaiduFormatter() Formatter {
	f := new(BaiduFormatter)
	f.logLevelToName = make(map[int]string)
	f.logLevelToName[LOGLEVEL_DEBUG] = "DEBUG"
	f.logLevelToName[LOGLEVEL_INFO] = "INFO"
	f.logLevelToName[LOGLEVEL_WARN] = "WARNING"
	f.logLevelToName[LOGLEVEL_FATAL] = "FATAL"
	return f
}

func (f *BaiduFormatter) GetHeader(level int) string {
	return f.logLevelToName[level] + ": " +
		time.Now().Format("2006-01-02 15:04:05") + ": "
}

func getCallerInfo() string {
	pc, fn, line, _ := runtime.Caller(2)
	pcs := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	fns := strings.Split(fn, "/")
	return fmt.Sprintf("[%s:%d]%s", fns[len(fns)-1], line, pcs[len(pcs)-1])
}
