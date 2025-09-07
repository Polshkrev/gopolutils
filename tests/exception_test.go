package tests

import (
	"testing"

	"github.com/Polshkrev/gopolutils"
)

func TestExceptionIsSuccess(test *testing.T) {
	var expectedName gopolutils.ExceptionName = gopolutils.ValueError
	var exception *gopolutils.Exception = gopolutils.NewNamedException(expectedName, "This is a mock exception")
	var result bool = exception.Is(expectedName)
	if !result {
		test.Errorf("Expected: %t got: %t", true, result)
	}
}

func TestExceptionIsFail(test *testing.T) {
	var expectedName gopolutils.ExceptionName = gopolutils.ValueError
	var testName gopolutils.ExceptionName = gopolutils.IOError
	var exception *gopolutils.Exception = gopolutils.NewNamedException(expectedName, "This is a mock exception")
	var result bool = exception.Is(testName)
	if result {
		test.Errorf("Expected: %t got: %t", false, result)
	}
}
