package tests

import (
	"testing"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
)

func TestNewQueue(t *testing.T) {
	var queue *collections.Queue[int] = collections.NewQueue[int]()

	if queue == nil {
		t.Fatal("expected queue")
	} else if !queue.IsEmpty() {
		t.Fatal("new queue should be empty")
	} else if queue.Size() != 0 {
		t.Fatalf("expected size 0, got %d", queue.Size())
	}
}

func TestQueueAppend(t *testing.T) {
	var queue *collections.Queue[int] = collections.NewQueue[int]()

	queue.Append(1)
	queue.Append(2)
	queue.Append(3)

	if queue.Size() != 3 {
		t.Fatalf("expected size 3, got %d", queue.Size())
	}

	var expected []int = []int{1, 2, 3}
	var i, value int
	for i, value = range expected {
		var item *int
		var exception *gopolutils.Exception
		item, exception = queue.At(gopolutils.Size(i))

		if exception != nil {
			t.Fatal(exception)
		} else if *item != value {
			t.Fatalf("expected %d, got %d", value, *item)
		}
	}
}

func TestQueueExtend(t *testing.T) {
	var queue *collections.Queue[int] = collections.NewQueue[int]()

	var values *collections.Array[int] = collections.NewArray[int]()
	values.Append(1)
	values.Append(2)
	values.Append(3)

	queue.Extend(values)

	if queue.Size() != 3 {
		t.Fatalf("expected size 3, got %d", queue.Size())
	}

	var expected []int = []int{1, 2, 3}
	var i, value int
	for i, value = range expected {
		var item *int
		var exception *gopolutils.Exception
		item, exception = queue.At(gopolutils.Size(i))
		if exception != nil {
			t.Fatal(exception)
		} else if *item != value {
			t.Fatalf("expected %d, got %d", value, *item)
		}
	}
}

func TestQueueAt(t *testing.T) {
	var queue *collections.Queue[int] = collections.NewQueue[int]()

	queue.Append(42)
	var item *int
	var exception *gopolutils.Exception
	item, exception = queue.At(0)

	if exception != nil {
		t.Fatal(exception)
	} else if *item != 42 {
		t.Fatalf("expected 42, got %d", *item)
	}
}

func TestQueueAtEmpty(t *testing.T) {
	var queue *collections.Queue[int] = collections.NewQueue[int]()
	var item *int
	var exception *gopolutils.Exception
	item, exception = queue.At(0)

	if exception == nil {
		t.Fatal("expected exception")
	} else if item != nil {
		t.Fatal("expected nil item")
	}
}

func TestQueueAtOutOfRange(t *testing.T) {
	var queue *collections.Queue[int] = collections.NewQueue[int]()

	queue.Append(1)

	var item *int
	var exception *gopolutils.Exception
	item, exception = queue.At(2)

	if exception == nil {
		t.Fatal("expected exception")
	} else if item != nil {
		t.Fatal("expected nil item")
	}
}

func TestQueueUpdate(t *testing.T) {
	var queue *collections.Queue[int] = collections.NewQueue[int]()

	queue.Append(1)

	var exception *gopolutils.Exception = queue.Update(0, 100)
	if exception != nil {
		t.Fatal(exception)
	}

	var item *int
	item, exception = queue.At(0)
	if exception != nil {
		t.Fatalf("unexpected exception: %s", exception.Error())
	} else if *item != 100 {
		t.Fatalf("expected 100, got %d", *item)
	}
}

func TestQueueUpdateEmpty(t *testing.T) {
	var queue *collections.Queue[int] = collections.NewQueue[int]()

	if queue.Update(0, 1) == nil {
		t.Fatal("expected exception")
	}
}

func TestQueueUpdateOutOfRange(t *testing.T) {
	var queue *collections.Queue[int] = collections.NewQueue[int]()

	queue.Append(1)

	if queue.Update(2, 5) == nil {
		t.Fatal("expected exception")
	}

	var item *int
	var exception *gopolutils.Exception
	item, exception = queue.At(0)
	if exception != nil {
		t.Fatal(exception)
	} else if *item != 1 {
		t.Fatal("queue should not have been modified")
	}
}

func TestQueueRemove(t *testing.T) {
	var queue *collections.Queue[int] = collections.NewQueue[int]()

	queue.Append(1)
	queue.Append(2)
	queue.Append(3)

	var exception *gopolutils.Exception = queue.Remove(1)
	if exception != nil {
		t.Fatal(exception)
	}

	var expected []int = []int{1, 3}

	if queue.Size() != 2 {
		t.Fatalf("expected size 2, got %d", queue.Size())
	}
	var i, value int
	for i, value = range expected {

		var item *int
		item, exception = queue.At(gopolutils.Size(i))

		if *item != value {
			t.Fatalf("expected %d, got %d", value, *item)
		}
	}
}

func TestQueueRemoveEmpty(t *testing.T) {
	var queue *collections.Queue[int] = collections.NewQueue[int]()

	if queue.Remove(0) == nil {
		t.Fatal("expected exception")
	}
}

func TestQueueRemoveOutOfRange(t *testing.T) {
	var queue *collections.Queue[int] = collections.NewQueue[int]()

	queue.Append(1)

	if queue.Remove(2) == nil {
		t.Fatal("expected exception")
	} else if queue.Size() != 1 {
		t.Fatal("queue should not have been modified")
	}
}

func TestDequeue(t *testing.T) {
	var queue *collections.Queue[int] = collections.NewQueue[int]()

	queue.Append(1)
	queue.Append(2)
	queue.Append(3)

	var expected []int = []int{1, 2, 3}
	var i int
	for i = range expected {
		var value int = expected[i]

		var item *int
		var exception *gopolutils.Exception
		item, exception = queue.Dequeue()
		if exception != nil {
			t.Fatal(exception)
		} else if *item != value {
			t.Fatalf("expected %d, got %d", value, *item)
		}
	}

	if !queue.IsEmpty() {
		t.Fatal("expected empty queue")
	}
}

func TestDequeueEmpty(t *testing.T) {
	var queue *collections.Queue[int] = collections.NewQueue[int]()
	var item *int
	var exception *gopolutils.Exception
	item, exception = queue.Dequeue()

	if exception == nil {
		t.Fatal("expected exception")
	} else if item != nil {
		t.Fatal("expected nil item")
	}
}

func TestQueuePeek(t *testing.T) {
	var queue *collections.Queue[int] = collections.NewQueue[int]()

	queue.Append(10)
	queue.Append(20)

	var item *int
	var exception *gopolutils.Exception
	item, exception = queue.Peek()
	if exception != nil {
		t.Fatal(exception)
	} else if *item != 10 {
		t.Fatalf("expected 10, got %d", *item)
	} else if queue.Size() != 2 {
		t.Fatal("peek should not remove element")
	}
}

func TestQueuePeekEmpty(t *testing.T) {
	var queue *collections.Queue[int] = collections.NewQueue[int]()

	var item *int
	var exception *gopolutils.Exception
	item, exception = queue.Peek()
	if exception == nil {
		t.Fatal("expected exception")
	} else if item != nil {
		t.Fatal("expected nil item")
	}
}

func TestQueueItems(t *testing.T) {
	var queue *collections.Queue[int] = collections.NewQueue[int]()

	queue.Append(1)
	queue.Append(2)

	var items *[]int = queue.Items()

	(*items)[0] = 42
	var item *int
	var exception *gopolutils.Exception
	item, exception = queue.At(0)
	if exception != nil {
		t.Fatal(exception)
	} else if *item != 42 {
		t.Fatal("expected modification through Items()")
	}
}

func TestQueueCollect(t *testing.T) {
	var queue *collections.Queue[int] = collections.NewQueue[int]()

	queue.Append(1)
	queue.Append(2)

	var values []int = queue.Collect()

	if len(values) != 2 {
		t.Fatal("unexpected length")
	} else if values[0] != 1 || values[1] != 2 {
		t.Fatal("unexpected values")
	}
}

func TestQueueIterator(t *testing.T) {
	var queue *collections.Queue[int] = collections.NewQueue[int]()

	queue.Append(1)
	queue.Append(2)
	queue.Append(3)

	var count int = 0

	queue.Iterator().ForEach(func(value int) {
		count++
	})

	if count != 3 {
		t.Fatalf("expected 3 elements, got %d", count)
	}
}

func TestQueueSize(t *testing.T) {
	var queue *collections.Queue[int] = collections.NewQueue[int]()
	var i int
	for i = range 10 {
		queue.Append(i)
	}

	if queue.Size() != 10 {
		t.Fatalf("expected size 10, got %d", queue.Size())
	}
}

func TestQueueIsEmpty(t *testing.T) {
	var queue *collections.Queue[int] = collections.NewQueue[int]()

	if !queue.IsEmpty() {
		t.Fatal("expected empty queue")
	}

	queue.Append(1)

	if queue.IsEmpty() {
		t.Fatal("expected non-empty queue")
	}
}

func TestAppendAfterDequeue(t *testing.T) {
	var queue *collections.Queue[int] = collections.NewQueue[int]()

	queue.Append(1)
	queue.Append(2)

	queue.Dequeue()
	queue.Append(3)

	var expected []int = []int{2, 3}
	var i, value int
	for i, value = range expected {
		var item *int
		var exception *gopolutils.Exception
		item, exception = queue.At(gopolutils.Size(i))
		if exception != nil {
			t.Fatal(exception)
		} else if *item != value {
			t.Fatalf("expected %d, got %d", value, *item)
		}
	}
}

func TestQueueAtIndexEqualsSize(t *testing.T) {
	var queue *collections.Queue[int] = collections.NewQueue[int]()

	queue.Append(1)

	var item *int
	var exception *gopolutils.Exception
	item, exception = queue.At(queue.Size())
	if exception == nil {
		t.Fatal("expected out of range exception")
	} else if item != nil {
		t.Fatal("expected nil item")
	}
}

func TestQueueUpdateIndexEqualsSize(t *testing.T) {
	var queue *collections.Queue[int] = collections.NewQueue[int]()
	queue.Append(1)

	if queue.Update(queue.Size(), 5) == nil {
		t.Fatal("expected out of range exception")
	}
}

func TestQueueRemoveIndexEqualsSize(t *testing.T) {
	var queue *collections.Queue[int] = collections.NewQueue[int]()
	queue.Append(1)

	if queue.Remove(queue.Size()) == nil {
		t.Fatal("expected out of range exception")
	}
}
