package gopolutils

import (
	"fmt"
	"os"
	"time"
)

const (
	// Output capacity.
	availableOutputs uint8 = 2
)

const (
	// Format to structure the log timstamp. Copied from the time module docs.
	timestampFormat string = "2006-01-02 15:04:05"
)

var (
	// Sef-explanatory.
	outputCount uint8 = 0
)

// An enum representation of the severity of a log message.
type LoggingLevel = uint8

const (
	// Lowest severity log message. Used to log debug information for development.
	DEBUG LoggingLevel = iota
	// Used to log info that should be read. It is not an error, but is not a debug message.
	INFO
	// Used to log a non-crashing warning, such as a file already existing when calling a create function.
	WARNING
	// Used to log a non-crashing error. This level should be the default when logging an error.
	ERROR
	// Used to log a crashing error. Used to log a message of a panic or breaking state.
	CRITICAL
)

// Represent a logging level as a string.
// Returns a string representation of a logging level.
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

// A logger.
type Logger struct {
	name    string
	level   LoggingLevel
	outputs [availableOutputs]*os.File
}

// Construct a new logger with a given name and default logging level.
// The default logging level passed into this constructor is the minimum level of severity that will be ouput by the logger.
// Returns a pointer to a new logger.
func NewLogger(name string, level LoggingLevel) *Logger {
	var logger *Logger = new(Logger)
	logger.name = name
	logger.level = level
	return logger
}

// Private method to append an output to the logger.
func (logger *Logger) append(output *os.File) {
	logger.outputs[outputCount] = output
}

// Set the minimal logging level for the logger.
// If the given logging level is less than or equal to the logging level already set in the logger, the method returns without modifying the logger.
func (logger *Logger) SetLevel(level LoggingLevel) {
	if level <= logger.level {
		return
	}
	logger.level = level
}

// Bind the standard output to the logger.
// If the logger has already allocated the maximum number of allowed outputs, a ValueError is returned.
func (logger *Logger) AddConsole() *Exception {
	if outputCount >= availableOutputs {
		return NewNamedException("ValueError", "The number of outputs has exceeded the maximum allowed.")
	}
	logger.append(os.Stdout)
	outputCount++
	return nil
}

// Bind a file to the logger.
// If the logger has already allocated the maximum number of allowed outputs, a ValueError is returned.
// If the given file can not be found, an Exception is returned.
func (logger *Logger) AddFile(fileName string) *Exception {
	if outputCount >= availableOutputs {
		return NewNamedException("ValueError", "The number of outputs has exceeded the maximum allowed.")
	}
	var file *os.File
	var except error
	file, except = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if except != nil {
		return NewException(except.Error())
	}
	logger.append(file)
	outputCount++
	return nil
}

// Bind only the standard output to the logger.
// If the logger has already allocated the maximum number of allowed outputs, a ValueError is returned.
func (logger *Logger) ConsoleOnly() *Exception {
	var except *Exception = logger.AddConsole()
	if except != nil {
		return except
	}
	outputCount = 2
	return nil
}

// Bind only a file to the logger.
// If the logger has already allocated the maximum number of allowed outputs, a ValueError is returned.
// If the given file can not be found, an Exception is returned.
func (logger *Logger) FileOnly(fileName string) *Exception {
	var except *Exception = logger.AddFile(fileName)
	if except != nil {
		return except
	}
	outputCount = 2
	return nil
}

// Bind both a file and the standard output to the logger.
// If the logger has already allocated the maximum number of allowed outputs, a ValueError is returned.
// If the given file can not be found, an Exception is returned.
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

// Log a message.
//
// If the default logging level of the logger is greater than the given logging level of the message, the message will not be logged.
func (logger *Logger) Log(message string, level LoggingLevel) {
	if level < logger.level {
		return
	}
	var output *os.File
	for _, output = range logger.outputs {
		publishMessage(output, getTimestamp(), logger.name, message, level)
	}
}

// Deallocate the logger.
// If the logger has a file bound, the file will need to be closed with this method.
// A good practice is to call the this method deferred even if a file is not bound;
// this method will not close the standard output.
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

// Private method to seperate the responsibility of the log method.
// This is the method that actually logs the message to the given stream.
func publishMessage(stream *os.File, timestamp, name, message string, level LoggingLevel) {
	fmt.Fprintf(stream, "%s:%s[%s] - %s\n", timestamp, name, lltostr(level), message)
}

// Private method to construct a formatted timestamp.
// Returns a string representation of a correctly fotmatted timestamp.
func getTimestamp() string {
	var now time.Time = time.Now()
	return now.Format(timestampFormat)
}

// Determine if the given stream is a file.
// Returns true if the given stream is neither the standard output, standard input, or standard error.
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
