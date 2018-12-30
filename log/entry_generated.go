// +build ignore

// Generated by tmpl
// https://github.com/benbjohnson/tmpl
//
// DO NOT EDIT!
// Source: entry_generated.go.tmpl

package log

import (
	"fmt"
)

func (entry *Entry) Trace(args ...interface{}) {
	if entry.EntryLevel >= TraceLevel {
		entry.log(TraceLevel, fmt.Sprint(args...))
	}
}

func (entry *Entry) Tracef(format string, args ...interface{}) {
	if entry.EntryLevel >= TraceLevel {
		entry.log(TraceLevel, fmt.Sprintf(format, args...))
	}
}

func (entry *Entry) Debug(args ...interface{}) {
	if entry.EntryLevel >= DebugLevel {
		entry.log(DebugLevel, fmt.Sprint(args...))
	}
}

func (entry *Entry) Debugf(format string, args ...interface{}) {
	if entry.EntryLevel >= DebugLevel {
		entry.log(DebugLevel, fmt.Sprintf(format, args...))
	}
}

func (entry *Entry) Info(args ...interface{}) {
	if entry.EntryLevel >= InfoLevel {
		entry.log(InfoLevel, fmt.Sprint(args...))
	}
}

func (entry *Entry) Infof(format string, args ...interface{}) {
	if entry.EntryLevel >= InfoLevel {
		entry.log(InfoLevel, fmt.Sprintf(format, args...))
	}
}

func (entry *Entry) Warn(args ...interface{}) {
	if entry.EntryLevel >= WarnLevel {
		entry.log(WarnLevel, fmt.Sprint(args...))
	}
}

func (entry *Entry) Warnf(format string, args ...interface{}) {
	if entry.EntryLevel >= WarnLevel {
		entry.log(WarnLevel, fmt.Sprintf(format, args...))
	}
}

func (entry *Entry) Error(args ...interface{}) {
	if entry.EntryLevel >= ErrorLevel {
		entry.log(ErrorLevel, fmt.Sprint(args...))
	}
}

func (entry *Entry) Errorf(format string, args ...interface{}) {
	if entry.EntryLevel >= ErrorLevel {
		entry.log(ErrorLevel, fmt.Sprintf(format, args...))
	}
}
