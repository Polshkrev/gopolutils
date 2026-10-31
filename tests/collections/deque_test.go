package tests

import (
	"testing"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
)

func TestNewDeque(test *testing.T) {
	var list *collections.Deque[int] = collections.NewDeque[int]()

	if list == nil {
		test.Fatal("NewDeque() returned nil")
	} else if !list.IsEmpty() {
		test.Error("NewDeque() should return an empty deque")
	}
}

func TestDequeUseEmpty(test *testing.T) {
	var list *collections.Deque[int] = collections.NewDeque[int]()

	var exception *gopolutils.Exception = list.Use(0)

	if exception == nil {
		test.Fatal("Use(0) should return an exception for an empty deque")
	} else if !list.IsEmpty() {
		test.Error("Use() should not modify an empty deque")
	}
}

func TestDequeUseInvalidIndex(test *testing.T) {
	var list *collections.Deque[int] = collections.NewDeque[int]()

	list.Append(10)
	list.Append(20)
	list.Append(30)

	var exception *gopolutils.Exception = list.Use(list.Size())
	if exception == nil {
		test.Fatal("Use(size) should return an exception")
	}

	var expected []int = []int{10, 20, 30}
	assertDequeValues(test, list, expected)
}

func TestDequeUseFirst(test *testing.T) {
	var list *collections.Deque[int] = collections.NewDeque[int]()

	list.Append(10)
	list.Append(20)
	list.Append(30)

	var exception *gopolutils.Exception = list.Use(0)
	if exception != nil {
		test.Fatalf("Use(0) returned an unexpected exception: %v", exception)
	}

	var expected []int = []int{10, 20, 30}
	assertDequeValues(test, list, expected)
}

func TestDequeUseMiddle(test *testing.T) {
	var list *collections.Deque[int] = collections.NewDeque[int]()

	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Append(40)

	var exception *gopolutils.Exception = list.Use(2)
	if exception != nil {
		test.Fatalf("Use(2) returned an unexpected exception: %v", exception)
	}

	var expected []int = []int{30, 10, 20, 40}
	assertDequeValues(test, list, expected)
}

func TestDequeUseLast(test *testing.T) {
	var list *collections.Deque[int] = collections.NewDeque[int]()

	list.Append(10)
	list.Append(20)
	list.Append(30)

	var exception *gopolutils.Exception = list.Use(2)
	if exception != nil {
		test.Fatalf("Use(2) returned an unexpected exception: %v", exception)
	}

	var expected []int = []int{30, 10, 20}
	assertDequeValues(test, list, expected)
}

func TestDequeUseSingleElement(test *testing.T) {
	var list *collections.Deque[int] = collections.NewDeque[int]()

	list.Append(10)

	var exception *gopolutils.Exception = list.Use(0)
	if exception != nil {
		test.Fatalf("Use(0) returned an unexpected exception: %v", exception)
	}

	var expected []int = []int{10}
	assertDequeValues(test, list, expected)
}

func TestDequeUseLinks(test *testing.T) {
	var list *collections.Deque[int] = collections.NewDeque[int]()

	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Append(40)

	var exception *gopolutils.Exception = list.Use(2)

	if exception != nil {
		test.Fatalf("Use(2) returned an unexpected exception: %v", exception)
	} else if list.Head().HasPrevious() {
		test.Error("head should not have a previous node")
	} else if list.Tail().HasNext() {
		test.Error("tail should not have a next node")
	}

	var forward []int

	var node *collections.Node[int] = list.Head()

	for {
		forward = append(forward, node.Data())

		if !node.HasNext() {
			break
		}

		var next collections.Node[int]
		var next_exception *gopolutils.Exception

		next, next_exception = node.Next()

		if next_exception != nil {
			test.Fatalf("Next() returned an unexpected exception: %v", next_exception)
		}

		node = &next
	}

	var expected_forward []int = []int{30, 10, 20, 40}

	assertIntSlicesEqual(test, forward, expected_forward)

	var backward []int

	node = list.Tail()

	for {
		backward = append(backward, node.Data())

		if !node.HasPrevious() {
			break
		}

		var previous collections.Node[int]
		var previous_exception *gopolutils.Exception

		previous, previous_exception = node.Previous()

		if previous_exception != nil {
			test.Fatalf("Previous() returned an unexpected exception: %v", previous_exception)
		}

		node = &previous
	}

	var expected_backward []int = []int{40, 20, 10, 30}

	assertIntSlicesEqual(test, backward, expected_backward)
}

func TestDequeUseRepeated(test *testing.T) {
	var list *collections.Deque[int] = collections.NewDeque[int]()

	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Append(40)

	var exception *gopolutils.Exception
	if exception = list.Use(2); exception != nil {
		test.Fatalf("Use(2) returned an unexpected exception: %v", exception)
	} else if exception = list.Use(3); exception != nil {
		test.Fatalf("Use(3) returned an unexpected exception: %v", exception)
	} else if exception = list.Use(2); exception != nil {
		test.Fatalf("Use(2) returned an unexpected exception: %v", exception)
	}

	var expected []int = []int{10, 40, 30, 20}
	assertDequeValues(test, list, expected)
}

func assertDequeValues(test *testing.T, list *collections.Deque[int], expected []int) {
	test.Helper()
	var forward []int
	var node *collections.Node[int] = list.Head()
	for node != nil {
		forward = append(forward, node.Data())

		if !node.HasNext() {
			break
		}

		var next collections.Node[int]
		var exception *gopolutils.Exception
		next, exception = node.Next()
		if exception != nil {
			test.Fatalf("Next() returned an unexpected exception: %v", exception)
		}
		node = &next
	}
	assertIntSlicesEqual(test, forward, expected)
}

func assertIntSlicesEqual(test *testing.T, actual []int, expected []int) {
	test.Helper()
	if len(actual) != len(expected) {
		test.Fatalf("slice length = %d, want %d", len(actual), len(expected))
	}
	var index int
	for index = range expected {
		var value int = expected[index]
		if actual[index] != value {
			test.Errorf("element at index %d = %d, want %d", index, actual[index], value)
		}
	}
}
