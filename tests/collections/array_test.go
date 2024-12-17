package tests

import (
	"testing"
	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
)

var (
	arrayMock *collections.Array[int] = collections.NewArray[int]()
)

func setupArrayMock(testing.TB) func(testing.TB) {
	arrayMock.Append(0)
	arrayMock.Append(1)
	arrayMock.Append(2)
	return func(testing.TB) {}
}

func TestArrayConstructNotNil(test *testing.T) {
	var nilArray *collections.Array[int] = collections.NewArray[int]()
	if nilArray == nil {
		test.Errorf("Array constructor returned nil.\n")
	}
}

func TestArrayAppendSuccess(test *testing.T) {
	var teardown func(testing.TB) = setupArrayMock(test)
	defer teardown(test)
	if !collections.In(arrayMock, 1) {
		test.Errorf("Can not find '%d' in array '%+v'\n", 1, *arrayMock)
	}
}

func TestArrayAppendFail(test *testing.T) {
	var teardown func(testing.TB) = setupArrayMock(test)
	defer teardown(test)
	if collections.In(arrayMock, 10) {
		test.Errorf("Can not find '%d' in array '%+v'\n", 10, *arrayMock)
	}
}

func TestArrayAtSuccess(test *testing.T) {
	var teardown func(testing.TB) = setupArrayMock(test)
	defer teardown(test)
	var item *int
	var except *gopolutils.Exception
	item, except = arrayMock.At(1)
	if *item != 1 || except != nil{
		test.Errorf("Can not find '%d' in array '%+v'. %s\n", 1, *arrayMock, except.Error())
	}
}

func TestArrayAtFail(test *testing.T) {
	var teardown func(testing.TB) = setupArrayMock(test)
	defer teardown(test)
	var item *int
	var except *gopolutils.Exception
	item, except = arrayMock.At(10)
	if except == nil {
		test.Errorf("Except at index '%d' is nil in array '%+v' with value '%d'.\n", 10, *arrayMock, *item)
	}
}