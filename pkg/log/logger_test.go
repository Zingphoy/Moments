package log

import "testing"

var PREFIX = "test "

func init() {
	InitLogger()
}

func TestDebug(t *testing.T) {
	Debug(PREFIX + "debug")
}

func TestInfo(t *testing.T) {
	Info(PREFIX + "info")
}

func TestWarn(t *testing.T) {
	Info(PREFIX + "warn")
}

func TestError(t *testing.T) {
	Info(PREFIX + "error")
}

func TestFatal(t *testing.T) {
	Info(PREFIX + "fatal")
}
