package tests

import (
	"testing"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
)

func TestNewArray(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()
	if array == nil {
		test.Fatal("expected array")
	} else if !array.IsEmpty() {
		test.Fatal("new array should be empty")
	} else if array.Size() != 0 {
		test.Fatalf("expected size 0, got %d", array.Size())
	}
}

func TestArrayAppend(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()
	array.Append(10)
	array.Append(20)
	array.Append(30)
	if array.Size() != 3 {
		test.Fatalf("expected size 3, got %d", array.Size())
	}
	var expected []int = []int{10, 20, 30}
	var i, value int
	for i, value = range expected {
		var item *int
		var exception *gopolutils.Exception
		item, exception = array.At(gopolutils.Size(i))
		if exception != nil {
			test.Fatal(exception)
		} else if *item != value {
			test.Fatalf("expected %d, got %d", value, *item)
		}
	}
}

func TestArrayExtend(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()

	var other *collections.Array[int] = collections.NewArray[int]()
	other.Append(1)
	other.Append(2)
	other.Append(3)

	array.Extend(other)

	if array.Size() != 3 {
		test.Fatalf("expected size 3, got %d", array.Size())
	}

	var expected []int = []int{1, 2, 3}

	var i, value int
	for i, value = range expected {
		var item *int
		var exception *gopolutils.Exception
		item, exception = array.At(gopolutils.Size(i))
		if exception != nil {
			test.Fatalf("exception is not nil: %s", exception.Error())
		} else if *item != value {
			test.Fatalf("expected %d, got %d", value, *item)
		}
	}
}

func TestArrayAt(test *testing.T) {
	var array *collections.Array[string] = collections.NewArray[string]()

	array.Append("hello")

	var item *string
	var exception *gopolutils.Exception
	item, exception = array.At(0)

	if exception != nil {
		test.Fatal(exception)
	} else if *item != "hello" {
		test.Fatalf("expected hello, got %s", *item)
	}
}

func TestArrayAtEmpty(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()

	var item *int
	var exception *gopolutils.Exception
	item, exception = array.At(0)

	if exception == nil {
		test.Fatal("expected exception")
	} else if item != nil {
		test.Fatal("expected nil item")
	}
}

func TestArrayAtOutOfRange(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()

	array.Append(1)

	var item *int
	var exception *gopolutils.Exception
	item, exception = array.At(2)

	if exception == nil {
		test.Fatal("expected exception")
	} else if item != nil {
		test.Fatal("expected nil item")
	}
}

func TestArrayUpdate(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()

	array.Append(5)

	var exception *gopolutils.Exception = array.Update(0, 42)

	if exception != nil {
		test.Fatal(exception)
	}
	var item *int
	item, exception = array.At(0)
	if exception != nil {
		test.Fatalf("exception is not nil: %s", exception.Error())
	} else if *item != 42 {
		test.Fatalf("expected 42, got %d", *item)
	}
}

func TestArrayUpdateEmpty(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()

	if array.Update(0, 10) == nil {
		test.Fatal("expected exception")
	}
}

func TestArrayUpdateOutOfRange(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()

	array.Append(1)

	if array.Update(5, 10) == nil {
		test.Fatal("expected exception")
	}

	var item *int
	var exception *gopolutils.Exception
	item, exception = array.At(0)
	if exception != nil {
		test.Fatalf("exception is not nil: %s", exception.Error())
	} else if *item != 1 {
		test.Fatal("array should not have been modified")
	}
}

func TestArrayRemove(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()

	array.Append(1)
	array.Append(2)
	array.Append(3)

	var exception *gopolutils.Exception = array.Remove(1)

	if exception != nil {
		test.Fatal(exception)
	}

	var expected []int = []int{1, 3}

	if array.Size() != 2 {
		test.Fatalf("expected size 2, got %d", array.Size())
	}
	var i, value int
	for i, value = range expected {
		var item *int
		var exception *gopolutils.Exception
		item, exception = array.At(gopolutils.Size(i))
		if exception != nil {
			test.Fatalf("exception is not nil: %s", exception.Error())
		} else if *item != value {
			test.Fatalf("expected %d, got %d", value, *item)
		}
	}
}

func TestArrayRemoveEmpty(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()

	if array.Remove(0) == nil {
		test.Fatal("expected exception")
	}
}

func TestArrayRemoveOutOfRange(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()

	array.Append(1)

	if array.Remove(5) == nil {
		test.Fatal("expected exception")
	} else if array.Size() != 1 {
		test.Fatal("array should not have been modified")
	}
}

func TestArrayArrayItems(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()

	array.Append(1)
	array.Append(2)

	var items *[]int = array.Items()

	if len(*items) != 2 {
		test.Fatalf("expected length 2, got %d", len(*items))
	}

	(*items)[0] = 99
	var item *int
	var exception *gopolutils.Exception
	item, exception = array.At(0)
	if exception != nil {
		test.Fatalf("exception is not nil: %s", exception.Error())
	} else if *item != 99 {
		test.Fatal("expected modification through Items()")
	}
}

func TestArrayCollect(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()

	array.Append(1)
	array.Append(2)

	var items []int = array.Collect()

	if len(items) != 2 {
		test.Fatal("unexpected length")
	} else if items[0] != 1 || items[1] != 2 {
		test.Fatal("unexpected contents")
	}
}

func TestArrayIterator(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()

	array.Append(1)
	array.Append(2)
	array.Append(3)

	var iterator *collections.Iterator[int] = array.Iterator()

	var count int = 0

	iterator.ForEach(func(i int) {
		count += i
	})

	if count != 6 {
		test.Fatalf("expected 3 elements, got %d", count)
	}
}

func TestArraySize(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()

	var i int
	for i = range 10 {
		array.Append(i)
	}

	if array.Size() != 10 {
		test.Fatalf("expected size 10, got %d", array.Size())
	}
}

func TestArrayIsEmpty(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()

	if !array.IsEmpty() {
		test.Fatal("expected empty array")
	}

	array.Append(1)

	if array.IsEmpty() {
		test.Fatal("expected non-empty array")
	}
}

func TestArrayRemoveLastElement(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()

	array.Append(42)

	var exception *gopolutils.Exception = array.Remove(0)
	if exception != nil {
		test.Fatal(exception)
	} else if !array.IsEmpty() {
		test.Fatal("expected empty array")
	}
}

func TestArrayAppendAfterRemove(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()

	array.Append(1)
	array.Append(2)

	array.Remove(0)
	array.Append(3)

	var expected []int = []int{2, 3}
	var i, value int
	for i, value = range expected {
		var item *int
		var exception *gopolutils.Exception
		item, exception = array.At(gopolutils.Size(i))
		if exception != nil {
			test.Fatalf("exception is not nil: %s", exception.Error())
		} else if *item != value {
			test.Fatalf("expected %d, got %d", value, *item)
		}
	}
}

func TestArrayAtIndexEqualsSize(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()
	array.Append(1)
	var item *int
	var exception *gopolutils.Exception
	item, exception = array.At(array.Size())

	if exception == nil {
		test.Fatal("expected out of range exception")
	} else if item != nil {
		test.Fatal("expected nil item")
	}
}

func TestArrayUpdateIndexEqualsSize(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()
	array.Append(1)

	if array.Update(array.Size(), 5) == nil {
		test.Fatal("expected out of range exception")
	}
}

func TestArrayRemoveIndexEqualsSize(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()
	array.Append(1)

	if array.Remove(array.Size()) == nil {
		test.Fatal("expected out of range exception")
	}
}
