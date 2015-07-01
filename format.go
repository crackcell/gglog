/***************************************************************
 *
 * Copyright (c) 2015, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 * Log format
 *
 * @file format.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed Jul  1 15:04:36 2015
 *
 **/

package gglog

import (
// "fmt"
)

//===================================================================
// Public APIs
//===================================================================

type Formatter interface {
	GetPrefix(userPrefix string, level int) string
	GetPreFormat() string
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

func (f *StandardFormatter) GetPrefix(userPrefix string, level int) string {
	return userPrefix + " " + f.logLevelToName[level] + " "
}

func (f *StandardFormatter) GetPreFormat() string {
	return ""
}

//===================================================================
// Private
//===================================================================
