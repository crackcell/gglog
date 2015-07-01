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

func (m *multiLogger) SetLogLevel(logLevelMask int) {}

func (m *multiLogger) Debug(v ...interface{}) {
	for _, l := range m.loggers {
		if l != nil {
			l.Debug(v...)
		}
	}
}

func (m *multiLogger) Debugf(format string, v ...interface{}) {
	for _, l := range m.loggers {
		if l != nil {
			l.Debugf(format, v...)
		}
	}
}

func (m *multiLogger) Info(v ...interface{}) {
	for _, l := range m.loggers {
		if l != nil {
			l.Info(v...)
		}
	}
}

func (m *multiLogger) Infof(format string, v ...interface{}) {
	for _, l := range m.loggers {
		if l != nil {
			l.Infof(format, v...)
		}
	}
}

func (m *multiLogger) Warn(v ...interface{}) {
	for _, l := range m.loggers {
		if l != nil {
			l.Warn(v...)
		}
	}
}

func (m *multiLogger) Warnf(format string, v ...interface{}) {
	for _, l := range m.loggers {
		if l != nil {
			l.Warnf(format, v...)
		}
	}
}

func (m *multiLogger) Fatal(v ...interface{}) {
	for _, l := range m.loggers {
		if l != nil {
			l.Fatal(v...)
		}
	}
}

func (m *multiLogger) Fatalf(format string, v ...interface{}) {
	for _, l := range m.loggers {
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
