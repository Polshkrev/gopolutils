package tests

import (
	"testing"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
)

func TestNewDoublyLinkedList(test *testing.T) {
	var list *collections.DoublyLinkedList[int] = collections.NewDoublyLinkedList[int]()

	if list == nil {
		test.Fatal("NewDoublyLinkedList() returned nil")
	} else if !list.IsEmpty() {
		test.Error("IsEmpty() = false, want true")
	} else if list.Size() != 0 {
		test.Errorf("Size() = %d, want 0", list.Size())
	} else if list.HasHead() {
		test.Error("HasHead() = true, want false")
	} else if list.HasTail() {
		test.Error("HasTail() = true, want false")
	}
}

func TestAppend(test *testing.T) {
	var list *collections.DoublyLinkedList[int] = collections.NewDoublyLinkedList[int]()

	list.Append(10)

	if list.IsEmpty() {
		test.Fatal("IsEmpty() = true after Append(), want false")
	} else if list.Size() != 1 {
		test.Errorf("Size() = %d, want 1", list.Size())
	} else if list.Head() != list.Tail() {
		test.Error("head and tail should reference the same node after one append")
	}

	list.Append(20)
	list.Append(30)

	var expected []int = []int{10, 20, 30}

	assertListContents(test, list, expected)
	assertListLinks(test, list)
}

func TestAppendMaintainsBidirectionalLinks(test *testing.T) {
	var list *collections.DoublyLinkedList[int] = collections.NewDoublyLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)

	assertListLinks(test, list)
}

type atTestCase struct {
	name  string
	index gopolutils.Size
	want  string
}

func TestAt(test *testing.T) {
	var list *collections.DoublyLinkedList[string] = collections.NewDoublyLinkedList[string]()

	list.Append("first")
	list.Append("second")
	list.Append("third")

	var testCases []atTestCase = []atTestCase{
		{
			name:  "first element",
			index: 0,
			want:  "first",
		},
		{
			name:  "middle element",
			index: 1,
			want:  "second",
		},
		{
			name:  "last element",
			index: 2,
			want:  "third",
		},
	}

	var index int
	for index = range testCases {
		var testCase atTestCase = testCases[index]
		test.Run(testCase.name, func(test *testing.T) {
			var actual *string
			var exception *gopolutils.Exception

			actual, exception = list.At(testCase.index)

			if exception != nil {
				test.Fatalf("At(%d) returned an unexpected exception: %v", testCase.index, exception)
			} else if actual == nil {
				test.Fatal("At() returned a nil data pointer")
			} else if *actual != testCase.want {
				test.Errorf("At(%d) = %q, want %q", testCase.index, *actual, testCase.want)
			}
		})
	}
}

func TestAtReturnsMutablePointer(test *testing.T) {
	var list *collections.DoublyLinkedList[int] = collections.NewDoublyLinkedList[int]()

	list.Append(10)

	var actual *int
	var exception *gopolutils.Exception

	actual, exception = list.At(0)

	if exception != nil {
		test.Fatalf("At(0) returned an unexpected exception: %v", exception)
	} else if actual == nil {
		test.Fatal("At(0) returned a nil data pointer")
	}

	*actual = 99

	var updated *int

	updated, exception = list.At(0)

	if exception != nil {
		test.Fatalf("At(0) returned an unexpected exception after mutation: %v", exception)
	} else if updated == nil {
		test.Fatal("At(0) returned a nil data pointer after mutation")
	} else if *updated != 99 {
		test.Errorf("stored value = %d, want 99", *updated)
	}
}

type validTestCase struct {
	name  string
	list  *collections.DoublyLinkedList[int]
	index gopolutils.Size
}

func TestAtInvalidIndex(test *testing.T) {
	var testCases []validTestCase = []validTestCase{
		{
			name:  "empty list",
			list:  collections.NewDoublyLinkedList[int](),
			index: 0,
		},
		{
			name:  "out of range",
			list:  collections.NewDoublyLinkedList[int](),
			index: 1,
		},
	}

	testCases[1].list.Append(10)

	var index int
	for index = range testCases {
		var testCase validTestCase = testCases[index]
		test.Run(testCase.name, func(test *testing.T) {
			var actual *int
			var exception *gopolutils.Exception

			actual, exception = testCase.list.At(testCase.index)

			if exception == nil {
				test.Fatal("At() returned no exception for an invalid index")
			}

			if actual != nil {
				test.Error("At() returned a non-nil pointer for an invalid index")
			}
		})
	}
}

func TestCollect(test *testing.T) {
	var list *collections.DoublyLinkedList[int] = collections.NewDoublyLinkedList[int]()

	var expected []int = []int{}

	assertSlicesEqual(test, list.Collect(), expected)

	list.Append(10)
	list.Append(20)
	list.Append(30)

	expected = []int{10, 20, 30}

	assertSlicesEqual(test, list.Collect(), expected)
}

func TestDoublyLinkedItems(test *testing.T) {
	var list *collections.DoublyLinkedList[int] = collections.NewDoublyLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)

	var items *[]int = list.Items()

	if items == nil {
		test.Fatal("Items() returned a nil slice pointer")
	}

	var expected []int = []int{1, 2, 3}

	assertSlicesEqual(test, *items, expected)

	(*items)[0] = 100

	assertListContents(test, list, expected)
}

func TestUpdate(test *testing.T) {
	var list *collections.DoublyLinkedList[int] = collections.NewDoublyLinkedList[int]()

	list.Append(10)
	list.Append(20)
	list.Append(30)

	var exception *gopolutils.Exception = list.Update(1, 99)

	if exception != nil {
		test.Fatalf("Update(1, 99) returned an unexpected exception: %v", exception)
	}

	var expected []int = []int{10, 99, 30}

	assertListContents(test, list, expected)
	assertListLinks(test, list)
}

func TestUpdateInvalidIndex(test *testing.T) {
	var testCases []validTestCase = []validTestCase{
		{
			name:  "empty list",
			list:  collections.NewDoublyLinkedList[int](),
			index: 0,
		},
		{
			name:  "out of range",
			list:  collections.NewDoublyLinkedList[int](),
			index: 1,
		},
	}

	testCases[1].list.Append(10)

	var index int
	for index = range testCases {
		var testCase validTestCase = testCases[index]
		test.Run(testCase.name, func(test *testing.T) {
			var exception *gopolutils.Exception

			exception = testCase.list.Update(testCase.index, 99)

			if exception == nil {
				test.Fatal("Update() returned no exception for an invalid index")
			}

			assertListContents(test, testCase.list, testCase.list.Collect())
		})
	}
}

type removeTestCase struct {
	name          string
	index         gopolutils.Size
	expectedItems []int
}

func TestRemove(test *testing.T) {
	var testCases []removeTestCase = []removeTestCase{
		{
			name:          "remove first",
			index:         0,
			expectedItems: []int{20, 30, 40},
		},
		{
			name:          "remove middle",
			index:         1,
			expectedItems: []int{10, 30, 40},
		},
		{
			name:          "remove last",
			index:         3,
			expectedItems: []int{10, 20, 30},
		},
	}

	var index int
	for index = range testCases {
		var testCase removeTestCase = testCases[index]
		test.Run(testCase.name, func(test *testing.T) {
			var list *collections.DoublyLinkedList[int] = collections.NewDoublyLinkedList[int]()

			list.Append(10)
			list.Append(20)
			list.Append(30)
			list.Append(40)

			var exception *gopolutils.Exception = list.Remove(testCase.index)

			if exception != nil {
				test.Fatalf("Remove(%d) returned an unexpected exception: %v", testCase.index, exception)
			}

			assertListContents(test, list, testCase.expectedItems)
			assertListLinks(test, list)
		})
	}
}

func TestRemoveOnlyElement(test *testing.T) {
	var list *collections.DoublyLinkedList[int] = collections.NewDoublyLinkedList[int]()

	list.Append(10)

	var exception *gopolutils.Exception = list.Remove(0)

	if exception != nil {
		test.Fatalf("Remove(0) returned an unexpected exception: %v", exception)
	}

	assertListContents(test, list, []int{})

	if list.HasHead() {
		test.Error("HasHead() = true after removing the only element")
	} else if list.HasTail() {
		test.Error("HasTail() = true after removing the only element")
	}

	assertListLinks(test, list)
}

func TestRemoveInvalidIndex(test *testing.T) {
	var testCases []validTestCase = []validTestCase{
		{
			name:  "empty list",
			list:  collections.NewDoublyLinkedList[int](),
			index: 0,
		},
		{
			name:  "out of range",
			list:  collections.NewDoublyLinkedList[int](),
			index: 2,
		},
	}

	testCases[1].list.Append(10)
	testCases[1].list.Append(20)

	var index int
	for index = range testCases {
		var testCase validTestCase = testCases[index]
		test.Run(testCase.name, func(test *testing.T) {
			var expected []int = testCase.list.Collect()

			var exception *gopolutils.Exception = testCase.list.Remove(testCase.index)

			if exception == nil {
				test.Fatal("Remove() returned no exception for an invalid index")
			}

			assertListContents(test, testCase.list, expected)
			assertListLinks(test, testCase.list)
		})
	}
}

type stringTestCase struct {
	name string
	want string
}

func TestDoublyLinkedString(test *testing.T) {
	var list *collections.DoublyLinkedList[int] = collections.NewDoublyLinkedList[int]()

	var testCases []stringTestCase = []stringTestCase{
		{
			name: "empty list",
			want: "[]",
		},
		{
			name: "populated list",
			want: "[10, 20, 30]",
		},
	}

	var index int
	for index = range testCases {
		var testCase stringTestCase = testCases[index]
		test.Run(testCase.name, func(test *testing.T) {
			if index == 1 {
				list.Append(10)
				list.Append(20)
				list.Append(30)
			}

			var actual string = list.String()

			if actual != testCase.want {
				test.Errorf("String() = %q, want %q", actual, testCase.want)
			}
		})
	}
}

func TestHasHeadAndTail(test *testing.T) {
	var list *collections.DoublyLinkedList[int] = collections.NewDoublyLinkedList[int]()

	if list.HasHead() || list.HasTail() {
		test.Fatal("empty list unexpectedly has a head or tail")
	}

	list.Append(10)

	if !list.HasHead() || !list.HasTail() {
		test.Fatal("single-element list should have both a head and tail")
	}

	list.Append(20)

	if !list.HasHead() || !list.HasTail() {
		test.Error("nonempty list should have both a head and tail")
	}
}

func assertListContents[Type comparable](test *testing.T, list *collections.DoublyLinkedList[Type], expected []Type) {
	test.Helper()

	assertSlicesEqual(test, list.Collect(), expected)

	if list.Size() != gopolutils.Size(len(expected)) {
		test.Errorf("Size() = %d, want %d", list.Size(), len(expected))
	} else if list.IsEmpty() != (len(expected) == 0) {
		test.Errorf("IsEmpty() = %test, want %test", list.IsEmpty(), len(expected) == 0)
	}
}

func assertSlicesEqual[Type comparable](test *testing.T, actual []Type, expected []Type) {
	test.Helper()

	if len(actual) != len(expected) {
		test.Fatalf("slice length = %d, want %d", len(actual), len(expected))
	}

	for index, expectedValue := range expected {
		if actual[index] != expectedValue {
			test.Errorf("element at index %d = %v, want %v", index, actual[index], expectedValue)
		}
	}
}

func assertListLinks[Type comparable](test *testing.T, list *collections.DoublyLinkedList[Type]) {
	test.Helper()

	if list.IsEmpty() {
		if list.HasHead() {
			test.Error("empty list has a head")
		} else if list.HasTail() {
			test.Error("empty list has a tail")
		}
		return
	} else if !list.HasHead() {
		test.Fatal("nonempty list has no head")
	} else if !list.HasTail() {
		test.Fatal("nonempty list has no tail")
	}

	var expected []Type = list.Collect()
	var forward []Type
	var backward []Type

	var node *collections.Node[Type] = list.Head()

	var index gopolutils.Size

	for index = gopolutils.Size(0); index < list.Size(); index++ {
		if node == nil {
			test.Fatalf("forward traversal ended at index %d", index)
		}

		forward = append(forward, node.Data())

		if index+1 == list.Size() {
			if node.HasNext() {
				test.Error("last node has an unexpected next node")
			}
			break
		} else if !node.HasNext() {
			test.Fatalf("node at index %d has no next node", index)
		}

		var next collections.Node[Type]
		var exception *gopolutils.Exception

		next, exception = node.Next()

		if exception != nil {
			test.Fatalf("Next() returned an unexpected exception: %v", exception)
		}
		node = &next
	}

	assertSlicesEqual(test, forward, expected)

	node = list.Tail()

	for index = gopolutils.Size(0); index < list.Size(); index++ {
		if node == nil {
			test.Fatalf("backward traversal ended at index %d", index)
		}

		backward = append(backward, node.Data())

		if index+1 == list.Size() {
			if node.HasPrevious() {
				test.Error("first node has an unexpected previous node")
			}
			break
		} else if !node.HasPrevious() {
			test.Fatalf("node at backward position %d has no previous node", index)
		}

		var previous collections.Node[Type]
		var exception *gopolutils.Exception

		previous, exception = node.Previous()

		if exception != nil {
			test.Fatalf("Previous() returned an unexpected exception: %v", exception)
		}

		node = &previous
	}

	var expectedBackward []Type = make([]Type, len(expected))
	var expectedIndex int
	for expectedIndex = range expected {
		expectedBackward[expectedIndex] = expected[len(expected)-1-expectedIndex]
	}

	assertSlicesEqual(test, backward, expectedBackward)
}
