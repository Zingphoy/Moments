package log

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var (
	DEFAULTPREFIX  = ""
	LOGCALLERDEPTH = 2
	logger         *log.Logger
	logPrefix      string
	levelTag       = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	DEBUGFLAG      = false
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

// TODO: InitLogger 配置化，如log路径、log
func InitLogger(debug bool) {
	DEBUGFLAG = debug
	logFile, err := os.OpenFile("/Users/bytedance//Developer/Moments/build/moments.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	logger = log.New(logFile, DEFAULTPREFIX, log.Ldate|log.Ltime)
}

func setLogPrefix(level Level) {
	_, file, line, ok := runtime.Caller(LOGCALLERDEPTH)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelTag[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelTag[level])
	}
	if DEBUGFLAG {
		logPrefix = "<UT>" + logPrefix
	}
	logger.SetPrefix(logPrefix)
}

func Debug(v ...interface{}) {
	setLogPrefix(DEBUG)
	logger.Println(v...)
}

func Info(v ...interface{}) {
	setLogPrefix(INFO)
	logger.Println(v...)
}

func Warn(v ...interface{}) {
	setLogPrefix(WARN)
	logger.Println(v...)
}

func Error(v ...interface{}) {
	setLogPrefix(ERROR)
	logger.Println(v...)
}

func Fatal(v ...interface{}) {
	setLogPrefix(FATAL)
	logger.Println(v...)
}
