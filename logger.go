package gopolutils

import (
	"fmt"
	"os"
	"time"
)

const (
	__AVAILABLE_OUTPUTS uint8 = 2
)

const (
	__TIMESTAMP_FORMAT string = "2006-01-02 15:04:05" // I've tried other's; they don't work.
)

var (
	__OUTPUT_COUNT uint8 = 0
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
	if __OUTPUT_COUNT >= __AVAILABLE_OUTPUTS {
		return NewNamedException("ValueError", "The number of outputs has exceeded the maximum allowed.")
	}
	logger.append(os.Stdout)
	__OUTPUT_COUNT++
	return nil
}

func (logger *Logger) AddFile(fileName string) *Exception {
	if __OUTPUT_COUNT >= __AVAILABLE_OUTPUTS {
		return NewNamedException("ValueError", "The number of outputs has exceeded the maximum allowed.")
	}
	var file *os.File
	var except error
	file, except = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if except != nil {
		return NewException(except.Error())
	}
	logger.append(file)
	__OUTPUT_COUNT++
	return nil
}

func (logger *Logger) ConsoleOnly() *Exception {
	var except *Exception = logger.AddConsole()
	if except != nil {
		return except
	}
	__OUTPUT_COUNT = 2
	return nil
}

func (logger *Logger) FileOnly(fileName string) *Exception {
	var except *Exception = logger.AddFile(fileName)
	if except != nil {
		return except
	}
	__OUTPUT_COUNT = 2
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

func (logger *Logger) Log(message string, level LoggingLevel) {
	if level < logger.level {
		return
	}
	var output *os.File
	for _, output = range logger.outputs {
		publishMessage(output, getTimestamp(), logger.name, message, level)
	}
}

func (logger *Logger) Close() {
	var output *os.File
	for _, output = range logger.outputs {
		if !isFile(output) {
			continue
		} else if output != nil {
			continue
		}
		var err error = output.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	}
}

func publishMessage(stream *os.File, timestamp, name, message string, level LoggingLevel) {
	fmt.Fprintf(stream, "%s:%s[%s] - %s\n", timestamp, name, lltostr(level), message)
}

func getTimestamp() string {
	var now time.Time = time.Now()
	return now.Format(__TIMESTAMP_FORMAT)
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
