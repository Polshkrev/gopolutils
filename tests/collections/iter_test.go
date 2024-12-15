package tests

import (
	"testing"

	"github.com/Polshkrev/gopolutils/collections"
)

var (
	arrayMock *collections.Array[int] = collections.NewArray[int]()
)

func setupSuite(testing.TB) func(testing.TB) {
	arrayMock.Append(0)
	arrayMock.Append(1)
	arrayMock.Append(2)
	return func(testing.TB) {}
}

func TestInSucces(test *testing.T) {
	var item int = 0
	var result bool = collections.In(arrayMock, item)
	if !result {
		test.Errorf("Item '%d' is evaluated to not be in array '%+v'.", item, arrayMock)
	}
}

func TestInFail(test *testing.T) {
	var item int = 5
	var result bool = collections.In(arrayMock, item)
	if result {
		test.Errorf("Item '%d' is evaluated to be in array '%+v'.", item, arrayMock)
	}
}
