package tests

import (
	"reflect"
	"testing"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
)

func TestSafeArrayConstructNotNil(test *testing.T) {
	var nilArray *collections.Array[int] = collections.NewArray[int]()
	if nilArray == nil {
		test.Errorf("Array constructor returned nil.\n")
	}
}

func TestSafeArrayAppendSuccess(test *testing.T) {
	var mock *collections.Array[int] = collections.NewArray[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	if !collections.In(mock, 1) {
		test.Errorf("Can not find '%d' in array '%+v'\n", 1, *mock)
	}
}

func TestSafeArrayAppendFail(test *testing.T) {
	var mock *collections.Array[int] = collections.NewArray[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	if collections.In(mock, 10) {
		test.Errorf("Can not find '%d' in array '%+v'\n", 10, *mock)
	}
}

func TestSafeArrayAtSuccess(test *testing.T) {
	var mock *collections.Array[int] = collections.NewArray[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var item *int
	var except *gopolutils.Exception
	item, except = mock.At(1)
	if *item != 1 || except != nil {
		test.Errorf("Can not find '%d' in array '%+v'. %s\n", 1, *mock, except.Error())
	}
}

func TestSafeArrayAtFail(test *testing.T) {
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

func TestSafeArrayUpdateSuccess(test *testing.T) {
	var mock *collections.Array[int] = collections.NewArray[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var except *gopolutils.Exception = mock.Update(0, 3)
	var item *int
	var exceptAt *gopolutils.Exception
	item, exceptAt = mock.At(0)
	if except != nil || exceptAt != nil || *item != 3 {
		test.Errorf("Can not find '%d' in array '%+v'. %s\n", 0, *mock, except.Error())
	}
}

func TestSafeArrayUpdateFail(test *testing.T) {
	var mock *collections.Array[int] = collections.NewArray[int]()
	var except *gopolutils.Exception = mock.Update(0, 3)
	var exceptAt *gopolutils.Exception
	_, exceptAt = mock.At(0)
	if except == nil || exceptAt == nil {
		test.Errorf("Can not find '%d' in array '%+v'.\n", 0, *mock)
	}
}

func TestSafeArrayRemoveSuccess(test *testing.T) {
	var mock *collections.Array[int] = collections.NewArray[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var _ *gopolutils.Exception = mock.Remove(1)
	if collections.In(mock, 1) {
		test.Errorf("'%d' was not removed from array '%+v'.\n", 1, *mock)
	}
}

func TestSafeArrayRemoveFail(test *testing.T) {
	var mock *collections.Array[int] = collections.NewArray[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var except *gopolutils.Exception = mock.Remove(8)
	if except == nil || !collections.In(mock, 1) {
		test.Errorf("'%d' was not removed from array '%+v'.\n", 1, *mock)
	}
}

func TestSafeArrayCollectSuccess(test *testing.T) {
	var mock *collections.Array[int] = collections.NewArray[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var expect []int = []int{0, 1, 2}
	var result []int = mock.Collect()
	if !reflect.DeepEqual(result, expect) {
		test.Errorf("Array collect was not retuned correctly: '%+v'.\n", *mock)
	}
}

func TestSafeArrayCollectFail(test *testing.T) {
	var mock *collections.Array[int] = collections.NewArray[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var expect []int = []int{1, 2, 3}
	var result []int = mock.Collect()
	if reflect.DeepEqual(result, expect) {
		test.Errorf("Array collect was not retuned correctly: '%+v'.\n", *mock)
	}
}

func TestSafeArraySizeSuccess(test *testing.T) {
	var mock *collections.Array[int] = collections.NewArray[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var size gopolutils.Size = mock.Size()
	if size != 3 {
		test.Errorf("Array size was not retuned correctly: '%d'.\n", size)
	}
}

func TestSafeArraySizeFail(test *testing.T) {
	var mock *collections.Array[int] = collections.NewArray[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	mock.Append(4)
	var size gopolutils.Size = mock.Size()
	if size == 3 {
		test.Errorf("Array size was not retuned correctly: '%d'.\n", size)
	}
}

func TestSafeArrayIsEmptySuccess(test *testing.T) {
	var mock *collections.Array[int] = collections.NewArray[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var result bool = mock.IsEmpty()
	if result {
		test.Errorf("Array is empty was not retuned correctly: '%t'.\n", result)
	}
}

func TestSafeArrayIsEmptyFail(test *testing.T) {
	var mock *collections.Array[int] = collections.NewArray[int]()
	var result bool = mock.IsEmpty()
	if !result {
		test.Errorf("Array is empty was not retuned correctly: '%t'.\n", result)
	}
}

func TestSafeArrayItemsIsNotNil(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()
	array.Append(0)
	array.Append(1)
	array.Append(2)
	if array.Items() == nil {
		test.Errorf("Array items is nil.\n")
	}
}
