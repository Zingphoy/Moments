package log

import "testing"

var PREFIX = "test "

func init() {
	InitLogger(true)
}

func TestTrace(t *testing.T) {
	Trace(nil, PREFIX+"trace")
}

func TestDebug(t *testing.T) {
	Debug(nil, PREFIX+"debug")
}

func TestInfo(t *testing.T) {
	Info(nil, PREFIX+"info")
}

func TestWarn(t *testing.T) {
	Info(nil, PREFIX+"warn")
}

func TestError(t *testing.T) {
	Info(nil, PREFIX+"error")
}

func TestFatal(t *testing.T) {
	Info(nil, PREFIX+"fatal")
}
