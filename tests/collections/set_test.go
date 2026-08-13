package tests

import (
	"slices"
	"strings"
	"testing"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
)

func TestNewSet(t *testing.T) {
	var set *collections.Set[int] = collections.NewSet[int]()

	if set == nil {
		t.Fatal("expected set")
	} else if !set.IsEmpty() {
		t.Fatal("new set should be empty")
	} else if set.Size() != 0 {
		t.Fatalf("expected size 0, got %d", set.Size())
	}
}

func TestSetAppend(t *testing.T) {
	var set *collections.Set[int] = collections.NewSet[int]()

	set.Append(1)
	set.Append(2)
	set.Append(3)

	if set.Size() != 3 {
		t.Fatalf("expected size 3, got %d", set.Size())
	}
	var expected []int = []int{1, 2, 3}
	var i int
	for i = range expected {
		var value int = expected[i]
		if !set.Contains(value) {
			t.Fatalf("expected set to contain %d", value)
		}
	}
}

func TestAppendDuplicate(t *testing.T) {
	var set *collections.Set[int] = collections.NewSet[int]()

	set.Append(1)
	set.Append(1)
	set.Append(1)

	if set.Size() != 1 {
		t.Fatalf("expected size 1, got %d", set.Size())
	}
}

func TestSetExtend(t *testing.T) {
	var values *collections.Array[int] = collections.NewArray[int]()
	values.Append(1)
	values.Append(2)
	values.Append(2)
	values.Append(3)

	var set *collections.Set[int] = collections.NewSet[int]()
	set.Extend(values)

	if set.Size() != 3 {
		t.Fatalf("expected size 3, got %d", set.Size())
	}
}

func TestSetAtEmpty(t *testing.T) {
	var set *collections.Set[int] = collections.NewSet[int]()
	var exception *gopolutils.Exception
	_, exception = set.At(0)

	if exception == nil {
		t.Fatal("expected NotImplementedError")
	}
}

// func TestSetAt(t *testing.T) {
// 	var set *collections.Set[int] = collections.NewSet[int]()
// 	set.Append(1)
// 	var value *int
// 	var exception *gopolutils.Exception
// 	value, exception = set.At(0)
// 	if exception != nil {
// 		t.Fatal(exception)
// 	} else if *value != 1 {
// 		t.Fatalf("expected value %d, got %d", 1, *value)
// 	} else if exception == nil {
// 		t.Fatal("expected NotImplementedError")
// 	}
// }

func TestSetUpdate(t *testing.T) {
	var set *collections.Set[int] = collections.NewSet[int]()

	if set.Update(0, 1) == nil {
		t.Fatal("expected NotImplementedError")
	}
}

func TestSetRemove(t *testing.T) {
	var set *collections.Set[int] = collections.NewSet[int]()

	set.Append(1)
	set.Append(2)
	set.Append(3)

	var size gopolutils.Size = set.Size()
	var exception *gopolutils.Exception = set.Remove(1)
	if exception != nil {
		t.Fatal(exception)
	} else if set.Size() != size-1 {
		t.Fatal("remove failed")
	}
}

func TestSetRemoveEmpty(t *testing.T) {
	var set *collections.Set[int] = collections.NewSet[int]()

	if set.Remove(0) == nil {
		t.Fatal("expected ValueError")
	}
}

func TestDiscard(t *testing.T) {
	var set *collections.Set[int] = collections.NewSet[int]()

	set.Append(1)
	set.Append(2)

	set.Discard(1)

	if set.Contains(1) {
		t.Fatal("discard failed")
	} else if set.Size() != 1 {
		t.Fatal("unexpected size")
	}
}

func TestDiscardMissing(t *testing.T) {
	var set *collections.Set[int] = collections.NewSet[int]()

	set.Append(1)

	set.Discard(2)

	if set.Size() != 1 {
		t.Fatal("set should not have changed")
	}
}

func TestContains(t *testing.T) {
	var set *collections.Set[int] = collections.NewSet[int]()

	set.Append(10)

	if !set.Contains(10) {
		t.Fatal("expected value")
	} else if set.Contains(20) {
		t.Fatal("unexpected value")
	}
}

func TestSetItems(t *testing.T) {
	var set *collections.Set[int] = collections.NewSet[int]()

	set.Append(1)
	set.Append(2)
	set.Append(3)

	var items *[]int = set.Items()

	slices.Sort(*items)

	var expected []int = []int{1, 2, 3}

	if !slices.Equal(*items, expected) {
		t.Fatalf("expected %v, got %v", expected, items)
	}
}

func TestDifference(t *testing.T) {
	var left *collections.Set[int] = collections.NewSet[int]()
	var right *collections.Set[int] = collections.NewSet[int]()

	left.Append(1)
	left.Append(2)

	right.Append(2)
	right.Append(3)
	right.Append(4)

	var result *collections.Set[int] = left.Difference(*right)

	var expected []int = []int{3, 4}

	var values []int = result.Collect()
	slices.Sort(values)

	if !slices.Equal(values, expected) {
		t.Fatalf("expected %v, got %v", expected, values)
	}
}

func TestIntersection(t *testing.T) {
	var left *collections.Set[int] = collections.NewSet[int]()
	var right *collections.Set[int] = collections.NewSet[int]()

	left.Append(1)
	left.Append(2)
	left.Append(3)

	right.Append(2)
	right.Append(3)
	right.Append(4)

	var result *collections.Set[int] = left.Intersection(*right)

	var values []int = result.Collect()
	slices.Sort(values)

	var expected []int = []int{2, 3}

	if !slices.Equal(values, expected) {
		t.Fatalf("expected %v, got %v", expected, values)
	}
}

func TestSetCollect(t *testing.T) {
	var set *collections.Set[int] = collections.NewSet[int]()

	set.Append(3)
	set.Append(1)
	set.Append(2)

	var values []int = set.Collect()

	slices.Sort(values)

	var expected []int = []int{1, 2, 3}

	if !slices.Equal(values, expected) {
		t.Fatalf("expected %v, got %v", expected, values)
	}
}

func TestFrom(t *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()
	array.Append(1)
	array.Append(2)
	array.Append(2)
	array.Append(3)

	var set *collections.Set[int] = collections.NewSet[int]()
	set.From(array)

	if set.Size() != 3 {
		t.Fatalf("expected size 3, got %d", set.Size())
	}
}

func TestInto(t *testing.T) {
	var set *collections.Set[int] = collections.NewSet[int]()

	set.Append(1)
	set.Append(2)
	set.Append(3)

	var collection collections.Collection[int] = set.Into()

	var values []int = collection.Collect()

	slices.Sort(values)

	var expected []int = []int{1, 2, 3}

	if !slices.Equal(values, expected) {
		t.Fatalf("expected %v, got %v", expected, values)
	}
}

func TestSetIterator(t *testing.T) {
	var set *collections.Set[int] = collections.NewSet[int]()

	set.Append(1)
	set.Append(2)
	set.Append(3)

	var count int = 0

	set.Iterator().ForEach(func(int) {
		count++
	})

	if count != 3 {
		t.Fatalf("expected 3 elements, got %d", count)
	}
}

func TestString(t *testing.T) {
	var set *collections.Set[int] = collections.NewSet[int]()

	set.Append(1)
	set.Append(2)
	set.Append(3)

	var value string = set.String()

	if !strings.HasPrefix(value, "{") {
		t.Fatal("missing opening brace")
	} else if !strings.HasSuffix(value, "}") {
		t.Fatal("missing closing brace")
	}
	var expected []string = []string{"1", "2", "3"}
	var i int
	for i = range expected {
		var value string = expected[i]
		if !strings.Contains(value, value) {
			t.Fatalf("missing value %s, full: %s", value, value)
		}
	}
}

func TestSetIsEmpty(t *testing.T) {
	var set *collections.Set[int] = collections.NewSet[int]()

	if !set.IsEmpty() {
		t.Fatal("expected empty set")
	}

	set.Append(1)

	if set.IsEmpty() {
		t.Fatal("expected non-empty set")
	}
}

func TestSetRemoveIndexEqualsSize(t *testing.T) {
	var set *collections.Set[int] = collections.NewSet[int]()

	set.Append(1)

	if set.Remove(set.Size()) == nil {
		t.Fatal("expected exception")
	}
}
