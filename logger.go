package logger

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Logger struct {
	namespace string
	now       func() string
	started   time.Time
	writer    io.Writer
}

func New(ns string) *Logger {
	return NewWriter(ns, os.Stdout)
}

func NewWriter(ns string, writer io.Writer) *Logger {
	return &Logger{namespace: ns, writer: writer}
}

func (l *Logger) At(at string) *Logger {
	return l.Namespace("at=%s", at)
}

func (l *Logger) Error(err error) {
	l.Log("state=error error=%q", err)
}

func (l *Logger) Log(format string, args ...interface{}) {
	now := time.Now().UTC().Format("2006-01-02T15:04:05.000000")

	if l.started.IsZero() {
		l.writer.Write([]byte(fmt.Sprintf("date=%q %s %s\n", now, l.namespace, fmt.Sprintf(format, args...))))
	} else {
		elapsed := float64(time.Now().Sub(l.started).Nanoseconds()) / 1000
		l.writer.Write([]byte(fmt.Sprintf("date=%q %s %s elapsed=%0.3f\n", now, l.namespace, fmt.Sprintf(format, args...), elapsed)))
	}
}

func (l *Logger) Namespace(format string, args ...interface{}) *Logger {
	return &Logger{
		namespace: fmt.Sprintf("%s %s", l.namespace, fmt.Sprintf(format, args...)),
		started:   l.started,
		writer:    l.writer,
	}
}

func (l *Logger) Start() *Logger {
	return &Logger{
		namespace: l.namespace,
		started:   time.Now(),
		writer:    l.writer,
	}
}

func (l *Logger) Success(format string, args ...interface{}) {
	l.Log("state=success %s", fmt.Sprintf(format, args...))
}
