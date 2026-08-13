package tests

import (
	"testing"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
)

func TestNewStack(t *testing.T) {
	var stack *collections.Stack[int] = collections.NewStack[int]()

	if stack == nil {
		t.Fatal("expected stack")
	} else if !stack.IsEmpty() {
		t.Fatal("new stack should be empty")
	} else if stack.Size() != 0 {
		t.Fatalf("expected size 0, got %d", stack.Size())
	}
}

func TestStackAppend(t *testing.T) {
	var stack *collections.Stack[int] = collections.NewStack[int]()

	stack.Append(1)
	stack.Append(2)
	stack.Append(3)

	if stack.Size() != 3 {
		t.Fatalf("expected size 3, got %d", stack.Size())
	}

	var expected []int = []int{1, 2, 3}
	var i int
	for i = range expected {
		var value int = expected[i]
		var item *int
		var exception *gopolutils.Exception
		item, exception = stack.At(gopolutils.Size(i))

		if exception != nil {
			t.Fatal(exception)
		} else if *item != value {
			t.Fatalf("expected %d, got %d", value, *item)
		}
	}
}

func TestStackExtend(t *testing.T) {
	var stack *collections.Stack[int] = collections.NewStack[int]()

	var values *collections.Array[int] = collections.NewArray[int]()
	values.Append(1)
	values.Append(2)
	values.Append(3)

	stack.Extend(values)

	if stack.Size() != 3 {
		t.Fatalf("expected size 3, got %d", stack.Size())
	}
}

func TestStackAt(t *testing.T) {
	var stack *collections.Stack[int] = collections.NewStack[int]()

	stack.Append(42)
	var item *int
	var exception *gopolutils.Exception
	item, exception = stack.At(0)
	if exception != nil {
		t.Fatal(exception)
	} else if *item != 42 {
		t.Fatalf("expected 42, got %d", *item)
	}
}

func TestStackAtEmpty(t *testing.T) {
	var stack *collections.Stack[int] = collections.NewStack[int]()

	var item *int
	var exception *gopolutils.Exception
	item, exception = stack.At(0)
	if exception == nil {
		t.Fatal("expected exception")
	} else if item != nil {
		t.Fatal("expected nil item")
	}
}

func TestStackAtOutOfRange(t *testing.T) {
	var stack *collections.Stack[int] = collections.NewStack[int]()

	stack.Append(1)

	var item *int
	var exception *gopolutils.Exception
	item, exception = stack.At(2)
	if exception == nil {
		t.Fatal("expected exception")
	} else if item != nil {
		t.Fatal("expected nil item")
	}
}

func TestStackUpdate(t *testing.T) {
	var stack *collections.Stack[int] = collections.NewStack[int]()

	stack.Append(1)

	var exception *gopolutils.Exception = stack.Update(0, 100)
	if exception != nil {
		t.Fatal(exception)
	}

	var item *int
	item, exception = stack.At(0)
	if exception != nil {
		t.Fatal(exception)
	} else if *item != 100 {
		t.Fatalf("expected 100, got %d", *item)
	}
}

func TestStackUpdateEmpty(t *testing.T) {
	var stack *collections.Stack[int] = collections.NewStack[int]()

	if stack.Update(0, 1) == nil {
		t.Fatal("expected exception")
	}
}

func TestStackUpdateOutOfRange(t *testing.T) {
	var stack *collections.Stack[int] = collections.NewStack[int]()

	stack.Append(1)

	if stack.Update(2, 5) == nil {
		t.Fatal("expected exception")
	}
}

func TestStackRemove(t *testing.T) {
	var stack *collections.Stack[int] = collections.NewStack[int]()

	stack.Append(1)
	stack.Append(2)
	stack.Append(3)
	var exception *gopolutils.Exception = stack.Remove(1)
	if exception != nil {
		t.Fatal(exception)
	}

	var expected []int = []int{1, 3}

	if stack.Size() != 2 {
		t.Fatalf("expected size 2, got %d", stack.Size())
	}
	var i int
	for i = range expected {
		var value int = expected[i]
		var item *int
		var exception *gopolutils.Exception
		item, exception = stack.At(gopolutils.Size(i))
		if exception != nil {
			t.Fatal(exception)
		} else if *item != value {
			t.Fatalf("expected %d, got %d", value, *item)
		}
	}
}

func TestStackRemoveEmpty(t *testing.T) {
	var stack *collections.Stack[int] = collections.NewStack[int]()

	if stack.Remove(0) == nil {
		t.Fatal("expected exception")
	}
}

func TestStackRemoveOutOfRange(t *testing.T) {
	var stack *collections.Stack[int] = collections.NewStack[int]()

	stack.Append(1)

	if stack.Remove(2) == nil {
		t.Fatal("expected exception")
	}
}

func TestPop(t *testing.T) {
	var stack *collections.Stack[int] = collections.NewStack[int]()

	stack.Append(1)
	stack.Append(2)
	stack.Append(3)

	var expected []int = []int{3, 2, 1}
	var i int
	for i = range expected {
		var value int = expected[i]
		var item *int
		var exception *gopolutils.Exception
		item, exception = stack.Pop()

		if exception != nil {
			t.Fatal(exception)
		} else if *item != value {
			t.Fatalf("expected %d, got %d", value, *item)
		}
	}

	if !stack.IsEmpty() {
		t.Fatal("expected empty stack")
	}
}

func TestPopEmpty(t *testing.T) {
	var stack *collections.Stack[int] = collections.NewStack[int]()

	var item *int
	var exception *gopolutils.Exception
	item, exception = stack.Pop()
	if exception == nil {
		t.Fatal("expected exception")
	} else if item != nil {
		t.Fatal("expected nil item")
	}
}

func TestStackPeek(t *testing.T) {
	var stack *collections.Stack[int] = collections.NewStack[int]()

	stack.Append(10)
	stack.Append(20)
	var item *int
	var exception *gopolutils.Exception
	item, exception = stack.Peek()

	if exception != nil {
		t.Fatal(exception)
	} else if *item != 20 {
		t.Fatalf("expected 20, got %d", *item)
	} else if stack.Size() != 2 {
		t.Fatal("peek should not modify stack")
	}
}

func TestStackPeekEmpty(t *testing.T) {
	var stack *collections.Stack[int] = collections.NewStack[int]()

	var item *int
	var exception *gopolutils.Exception
	item, exception = stack.Peek()
	if exception == nil {
		t.Fatal("expected exception")
	} else if item != nil {
		t.Fatal("expected nil item")
	}
}

func TestStackItems(t *testing.T) {
	var stack *collections.Stack[int] = collections.NewStack[int]()

	stack.Append(1)
	stack.Append(2)

	var items *[]int = stack.Items()

	(*items)[0] = 42
	var item *int
	var exception *gopolutils.Exception
	item, exception = stack.At(0)
	if exception != nil {
		t.Fatal(exception)
	} else if *item != 42 {
		t.Fatal("expected modification through Items()")
	}
}

func TestStackCollect(t *testing.T) {
	var stack *collections.Stack[int] = collections.NewStack[int]()

	stack.Append(1)
	stack.Append(2)

	var values []int = stack.Collect()

	if len(values) != 2 {
		t.Fatal("unexpected length")
	} else if values[0] != 1 || values[1] != 2 {
		t.Fatal("unexpected values")
	}
}

func TestStackIterator(t *testing.T) {
	var stack *collections.Stack[int] = collections.NewStack[int]()

	stack.Append(1)
	stack.Append(2)
	stack.Append(3)

	var count int = 0

	stack.Iterator().ForEach(func(int) {
		count++
	})

	if count != 3 {
		t.Fatalf("expected 3 elements, got %d", count)
	}
}

func TestStackSize(t *testing.T) {
	var stack *collections.Stack[int] = collections.NewStack[int]()
	var i int
	for i = range 10 {
		stack.Append(i)
	}

	if stack.Size() != 10 {
		t.Fatalf("expected size 10, got %d", stack.Size())
	}
}

func TestStackIsEmpty(t *testing.T) {
	var stack *collections.Stack[int] = collections.NewStack[int]()

	if !stack.IsEmpty() {
		t.Fatal("expected empty stack")
	}

	stack.Append(1)

	if stack.IsEmpty() {
		t.Fatal("expected non-empty stack")
	}
}

func TestStackAppendAfterPop(t *testing.T) {
	var stack *collections.Stack[int] = collections.NewStack[int]()

	stack.Append(1)
	stack.Append(2)

	stack.Pop()
	stack.Append(3)

	var expected []int = []int{1, 3}
	var i int
	for i = range expected {
		var value int = expected[i]
		var item *int
		var exception *gopolutils.Exception
		item, exception = stack.At(gopolutils.Size(i))
		if exception != nil {
			t.Fatal(exception)
		} else if *item != value {
			t.Fatalf("expected %d, got %d", value, *item)
		}
	}
}

func TestStackAtIndexEqualsSize(t *testing.T) {
	var stack *collections.Stack[int] = collections.NewStack[int]()

	stack.Append(1)

	var item *int
	var exception *gopolutils.Exception
	item, exception = stack.At(stack.Size())
	if exception == nil {
		t.Fatal("expected out of range exception")
	} else if item != nil {
		t.Fatal("expected nil item")
	}
}

func TestStackUpdateIndexEqualsSize(t *testing.T) {
	var stack *collections.Stack[int] = collections.NewStack[int]()
	stack.Append(1)

	if stack.Update(stack.Size(), 42) == nil {
		t.Fatal("expected out of range exception")
	}
}

func TestStackRemoveIndexEqualsSize(t *testing.T) {
	var stack *collections.Stack[int] = collections.NewStack[int]()
	stack.Append(1)

	if stack.Remove(stack.Size()) == nil {
		t.Fatal("expected out of range exception")
	}
}
