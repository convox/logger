package logger

import "fmt"

type Logger struct {
	Namespace string
}

func New(ns string) *Logger {
	return &Logger{Namespace: ns}
}

func (l *Logger) At(at string) *Logger {
	return &Logger{Namespace: fmt.Sprintf("%s at=%s", l.Namespace, at)}
}

func (l *Logger) Log(format string, args ...interface{}) {
	fmt.Printf("%s %s\n", l.Namespace, fmt.Sprintf(format, args...))
}

func (l *Logger) Error(err error) {
	l.Log("state=error error=%q", err)
}

func (l *Logger) Success(format string, args ...interface{}) {
	l.Log("state=success %s", fmt.Sprintf(format, args...))
}
