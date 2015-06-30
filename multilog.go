/***************************************************************
 *
 * Copyright (c) 2014, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 *
 *
 * @file multilog.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed Oct 15 11:47:28 2014
 *
 **/

package gglog

//===================================================================
// Public APIs
//===================================================================

func NewMultiLogger(loggers ...Logger) Logger {
	return &multiLogger{loggers}
}

func (self *multiLogger) Debug(v ...interface{}) {
	for _, l := range self.loggers {
		if l != nil {
			l.Debug(v...)
		}
	}
}

func (self *multiLogger) Debugf(format string, v ...interface{}) {
	for _, l := range self.loggers {
		if l != nil {
			l.Debugf(format, v...)
		}
	}
}

func (self *multiLogger) Info(v ...interface{}) {
	for _, l := range self.loggers {
		if l != nil {
			l.Info(v...)
		}
	}
}

func (self *multiLogger) Infof(format string, v ...interface{}) {
	for _, l := range self.loggers {
		if l != nil {
			l.Infof(format, v...)
		}
	}
}

func (self *multiLogger) Warn(v ...interface{}) {
	for _, l := range self.loggers {
		if l != nil {
			l.Warn(v...)
		}
	}
}

func (self *multiLogger) Warnf(format string, v ...interface{}) {
	for _, l := range self.loggers {
		if l != nil {
			l.Warnf(format, v...)
		}
	}
}

func (self *multiLogger) Fatal(v ...interface{}) {
	for _, l := range self.loggers {
		if l != nil {
			l.Fatal(v...)
		}
	}
}

func (self *multiLogger) Fatalf(format string, v ...interface{}) {
	for _, l := range self.loggers {
		if l != nil {
			l.Fatalf(format, v...)
		}
	}
}

//===================================================================
// Private
//===================================================================

type multiLogger struct {
	loggers []Logger
}
