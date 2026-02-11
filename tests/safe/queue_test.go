package tests

import (
	"reflect"
	"testing"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
	"github.com/Polshkrev/gopolutils/safe"
)

func TestSafeQueueConstructNotNil(test *testing.T) {
	var nilQueue *safe.Queue[int] = safe.NewQueue[int]()
	if nilQueue == nil {
		test.Errorf("Queue constructor returned nil.\n")
	}
}

func TestSafeQueueAppendSuccess(test *testing.T) {
	var mock *safe.Queue[int] = safe.NewQueue[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	if !collections.In(mock, 1) {
		test.Errorf("Can not find '%d' in queue '%+v'\n", 1, *mock)
	}
}

func TestSafeQueueAppendFail(test *testing.T) {
	var mock *safe.Queue[int] = safe.NewQueue[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	if collections.In(mock, 10) {
		test.Errorf("Can not find '%d' in queue '%+v'\n", 10, *mock)
	}
}

func TestSafeQueueAtSuccess(test *testing.T) {
	var mock *safe.Queue[int] = safe.NewQueue[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var item *int
	var except *gopolutils.Exception
	item, except = mock.At(1)
	if *item != 1 || except != nil {
		test.Errorf("Can not find '%d' in queue '%+v'. %s\n", 1, *mock, except.Error())
	}
}

func TestSafeQueueAtFail(test *testing.T) {
	var mock *safe.Queue[int] = safe.NewQueue[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var item *int
	var except *gopolutils.Exception
	item, except = mock.At(20)
	if except == nil {
		test.Errorf("Except at index '%d' is nil in queue '%+v' with value '%d'.\n", 20, *mock, *item)
	}
}

func TestSafeQueueUpdateSuccess(test *testing.T) {
	var mock *safe.Queue[int] = safe.NewQueue[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var except *gopolutils.Exception = mock.Update(0, 3)
	var item *int
	var exceptAt *gopolutils.Exception
	item, exceptAt = mock.At(0)
	if except != nil || exceptAt != nil || *item != 3 {
		test.Errorf("Can not find '%d' in queue '%+v'. %s\n", 1, *mock, except.Error())
	}
}

func TestSafeQueueUpdateFail(test *testing.T) {
	var mock *safe.Queue[int] = safe.NewQueue[int]()
	var except *gopolutils.Exception = mock.Update(0, 3)
	var exceptAt *gopolutils.Exception
	_, exceptAt = mock.At(0)
	if except == nil || exceptAt == nil {
		test.Errorf("Can not find '%d' in queue '%+v'. %s\n", 1, *mock, except.Error())
	}
}

func TestSafeQueueDequeueSuccess(test *testing.T) {
	var mock *safe.Queue[int] = safe.NewQueue[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var item *int
	var except *gopolutils.Exception
	item, except = mock.Dequeue()
	if *item != 0 || except != nil {
		test.Errorf("Can not dequeue queue '%+v'. %s\n", *mock, except.Error())
	}
}

func TestSafeQueueDequeueFail(test *testing.T) {
	var mock *safe.Queue[int] = safe.NewQueue[int]()
	var except *gopolutils.Exception
	_, except = mock.Dequeue()
	if except == nil {
		test.Errorf("Item is evaluated in dequeued queue '%+v'.\n", *mock)
	}
}

func TestSafeQueueDequeueRemoves(test *testing.T) {
	var mock *safe.Queue[int] = safe.NewQueue[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var item *int
	var except *gopolutils.Exception
	item, except = mock.Dequeue()
	if *item != 0 || except != nil || collections.In(mock, 0) {
		test.Errorf("Did not dequeue queue '%+v' correctly. %s\n", *mock, except.Error())
	}
}

func TestSafeQueuePeekSuccess(test *testing.T) {
	var mock *safe.Queue[int] = safe.NewQueue[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var item *int
	var except *gopolutils.Exception
	item, except = mock.Peek()
	if *item != 0 || except != nil {
		test.Errorf("Did not peek queue '%+v' correctly. %s\n", *mock, except.Error())
	}
}

func TestSafeQueuePeekFail(test *testing.T) {
	var mock *safe.Queue[int] = safe.NewQueue[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var item *int
	item, _ = mock.Peek()
	if *item != 0 {
		test.Errorf("Item is evaluated in peeked queue '%+v' with value '%d'.\n", *mock, *item)
	}
}

func TestSafeQueuePeekDoesNotRemove(test *testing.T) {
	var mock *safe.Queue[int] = safe.NewQueue[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var item *int
	var except *gopolutils.Exception
	item, except = mock.Peek()
	if *item != 0 || except != nil || !collections.In(mock, 0) {
		test.Errorf("Did not peek queue '%+v' correctly. %s\n", *mock, except.Error())
	}
}

func TestSafeQueueCollectSuccess(test *testing.T) {
	var mock *safe.Queue[int] = safe.NewQueue[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var expect []int = []int{0, 1, 2}
	var result []int = mock.Collect()
	if !reflect.DeepEqual(result, expect) {
		test.Errorf("Queue collect was not retuned correctly: '%+v'.\n", *mock)
	}
}

func TestSafeQueueCollectFail(test *testing.T) {
	var mock *safe.Queue[int] = safe.NewQueue[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var expect []int = []int{1, 2, 3}
	var result []int = mock.Collect()
	if reflect.DeepEqual(result, expect) {
		test.Errorf("Queue collect was not retuned correctly: '%+v'.\n", *mock)
	}
}

func TestSafeQueueSizeSuccess(test *testing.T) {
	var mock *safe.Queue[int] = safe.NewQueue[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var size gopolutils.Size = mock.Size()
	if size != 3 {
		test.Errorf("Queue size was not retuned correctly: '%d'.\n", size)
	}
}

func TestSafeQueueSizeFail(test *testing.T) {
	var mock *safe.Queue[int] = safe.NewQueue[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	mock.Append(4)
	var size gopolutils.Size = mock.Size()
	if size == 3 {
		test.Errorf("Queue size was not retuned correctly: '%d'.\n", size)
	}
}

func TestSafeQueueIsEmptySuccess(test *testing.T) {
	var mock *safe.Queue[int] = safe.NewQueue[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var result bool = mock.IsEmpty()
	if result {
		test.Errorf("Queue is empty was not retuned correctly: '%t'.\n", result)
	}
}

func TestSafeQueueIsEmptyFail(test *testing.T) {
	var mock *safe.Queue[int] = safe.NewQueue[int]()
	var result bool = mock.IsEmpty()
	if !result {
		test.Errorf("Queue is empty was not retuned correctly: '%t'.\n", result)
	}
}

func TestSafeQueueItemsIsNotNil(test *testing.T) {
	var mock *safe.Queue[int] = safe.NewQueue[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	if mock.Items() == nil {
		test.Errorf("Queue items are nil.\n")
	}
}
