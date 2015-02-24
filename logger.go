package logger

import (
	"fmt"
	"time"
)

type Logger struct {
	namespace string
}

func New(ns string) *Logger {
	return &Logger{namespace: ns}
}

func (l *Logger) At(at string) *Logger {
	return l.Namespace("at=%s", at)
}

func (l *Logger) Log(format string, args ...interface{}) {
	fmt.Printf("date=%s %s %s\n", time.Now().Format("2006-01-02 15:04:05.000"), l.namespace, fmt.Sprintf(format, args...))
}

func (l *Logger) Namespace(format string, args ...interface{}) *Logger {
	return &Logger{namespace: fmt.Sprintf("%s %s", l.namespace, fmt.Sprintf(format, args...))}
}

func (l *Logger) Error(err error) {
	l.Log("state=error error=%q", err)
}

func (l *Logger) Success(format string, args ...interface{}) {
	l.Log("state=success %s", fmt.Sprintf(format, args...))
}
