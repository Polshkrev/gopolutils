package gopolutils

import (
	"os"
)

const (
	AVAILABLE_OUTPUTS uint8 = 2
)

var (
	OUTPUT_COUNT uint8 = 0
)

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

func (logger *Logger) append(output *os.File) {
	logger.outputs = append(logger.outputs, output)
}

func (logger *Logger) AddConsole() *Exception {
	if OUTPUT_COUNT >= AVAILABLE_OUTPUTS {
		return NewNamedException("ValueError", "The number of outputs has exceeded the maximum allowed.")
	}
	logger.append(os.Stdout)
	OUTPUT_COUNT++
	return nil
}
