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
 * @file rollfilelog.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Thu Jul  2 15:53:02 2015
 *
 **/

package gglog

import (
	"os"
)

//===================================================================
// Public APIs
//===================================================================

type FileWithLock struct {
	mu   sync.Mutex
	file *os.File
}

type FileQueue struct {
}

type LogRoller interface {
	Roll() error
}

func NewRollFileLogger(path string, prefix string, format Formatter,
	mask int) (Logger, error) {

	l := new(rollFileLogger)

}

//===================================================================
// Private
//===================================================================

type rollFileLogger struct {
	files []*os.File
}
