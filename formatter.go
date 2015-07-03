/***************************************************************
 *
 * Copyright (c) 2015, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 * Log formatter
 *
 * @file format.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed Jul  1 15:04:36 2015
 *
 **/

package gglog

import (
	"time"
)

//===================================================================
// Public APIs
//===================================================================

type Formatter interface {
	GetHeader(level int) string
}

type StandardFormatter struct {
	logLevelToName map[int]string
}

func NewStandardFormatter() Formatter {
	f := new(StandardFormatter)
	f.logLevelToName = make(map[int]string)
	f.logLevelToName[LOGLEVEL_DEBUG] = "DEBUG"
	f.logLevelToName[LOGLEVEL_INFO] = "INFO"
	f.logLevelToName[LOGLEVEL_WARN] = "WARNING"
	f.logLevelToName[LOGLEVEL_FATAL] = "FATAL"
	return f
}

func (f *StandardFormatter) GetHeader(level int) string {
	return f.logLevelToName[level] + " " +
		time.Now().Format("2006-01-02 15:04:05") + " "
}

//===================================================================
// Private
//===================================================================
