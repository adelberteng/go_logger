package go_logger

import (
	"os"
	"io"
	"log"
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
		"debug": 0,
		"info": 1,
		"warning": 2,
		"error": 3,
		"critical": 4,
	}
	levelInt := levelMap[s]
	return levelInt
}

type Logger struct {
	logger *log.Logger
	level string
}

func (l *Logger) Debug(o interface{}) {
	levelInt := convertLevelToInt(l.level)
	if levelInt <= LevelDebug {
		l.logger.Println("[DEBUG] ", o)
	}
}

func (l *Logger) Info(o interface{}) {
	levelInt := convertLevelToInt(l.level)
	if levelInt <= LevelInfo {
		l.logger.Println("[INFO] ", o)
	}
}

func (l *Logger) Warning(o interface{}) {
	levelInt := convertLevelToInt(l.level)
	if levelInt <= LevelWarning {
		l.logger.Println("[WARNING] ", o)
	}
}

func (l *Logger) Error(o interface{}) {
	levelInt := convertLevelToInt(l.level)
	if levelInt <= LevelError {
		l.logger.Println("[ERROR] ", o)
	}
}

func (l *Logger) Critical(o interface{}) {
	levelInt := convertLevelToInt(l.level)
	if levelInt <= LevelCritical {
		l.logger.Println("[CRITICAL] ", o)
	}
}

func CreateLogger(f *os.File, level string) Logger {
	writers := []io.Writer{f, os.Stdout}
	logger := log.New(io.MultiWriter(writers...), "", log.Ldate|log.Ltime)
	newLogger := Logger{logger: logger, level: level}

	return newLogger
}