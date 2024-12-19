package tests

import (
	"reflect"
	"testing"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
)

func TestStackConstructNotNil(test *testing.T) {
	var nilStack *collections.Stack[int] = collections.NewStack[int]()
	if nilStack == nil {
		test.Errorf("Stack constructor returned nil.\n")
	}
}

func TestStackAppendSuccess(test *testing.T) {
	var mock *collections.Stack[int] = collections.NewStack[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	if !collections.In(mock, 1) {
		test.Errorf("Can not find '%d' in stack '%+v'\n", 1, *mock)
	}
}

func TestStackAppendFail(test *testing.T) {
	var mock *collections.Stack[int] = collections.NewStack[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	if collections.In(mock, 10) {
		test.Errorf("Can not find '%d' in stack '%+v'\n", 10, *mock)
	}
}

func TestStackAtSuccess(test *testing.T) {
	var mock *collections.Stack[int] = collections.NewStack[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var item *int
	var except *gopolutils.Exception
	item, except = mock.At(1)
	if *item != 1 || except != nil {
		test.Errorf("Can not find '%d' in stack '%+v'. %s\n", 1, *mock, except.Error())
	}
}

func TestStackAtFail(test *testing.T) {
	var mock *collections.Stack[int] = collections.NewStack[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var item *int
	var except *gopolutils.Exception
	item, except = mock.At(20)
	if except == nil {
		test.Errorf("Except at index '%d' is nil in stack '%+v' with value '%d'.\n", 20, *mock, *item)
	}
}

func TestStackUpdateSuccess(test *testing.T) {
	var mock *collections.Stack[int] = collections.NewStack[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var except *gopolutils.Exception = mock.Update(0, 3)
	var item *int
	var exceptAt *gopolutils.Exception
	item, exceptAt = mock.At(0)
	if except != nil || exceptAt != nil || *item != 3 {
		test.Errorf("Can not find '%d' in stack '%+v'. %s\n", 1, *mock, except.Error())
	}
}

func TestStackUpdateFail(test *testing.T) {
	var mock *collections.Stack[int] = collections.NewStack[int]()
	var except *gopolutils.Exception = mock.Update(0, 3)
	var exceptAt *gopolutils.Exception
	_, exceptAt = mock.At(0)
	if except == nil || exceptAt == nil {
		test.Errorf("Can not find '%d' in stack '%+v'. %s\n", 1, *mock, except.Error())
	}
}

func TestStackPopSuccess(test *testing.T) {
	var mock *collections.Stack[int] = collections.NewStack[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var item *int
	var except *gopolutils.Exception
	item, except = mock.Pop()
	if *item != 2 || except != nil || collections.In(mock, 2) {
		test.Errorf("Did not pop stack '%+v' correctly. %s\n", *mock, except.Error())
	}
}

func TestStackPopFail(test *testing.T) {
	var mock *collections.Stack[int] = collections.NewStack[int]()
	var except *gopolutils.Exception
	_, except = mock.Pop()
	if except == nil || collections.In(mock, 2) {
		test.Errorf("Item is evaluated in popped stack '%+v'.\n", *mock)
	}
}

func TestStackPopRemoves(test *testing.T) {
	var mock *collections.Stack[int] = collections.NewStack[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var item *int
	var except *gopolutils.Exception
	item, except = mock.Pop()
	if *item != 2 || except != nil || collections.In(mock, 2) {
		test.Errorf("Did not pop from stack '%+v' correctly. %s\n", *mock, except.Error())
	}
}

func TestStackPeekSuccess(test *testing.T) {
	var mock *collections.Stack[int] = collections.NewStack[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var item *int
	var except *gopolutils.Exception
	item, except = mock.Peek()
	if *item != 2 || except != nil || !collections.In(mock, 2) {
		test.Errorf("Did not peek stack '%+v' correctly. %s\n", *mock, except.Error())
	}
}

func TestStackPeekFail(test *testing.T) {
	var mock *collections.Stack[int] = collections.NewStack[int]()
	var except *gopolutils.Exception
	_, except = mock.Peek()
	if except == nil || collections.In(mock, 2) {
		test.Errorf("Item is evaluated in peeked stack '%+v'.\n", *mock)
	}
}

func TestStackPeekDoesNotRemove(test *testing.T) {
	var mock *collections.Stack[int] = collections.NewStack[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var item *int
	var except *gopolutils.Exception
	item, except = mock.Peek()
	if *item != 2 || except != nil || !collections.In(mock, 2) {
		test.Errorf("Did not peek stack '%+v' correctly. %s\n", *mock, except.Error())
	}
}

func TestStackCollectSuccess(test *testing.T) {
	var mock *collections.Stack[int] = collections.NewStack[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var expect []int = []int{0, 1, 2}
	var result []int = mock.Collect()
	if !reflect.DeepEqual(result, expect) {
		test.Errorf("Stack collect was not retuned correctly: '%+v'.\n", *mock)
	}
}

func TestStackCollectFail(test *testing.T) {
	var mock *collections.Stack[int] = collections.NewStack[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var expect []int = []int{1, 2, 3}
	var result []int = mock.Collect()
	if reflect.DeepEqual(result, expect) {
		test.Errorf("Stack collect was not retuned correctly: '%+v'.\n", *mock)
	}
}

func TestStackSizeSuccess(test *testing.T) {
	var mock *collections.Stack[int] = collections.NewStack[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var size uint64 = mock.Size()
	if size != 3 {
		test.Errorf("Stack size was not retuned correctly: '%d'.\n", size)
	}
}

func TestStackSizeFail(test *testing.T) {
	var mock *collections.Stack[int] = collections.NewStack[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	mock.Append(4)
	var size uint64 = mock.Size()
	if size == 3 {
		test.Errorf("Stack size was not retuned correctly: '%d'.\n", size)
	}
}

func TestStackIsEmptySuccess(test *testing.T) {
	var mock *collections.Stack[int] = collections.NewStack[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var result bool = mock.IsEmpty()
	if result {
		test.Errorf("Stack is empty was not retuned correctly: '%t'.\n", result)
	}
}

func TestStackIsEmptyFail(test *testing.T) {
	var mock *collections.Stack[int] = collections.NewStack[int]()
	var result bool = mock.IsEmpty()
	if !result {
		test.Errorf("Stack is empty was not retuned correctly: '%t'.\n", result)
	}
}

func TestStackItemsIsNotNil(test *testing.T) {
	var mock *collections.Stack[int] = collections.NewStack[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	if mock.Items() == nil {
		test.Errorf("Stack items are nil.\n")
	}
}
