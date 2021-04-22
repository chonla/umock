package logger

import (
	"log"
	"os"
)

const (
	TRACE = 0
	DEBUG = 1
	INFO  = 2
	WARN  = 3
	ERROR = 4
)

var (
	traceLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	debugLogger *log.Logger
	errorLogger *log.Logger
)

func init() {
	traceLogger = log.New(os.Stderr, "TRACE: ", log.Ldate|log.Ltime)
	infoLogger = log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime)
	warnLogger = log.New(os.Stderr, "WARN: ", log.Ldate|log.Ltime)
	debugLogger = log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
}

type Logger struct {
	level int
}

func New(logLevel int) *Logger {
	return &Logger{level: logLevel}
}

func (l *Logger) Level(logLevel int) {
	l.level = logLevel
}

func (l *Logger) Info(f string, v ...interface{}) {
	if l.level >= INFO {
		infoLogger.Printf(f, v...)
	}
}

func (l *Logger) Trace(f string, v ...interface{}) {
	if l.level >= TRACE {
		traceLogger.Printf(f, v...)
	}
}

func (l *Logger) Warn(f string, v ...interface{}) {
	if l.level >= WARN {
		warnLogger.Printf(f, v...)
	}
}

func (l *Logger) Debug(f string, v ...interface{}) {
	if l.level >= DEBUG {
		debugLogger.Printf(f, v...)
	}
}

func (l *Logger) Error(f string, v ...interface{}) {
	if l.level >= ERROR {
		errorLogger.Printf(f, v...)
	}
}
