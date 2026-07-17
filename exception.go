package gopolutils

import "fmt"

// Representation of a standardized exception.
type Exception struct {
	name    ExceptionName
	repr    string
	message string
}

// Protected method to assign the final print-out of the exception.
func (exception *Exception) assignRepr() {
	exception.repr = fmt.Sprintf("%s: %s", exception.name, exception.message)
}

// Protected method to assign the name of the exception.
func (exception *Exception) assignName(name ExceptionName) {
	exception.name = name
	exception.assignRepr()
}

// Protected method to assign the message of an exception.
func (exception *Exception) assignMessage(message string) {
	exception.message = message
}

// Obtain the raw message of the exception without the name.
// Returns the message of the exception.
func (exception Exception) Message() string {
	return exception.message
}

// Obtain the name of the exception.
// Returns the name of the exception.
func (exception Exception) Name() ExceptionName {
	return exception.name
}

// Construct a new exception with a default name and a given message.
// Returns a pointer to a new exception.
func NewException(format string, arguments ...any) *Exception {
	var exception *Exception = new(Exception)
	exception.assignName(BaseException)
	exception.assignMessage(fmt.Sprintf(format, arguments...))
	exception.assignRepr()
	return exception
}

// Construct a new exception with a given name and format specifiers with arguments.
// Returns a pointer to a new exception.
func NewNamedException(name ExceptionName, format string, arguments ...any) *Exception {
	var exception *Exception = new(Exception)
	exception.assignName(name)
	exception.assignMessage(fmt.Sprintf(format, arguments...))
	exception.assignRepr()
	return exception
}

// Method to adhere to the built-in error type.
// Returns a string representation of the exception.
func (exception Exception) Error() string {
	return exception.repr
}

// Determine if the exception has a specific name.
// Similar to the [errors.Is] function in the standard library.
// Returns true if the exception has the given name, else false.
func (exception Exception) Is(name ExceptionName) bool {
	return exception.name == name
}

// If the given exception is not nil, the function panics, else the function returns the given result.
func Must[Type any](result Type, except *Exception) Type {
	if except != nil {
		panic(except.Error())
	}
	return result
}
