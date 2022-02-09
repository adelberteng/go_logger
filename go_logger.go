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

type Logger interface {
	Debug(o interface{})
	Info(o interface{})
	Warning(o interface{})
	Error(o interface{})
	Critical(o interface{})
}

type GoLogger struct {
	logger *log.Logger
	level  string
}

func (l *GoLogger) Debug(o interface{}) {
	levelInt := convertLevelToInt(l.level)
	if levelInt <= LevelDebug {
		l.logger.Println("[DEBUG] ", o)
	}
}

func (l *GoLogger) Info(o interface{}) {
	levelInt := convertLevelToInt(l.level)
	if levelInt <= LevelInfo {
		l.logger.Println("[INFO] ", o)
	}
}

func (l *GoLogger) Warning(o interface{}) {
	levelInt := convertLevelToInt(l.level)
	if levelInt <= LevelWarning {
		l.logger.Println("[WARNING] ", o)
	}
}

func (l *GoLogger) Error(o interface{}) {
	levelInt := convertLevelToInt(l.level)
	if levelInt <= LevelError {
		l.logger.Println("[ERROR] ", o)
	}
}

func (l *GoLogger) Critical(o interface{}) {
	levelInt := convertLevelToInt(l.level)
	if levelInt <= LevelCritical {
		l.logger.Println("[CRITICAL] ", o)
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
