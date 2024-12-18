package tests

import (
	"testing"
	"reflect"
	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
)

func TestQueueConstructNotNil(test *testing.T) {
	var nilQueue *collections.Queue[int] = collections.NewQueue[int]()
	if nilQueue == nil {
		test.Errorf("Queue constructor returned nil.\n")
	}
}

func TestQueueAppendSuccess(test *testing.T) {
	var mock *collections.Queue[int] = collections.NewQueue[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	if !collections.In(mock, 1) {
		test.Errorf("Can not find '%d' in queue '%+v'\n", 1, *mock)
	}
}

func TestQueueAppendFail(test *testing.T) {
	var mock *collections.Queue[int] = collections.NewQueue[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	if collections.In(mock, 10) {
		test.Errorf("Can not find '%d' in queue '%+v'\n", 10, *mock)
	}
}

func TestQueueAtSuccess(test *testing.T) {
	var mock *collections.Queue[int] = collections.NewQueue[int]()
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

func TestQueueAtFail(test *testing.T) {
	var mock *collections.Queue[int] = collections.NewQueue[int]()
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

func TestQueueDequeueSuccess(test *testing.T) {
	var mock *collections.Queue[int] = collections.NewQueue[int]()
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

func TestQueueDequeueFail(test *testing.T) {
	var mock *collections.Queue[int] = collections.NewQueue[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var item *int
	var except *gopolutils.Exception
	item, except = mock.Dequeue()
	if except == nil || *item != 0 {
		test.Errorf("Except is nil in dequeued queue '%+v' with value '%d'.\n", *mock, *item)
	}
}

func TestQueueDequeueRemoves(test *testing.T) {
	var mock *collections.Queue[int] = collections.NewQueue[int]()
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

func TestQueuePeekSuccess(test *testing.T) {
	var mock *collections.Queue[int] = collections.NewQueue[int]()
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

func TestQueuePeekFail(test *testing.T) {
	var mock *collections.Queue[int] = collections.NewQueue[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var item *int
	var except *gopolutils.Exception
	item, except = mock.Peek()
	if except == nil || *item != 0 {
		test.Errorf("Except is nil in peeked queue '%+v' with value '%d'.\n", *mock, *item)
	}
}

func TestQueuePeekDoesNotRemove(test *testing.T) {
	var mock *collections.Queue[int] = collections.NewQueue[int]()
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

func TestQueueCollectSuccess(test *testing.T) {
	var mock *collections.Queue[int] = collections.NewQueue[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var expect []int = []int{0, 1, 2}
	var result []int = mock.Collect()
	if !reflect.DeepEqual(result, expect) {
		test.Errorf("Queue collect was not retuned correctly: '%+v'.\n", *mock)
	}
}

func TestQueueCollectFail(test *testing.T) {
	var mock *collections.Queue[int] = collections.NewQueue[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var expect []int = []int{1, 2, 3}
	var result []int = mock.Collect()
	if reflect.DeepEqual(result, expect) {
		test.Errorf("Queue collect was not retuned correctly: '%+v'.\n", *mock)
	}
}

func TestQueueSizeSuccess(test *testing.T) {
	var mock *collections.Queue[int] = collections.NewQueue[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var size uint64 = mock.Size()
	if size != 3 {
		test.Errorf("Queue size was not retuned correctly: '%d'.\n", size)
	}
}

func TestQueueSizeFail(test *testing.T) {
	var mock *collections.Queue[int] = collections.NewQueue[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	mock.Append(4)
	var size uint64 = mock.Size()
	if size == 3 {
		test.Errorf("Queue size was not retuned correctly: '%d'.\n", size)
	}
}

func TestQueueIsEmptySuccess(test *testing.T) {
	var mock *collections.Queue[int] = collections.NewQueue[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var result bool = mock.IsEmpty()
	if result {
		test.Errorf("Queue is empty was not retuned correctly: '%t'.\n", result)
	}
}

func TestQueueIsEmptyFail(test *testing.T) {
	var mock *collections.Queue[int] = collections.NewQueue[int]()
	var result bool = mock.IsEmpty()
	if !result {
		test.Errorf("Queue is empty was not retuned correctly: '%t'.\n", result)
	}
}

func TestQueueItemsIsNotNil(test *testing.T) {
	var mock *collections.Queue[int] = collections.NewQueue[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	if mock.Items() == nil {
		test.Errorf("Queue items are nil.\n")
	}
}