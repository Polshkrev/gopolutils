package gopolutils

import "fmt"

// Representation of a standardized exception.
type Exception struct {
	name    string
	repr    string
	message string
}

// Protected method to assign the final print-out of the exception.
func (exception *Exception) assignRepr() {
	exception.repr = fmt.Sprintf("%s: %s", exception.name, exception.message)
}

// Protected method to assign the name of the exception.
func (exception *Exception) assignName(name string) {
	exception.name = name
	exception.assignRepr()
}

// Obtain the raw message of the exception without the name.
// Returns the message of the exception.
func (exception Exception) Message() string {
	return exception.message
}

// Obtain the name of the exception.
// Returns the name of the exception.
func (exception Exception) Name() string {
	return exception.name
}

// Construct a new exception with a default name and a given message.
// Returns a pointer to a new exception.
func NewException(message string) *Exception {
	var exception *Exception = new(Exception)
	exception.assignName("Exception")
	exception.message = message
	exception.assignRepr()
	return exception
}

// Construct a new exception with a given name and message.
// Returns a pointer to a new exception.
func NewNamedException(name, message string) *Exception {
	var exception *Exception = new(Exception)
	exception.assignName(name)
	exception.message = message
	exception.assignRepr()
	return exception
}

// Method to adhere to the built-in error type.
// Returns a string representation of the exception.
func (exception Exception) Error() string {
	return exception.repr
}
