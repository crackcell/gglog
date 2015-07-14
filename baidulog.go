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
	outlog *FileLogger
	errlog *FileLogger
}

func (l *baiduLogger) SetLogLevel(mask int) {
	l.outlog.SetLogLevel(mask)
	l.errlog.SetLogLevel(mask)
}

func (l *baiduLogger) Debug(v ...interface{}) {
	l.Output(LOGLEVEL_DEBUG, fmt.Sprint(v...))
}

func (l *baiduLogger) Debugf(format string, v ...interface{}) {
	l.Output(LOGLEVEL_DEBUG, fmt.Sprintf(format, v...))
}

func (l *baiduLogger) Info(v ...interface{}) {
	l.Output(LOGLEVEL_INFO, fmt.Sprint(v...))
}

func (l *baiduLogger) Infof(format string, v ...interface{}) {
	l.Output(LOGLEVEL_INFO, fmt.Sprintf(format, v...))
}

func (l *baiduLogger) Warn(v ...interface{}) {
	l.Output(LOGLEVEL_WARN, fmt.Sprint(v...))
}

func (l *baiduLogger) Warnf(format string, v ...interface{}) {
	l.Output(LOGLEVEL_WARN, fmt.Sprintf(format, v...))
}

func (l *baiduLogger) Fatal(v ...interface{}) {
	l.Output(LOGLEVEL_FATAL, fmt.Sprint(v...))
}

func (l *baiduLogger) Fatalf(format string, v ...interface{}) {
	l.Output(LOGLEVEL_FATAL, fmt.Sprintf(format, v...))
}

func (l *baiduLogger) Output(level int, s string) {
	switch level {
	case LOGLEVEL_FATAL:
		l.errlog.Fatal(s)
	case LOGLEVEL_WARN:
		l.errlog.Warn(s)
	case LOGLEVEL_INFO:
		l.outlog.Info(s)
	case LOGLEVEL_DEBUG:
		l.outlog.Debug(s)
	}
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
