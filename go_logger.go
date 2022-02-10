package go_logger

import (
	"io"
	"log"
	"os"
)

const (
	LevelDebug = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelCritical
)

type Logger interface {
	Debug(v interface{})
	Info(v interface{})
	Warning(v interface{})
	Error(v interface{})
	Critical(v interface{})
}

type GoLogger struct {
	logger *log.Logger
	level  string
}

func convertLevelToInt(s string) int {
	levelMap := map[string]int{
		"debug":    0,
		"info":     1,
		"warning":  2,
		"error":    3,
		"critical": 4,
	}
	levelInt := levelMap[s]
	return levelInt
}


func (l *GoLogger) Debug(v interface{}) {
	levelInt := convertLevelToInt(l.level)
	if levelInt <= LevelDebug {
		l.logger.Println("[DEBUG] ", v)
	}
}

func (l *GoLogger) Info(v interface{}) {
	levelInt := convertLevelToInt(l.level)
	if levelInt <= LevelInfo {
		l.logger.Println("[INFO] ", v)
	}
}

func (l *GoLogger) Warning(v interface{}) {
	levelInt := convertLevelToInt(l.level)
	if levelInt <= LevelWarning {
		l.logger.Println("[WARNING] ", v)
	}
}

func (l *GoLogger) Error(v interface{}) {
	levelInt := convertLevelToInt(l.level)
	if levelInt <= LevelError {
		l.logger.Println("[ERROR] ", v)
	}
}

func (l *GoLogger) Critical(v interface{}) {
	levelInt := convertLevelToInt(l.level)
	if levelInt <= LevelCritical {
		l.logger.Println("[CRITICAL] ", v)
	}
}

var newLogger *GoLogger

func CreateLogger(f *os.File, level string) *GoLogger {
	// singleton
	if newLogger == nil {
		writers := []io.Writer{f, os.Stdout}
		logger := log.New(io.MultiWriter(writers...), "", log.Ldate|log.Ltime)
		newLogger = &GoLogger{logger: logger, level: level}
	}
	return newLogger
}
