package gopolutils

import "os"

type LoggingLevel = uint8

const (
	DEBUG LoggingLevel = iota
	INFO
	WARNING
	ERROR
	CRITICAL
)

func lltostr(level LoggingLevel) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case CRITICAL:
		return "CRITICAL"
	}
	return "Unknown logging level." // unreachable
}

type Logger struct {
	name    string
	level   LoggingLevel
	outputs []*os.File
}

func NewLogger(name string, level LoggingLevel) *Logger {
	var logger *Logger = new(Logger)
	logger.name = name
	logger.level = level
	logger.outputs = make([]*os.File, 0)
	return logger
}
