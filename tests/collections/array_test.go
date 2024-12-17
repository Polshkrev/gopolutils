package tests

import (
	"testing"
	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
)

func TestArrayConstructNotNil(test *testing.T) {
	var nilArray *collections.Array[int] = collections.NewArray[int]()
	if nilArray == nil {
		test.Errorf("Array constructor returned nil.\n")
	}
}

func TestArrayAppendSuccess(test *testing.T) {
	var mock *collections.Array[int] = collections.NewArray[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	if !collections.In(mock, 1) {
		test.Errorf("Can not find '%d' in array '%+v'\n", 1, *mock)
	}
}

func TestArrayAppendFail(test *testing.T) {
	var mock *collections.Array[int] = collections.NewArray[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	if collections.In(mock, 10) {
		test.Errorf("Can not find '%d' in array '%+v'\n", 10, *mock)
	}
}

func TestArrayAtSuccess(test *testing.T) {
	var mock *collections.Array[int] = collections.NewArray[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var item *int
	var except *gopolutils.Exception
	item, except = mock.At(1)
	if *item != 1 || except != nil{
		test.Errorf("Can not find '%d' in array '%+v'. %s\n", 1, *mock, except.Error())
	}
}

func TestArrayAtFail(test *testing.T) {
	var mock *collections.Array[int] = collections.NewArray[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var item *int
	var except *gopolutils.Exception
	item, except = mock.At(20)
	if except == nil {
		test.Errorf("Except at index '%d' is nil in array '%+v' with value '%d'.\n", 20, *mock, *item)
	}
}

func TestArrayRemoveSuccess(test *testing.T) {
	var mock *collections.Array[int] = collections.NewArray[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var _ *gopolutils.Exception = mock.Remove(1)
	if collections.In(mock, 1) {
		test.Errorf("'%d' was not removed from array '%+v'.\n", 1, *mock)
	}
}

func TestArrayRemoveFail(test *testing.T) {
	var mock *collections.Array[int] = collections.NewArray[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var except *gopolutils.Exception = mock.Remove(8)
	if except == nil || !collections.In(mock, 1) {
		test.Errorf("'%d' was not removed from array '%+v'.\n", 1, *mock)
	}
}