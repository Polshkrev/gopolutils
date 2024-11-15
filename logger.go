package gopolutils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	AVAILABLE_OUTPUTS uint8 = 2
)

const (
	TIMESTAMP_FORMAT string = "2006-01-02 15:04:05" // I've tried other's; they don't work.
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

func (logger *Logger) AddFile(fileName string) *Exception {
	if OUTPUT_COUNT >= AVAILABLE_OUTPUTS {
		return NewNamedException("ValueError", "The number of outputs has exceeded the maximum allowed.")
	}
	var file *os.File
	var except error
	file, except = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if except != nil {
		return NewException(except.Error())
	}
	logger.append(file)
	OUTPUT_COUNT++
	return nil
}

func (logger *Logger) ConsoleOnly() *Exception {
	var except *Exception = logger.AddConsole()
	if except != nil {
		return except
	}
	OUTPUT_COUNT = 2
	return nil
}

func (logger *Logger) FileOnly(fileName string) *Exception {
	var except *Exception = logger.AddFile(fileName)
	if except != nil {
		return except
	}
	OUTPUT_COUNT = 2
	return nil
}

func (logger *Logger) FullSetup(fileName string) *Exception {
	var except *Exception
	except = logger.AddConsole()
	if except != nil {
		return except
	}
	except = logger.AddFile(fileName)
	if except != nil {
		return except
	}
	return nil
}

func (logger *Logger) Write(message string, level LoggingLevel) {
	if level < logger.level {
		return
	}
	var output *os.File
	var timestamp = getTimestamp()
	var logMessage string = buildMessage(timestamp, logger.name, message, level)
	for _, output = range logger.outputs {
		var writer *bufio.Writer = bufio.NewWriter(output)
		var err error
		_, err = writer.WriteString(logMessage)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return
		}
		writer.Flush()
	}
}

func buildMessage(timestamp, name, message string, level LoggingLevel) string {
	var buffer strings.Builder = strings.Builder{}
	buffer.WriteString(timestamp)
	buffer.WriteString(":")
	buffer.WriteString(name)
	buffer.WriteString("[")
	buffer.WriteString(lltostr(level))
	buffer.WriteString("] - ")
	buffer.WriteString(message)
	buffer.WriteString("\n")
	return buffer.String()
}

func getTimestamp() string {
	var now time.Time = time.Now()
	return now.Format(TIMESTAMP_FORMAT)
}

func isFile(stream *os.File) bool {
	var found bool = false
	var streams [3]*os.File = [3]*os.File{os.Stdout, os.Stdin, os.Stderr}
	var output int
	for output = 0; output < 3; output++ {
		if streams[output] != stream {
			continue
		}
		found = true
	}
	return !found
}
