package log

import (
	"Moments/middleware"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
)

var (
	DEFAULTPREFIX  = ""
	LOGCALLERDEPTH = 2
	logger         *log.Logger
	logPrefix      string
	levelTag       = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	DEBUGFLAG      = false
	LOG2FILE       *os.File
	LOG2STDERR     = os.Stderr
)

type Level int

const (
	TRACE Level = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

// TODO: InitLogger 配置化，如log路径、log
func InitLogger(debug bool) {
	DEBUGFLAG = debug
	LOG2FILE = setLogFile()
	logger = log.New(LOG2FILE, DEFAULTPREFIX, log.Ldate|log.Ltime)
}

func setLogFile() *os.File {
	logFile, err := os.OpenFile("/Users/bytedance/Developer/Moments//build/moments.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	return logFile
}

// RedirectLogStd redirect log stream writes to specific log file
func RedirectLogFile() {
	logger.SetOutput(LOG2FILE)
}

// RedirectLogStd redirect log stream writes to os.Stderr for local test
func RedirectLogStd() {
	logger.SetOutput(LOG2STDERR)
}

func setLogPrefix(level Level, requestId string) {
	_, file, line, ok := runtime.Caller(LOGCALLERDEPTH)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d %s]", levelTag[level], filepath.Base(file), line, requestId)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelTag[level])
	}
	if DEBUGFLAG {
		logPrefix = "<UT>" + logPrefix
	}
	logger.SetPrefix(logPrefix)
}

func Trace(c *gin.Context, v ...interface{}) {
	requestId := ""
	if c != nil {
		requestId = c.GetHeader(middleware.TrackingHeader)
	}
	setLogPrefix(TRACE, requestId)
	logger.Println(v...)
}

func Debug(c *gin.Context, v ...interface{}) {
	requestId := ""
	if c != nil {
		requestId = c.GetHeader(middleware.TrackingHeader)
	}
	setLogPrefix(DEBUG, requestId)
	logger.Println(v...)
}

func Info(c *gin.Context, v ...interface{}) {
	requestId := ""
	if c != nil {
		requestId = c.GetHeader(middleware.TrackingHeader)
	}
	setLogPrefix(INFO, requestId)
	logger.Println(v...)
}

func Warn(c *gin.Context, v ...interface{}) {
	requestId := ""
	if c != nil {
		requestId = c.GetHeader(middleware.TrackingHeader)
	}
	setLogPrefix(WARN, requestId)
	logger.Println(v...)
}

func Error(c *gin.Context, v ...interface{}) {
	requestId := ""
	if c != nil {
		requestId = c.GetHeader(middleware.TrackingHeader)
	}
	setLogPrefix(ERROR, requestId)
	logger.Println(v...)
}

func Fatal(c *gin.Context, v ...interface{}) {
	requestId := ""
	if c != nil {
		requestId = c.GetHeader(middleware.TrackingHeader)
	}
	setLogPrefix(FATAL, requestId)
	logger.Println(v...)
}
