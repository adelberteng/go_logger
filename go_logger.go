package go_logger

import (
	"errors"
	"io"
	"log"
	"os"
	"sync"
)

const (
	LevelDebug = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelCritical
)

// 5 kinds of log message levels for different usage.
// if the message level lower than the logger level,
// the logger will not log the message.
type Logger interface {
	Debug(v interface{})
	Info(v interface{})
	Warning(v interface{})
	Error(v interface{})
	Critical(v interface{})
}

type GoLogger struct {
	Logger   *log.Logger
	LevelInt int
}

func convertLevelToInt(s string) int {
	levelMap := map[string]int{
		"debug":    0,
		"info":     1,
		"warning":  2,
		"error":    3,
		"critical": 4,
	}
	// if log level string wrong
	if levelInt, ok := levelMap[s]; !ok {
		return -1
	} else {
		return levelInt
	}
}

func (l *GoLogger) Debug(v interface{}) {
	if l.LevelInt <= LevelDebug {
		l.Logger.Println("[DEBUG] ", v)
	}
}

func (l *GoLogger) Info(v interface{}) {
	if l.LevelInt <= LevelInfo {
		l.Logger.Println("[INFO] ", v)
	}
}

func (l *GoLogger) Warning(v interface{}) {
	if l.LevelInt <= LevelWarning {
		l.Logger.Println("[WARNING] ", v)
	}
}

func (l *GoLogger) Error(v interface{}) {
	if l.LevelInt <= LevelError {
		l.Logger.Println("[ERROR] ", v)
	}
}

func (l *GoLogger) Critical(v interface{}) {
	if l.LevelInt <= LevelCritical {
		l.Logger.Println("[CRITICAL] ", v)
	}
}

var (
	newLogger *GoLogger
	once      sync.Once
)

func CreateLogger(f *os.File, level string) (*GoLogger, error) {
	// sync.Once to implement the singleton pattern.
	once.Do(
		func() {
			writers := []io.Writer{f, os.Stdout}
			logger := log.New(io.MultiWriter(writers...), "", log.Ldate|log.Ltime)
			levelInt := convertLevelToInt(level)
			newLogger = &GoLogger{
				Logger:   logger,
				LevelInt: levelInt,
			}
		},
	)
	// wrong log level string
	if newLogger.LevelInt == -1 {
		return &GoLogger{}, errors.New("invalid log level string")
	}

	return newLogger, nil
}
