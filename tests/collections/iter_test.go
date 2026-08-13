package tests

import (
	"slices"
	"testing"
	"time"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
)

func TestIterFrom(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()
	array.Append(1)
	array.Append(2)
	array.Append(3)

	var iterator *collections.Iterator[int] = collections.From(array)

	var values []int = iterator.Collect()

	var expected []int = []int{1, 2, 3}

	if !slices.Equal(values, expected) {
		test.Fatalf("expected %v, got %v", expected, values)
	}
}

func TestMap(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()

	array.Append(1)
	array.Append(2)
	array.Append(3)

	var iterator *collections.Iterator[int] = collections.From(array).Map(func(value int) int {
		return value * 2
	})

	var expected []int = []int{2, 4, 6}

	if !slices.Equal(iterator.Collect(), expected) {
		test.Fatal("unexpected mapped values")
	}
}

func TestMapEmpty(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()

	var values []int = collections.From(array).
		Map(func(value int) int {
			return value * 2
		}).
		Collect()

	if len(values) != 0 {
		test.Fatal("expected empty iterator")
	}
}

func TestFilter(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()
	var i int
	for i = 1; i <= 6; i++ {
		array.Append(i)
	}

	var values []int = collections.From(array).
		Filter(func(value int) bool {
			return value%2 == 0
		}).
		Collect()

	var expected []int = []int{2, 4, 6}

	if !slices.Equal(values, expected) {
		test.Fatalf("expected %v, got %v", expected, values)
	}
}

func TestFilterNone(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()

	array.Append(1)
	array.Append(2)

	var values []int = collections.From(array).
		Filter(func(value int) bool {
			return false
		}).
		Collect()

	if len(values) != 0 {
		test.Fatal("expected empty iterator")
	}
}

func TestFilterAll(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()

	array.Append(1)
	array.Append(2)

	var values []int = collections.From(array).
		Filter(func(value int) bool {
			return true
		}).
		Collect()

	var expected []int = []int{1, 2}

	if !slices.Equal(values, expected) {
		test.Fatal("unexpected values")
	}
}

func TestForEach(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()

	array.Append(1)
	array.Append(2)
	array.Append(3)

	var sum int = 0

	collections.From(array).ForEach(func(value int) {
		sum += value
	})

	if sum != 6 {
		test.Fatalf("expected 6, got %d", sum)
	}
}

func TestIterCollect(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()

	array.Append(10)
	array.Append(20)

	var values []int = collections.From(array).Collect()

	var expected []int = []int{10, 20}

	if !slices.Equal(values, expected) {
		test.Fatal("unexpected collected values")
	}
}

func TestIterSize(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()
	var i int
	for i = range 5 {
		array.Append(i)
	}

	if collections.From(array).Size() != 5 {
		test.Fatal("unexpected iterator size")
	}
}

func TestSizeEmpty(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()

	if collections.From(array).Size() != 0 {
		test.Fatal("expected size 0")
	}
}

func TestIterIsEmpty(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()

	if !collections.From(array).IsEmpty() {
		test.Fatal("expected empty iterator")
	}

	array.Append(1)

	if collections.From(array).IsEmpty() {
		test.Fatal("expected non-empty iterator")
	}
}

func TestEnumerate(test *testing.T) {
	var array *collections.Array[string] = collections.NewArray[string]()

	array.Append("a")
	array.Append("b")
	array.Append("c")

	var index gopolutils.Size = gopolutils.Size(0)
	var i gopolutils.Size
	var value string
	for i, value = range collections.Enumerate(array) {
		if i != index {
			test.Fatalf("expected index %d, got %d", index, i)
		} else if value != []string{"a", "b", "c"}[index] {
			test.Fatal("unexpected value")
		}
		index++
	}
	if index != 3 {
		test.Fatal("incorrect enumeration count")
	}
}

func TestIn(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()

	array.Append(1)
	array.Append(2)
	array.Append(3)

	if !collections.In(array, 2) {
		test.Fatal("expected value to exist")
	} else if collections.In(array, 5) {
		test.Fatal("did not expect value")
	}
}

func TestChain(test *testing.T) {
	var first *collections.Array[int] = collections.NewArray[int]()
	var second *collections.Array[int] = collections.NewArray[int]()
	var third *collections.Array[int] = collections.NewArray[int]()

	first.Append(1)
	first.Append(2)

	second.Append(3)

	third.Append(4)
	third.Append(5)

	var result collections.View[int] = collections.Chain(first, second, third)

	var expected []int = []int{1, 2, 3, 4, 5}

	if !slices.Equal(result.Collect(), expected) {
		test.Fatalf("expected %v, got %v", expected, result.Collect())
	}
}

func TestChainEmpty(test *testing.T) {
	var result collections.View[int] = collections.Chain[int]()

	if !result.IsEmpty() {
		test.Fatal("expected empty collection")
	}
}

func TestSumIntegers(test *testing.T) {
	if collections.Sum(1, 2, 3, 4, 5) != 15 {
		test.Fatal("incorrect sum")
	}
}

func TestSumFloats(test *testing.T) {
	if collections.Sum(1.5, 2.5, 3.0) != 7.0 {
		test.Fatal("incorrect sum")
	}
}

func TestSumEmpty(test *testing.T) {
	if collections.Sum[int]() != 0 {
		test.Fatal("expected zero")
	}
}

func TestMapFilterChain(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()
	var i int
	for i = 1; i <= 10; i++ {
		array.Append(i)
	}

	var values []int = collections.From(array).
		Map(func(value int) int {
			return value * value
		}).
		Filter(func(value int) bool {
			return value%2 == 0
		}).
		Collect()

	var expected []int = []int{4, 16, 36, 64, 100}

	if !slices.Equal(values, expected) {
		test.Fatalf("expected %v, got %v", expected, values)
	}
}

func TestReverseEmpty(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()

	defer func() {
		if recover() != nil {
			test.Fatal("Reverse panicked on an empty collection")
		}
	}()

	var count int = 0

	for range collections.Reverse(array) {
		count++
	}

	if count != 0 {
		test.Fatalf("expected 0 elements, got %d", count)
	}
}

func TestReverseOrder(test *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()

	array.Append(1)
	array.Append(2)
	array.Append(3)

	var expected []int = []int{3, 2, 1}

	var index int = 0
	var value int
	for _, value = range collections.Reverse(array) {
		if value != expected[index] {
			test.Fatalf("expected %d, got %d", expected[index], value)
		}
		index++
	}

	if index != len(expected) {
		test.Fatalf("expected %d elements, got %d", len(expected), index)
	}
}

func TestReverseIndices(test *testing.T) {
	var array *collections.Array[string] = collections.NewArray[string]()

	array.Append("a")
	array.Append("b")
	array.Append("c")

	var expectedIndices []gopolutils.Size = []gopolutils.Size{2, 1, 0}
	var expectedValues []string = []string{"c", "b", "a"}

	var index int = 0
	var i gopolutils.Size
	var value string
	for i, value = range collections.Reverse(array) {
		if i != expectedIndices[index] {
			test.Fatalf("expected index %d, got %d", expectedIndices[index], i)
		} else if value != expectedValues[index] {
			test.Fatalf("expected value %q, got %q", expectedValues[index], value)
		}

		index++
	}

	if index != len(expectedValues) {
		test.Fatalf("expected %d elements, got %d", len(expectedValues), index)
	}
}

func TestReverseTerminates(t *testing.T) {
	var array *collections.Array[int] = collections.NewArray[int]()
	array.Append(1)
	array.Append(2)
	array.Append(3)

	var done chan gopolutils.None = make(chan gopolutils.None)

	go func() {
		for range collections.Reverse(array) {
		}
		close(done)
	}()

	select {
	case <-done:
		// Success.
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Reverse did not terminate")
	}
}
