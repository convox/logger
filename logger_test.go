package logger

import (
	"fmt"
	"strings"
	"testing"
)

var writer = &BufferWriter{}

var log = &Logger{
	namespace: "ns=test",
	writer:    writer,
	now: func() string {
		return ""
	},
}

func TestAt(t *testing.T) {
	writer.Clear()
	log.At("target").Log("foo=bar")
	assertEquals(t, writer.Line(), `date="" ns=test at=target foo=bar`)
}

func TestError(t *testing.T) {
	writer.Clear()
	log.Error(fmt.Errorf("broken"))
	assertEquals(t, writer.Line(), `date="" ns=test state=error error="broken"`)
}

func TestLog(t *testing.T) {
	writer.Clear()
	log.Log("string=%q int=%d float=%0.2f", "foo", 42, 3.14159)
	assertEquals(t, writer.Line(), `date="" ns=test string="foo" int=42 float=3.14`)
}

func TestNamespace(t *testing.T) {
	writer.Clear()
	log.Namespace("foo=bar").Namespace("baz=qux").Log("fred=barney")
	assertEquals(t, writer.Line(), `date="" ns=test foo=bar baz=qux fred=barney`)
}

func TestStart(t *testing.T) {
	writer.Clear()
	log.Start().Success("num=%d", 42)
	assertContains(t, writer.Line(), "elapsed=")
}

func TestSuccess(t *testing.T) {
	writer.Clear()
	log.Success("num=%d", 42)
	assertEquals(t, writer.Line(), `date="" ns=test state=success num=42`)
}

func assertContains(t *testing.T, got, search string) {
	if strings.Index(got, search) == -1 {
		t.Errorf("\n   expected: %q\n to contain: %q", got, search)
	}
}

func assertEquals(t *testing.T, got, expected interface{}) {
	if expected != got {
		t.Errorf("\nexpected: %q\n     got: %q", expected, got)
	}
}

type BufferWriter struct {
	buffer string
}

func (bw *BufferWriter) Buffer() string {
	return bw.buffer
}

func (bw *BufferWriter) Clear() {
	bw.buffer = ""
}

func (bw *BufferWriter) Line() string {
	return strings.Split(bw.Buffer(), "\n")[0]
}

func (bw *BufferWriter) Write(data []byte) (int, error) {
	bw.buffer += string(data)
	return len(data), nil
}
