package gopolutils

import "fmt"

type Exception struct {
	name    string
	repr    string
	message string
}

func (exception *Exception) assignRepr() {
	exception.repr = fmt.Sprintf("%s: %s", exception.name, exception.message)
}

func (exception *Exception) assignName(name string) {
	exception.name = name
	exception.assignRepr()
}

func NewException(message string) *Exception {
	var exception *Exception = new(Exception)
	exception.message = message
	exception.assignRepr()
	return exception
}

func NewNamedException(name, message string) *Exception {
	var exception *Exception = new(Exception)
	exception.assignName(name)
	exception.message = message
	exception.assignRepr()
	return exception
}
