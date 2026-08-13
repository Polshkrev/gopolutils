package tests

import (
	"slices"
	"testing"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
)

func TestNewMap(t *testing.T) {
	var mapping *collections.Map[string, int] = collections.NewMap[string, int]()

	if mapping == nil {
		t.Fatal("expected map")
	} else if !mapping.IsEmpty() {
		t.Fatal("new map should be empty")
	} else if mapping.Size() != 0 {
		t.Fatalf("expected size 0, got %d", mapping.Size())
	}
}

func TestInsert(t *testing.T) {
	var mapping *collections.Map[string, int] = collections.NewMap[string, int]()
	var exception *gopolutils.Exception = mapping.Insert("one", 1)
	if exception != nil {
		t.Fatal(exception)
	} else if mapping.Size() != 1 {
		t.Fatalf("expected size 1, got %d", mapping.Size())
	}
	var value *int
	value, exception = mapping.At("one")

	if exception != nil {
		t.Fatal(exception)
	} else if *value != 1 {
		t.Fatalf("expected 1, got %d", *value)
	}
}

func TestInsertDuplicate(t *testing.T) {
	var mapping *collections.Map[string, int] = collections.NewMap[string, int]()

	var exception *gopolutils.Exception = mapping.Insert("one", 1)
	if exception != nil {
		t.Fatalf("unexpected exception: %s", exception.Error())
	} else if mapping.Insert("one", 2) == nil {
		t.Fatal("expected KeyError")
	} else if mapping.Size() != 1 {
		t.Fatal("map should not have been modified")
	}
	var value *int
	value, exception = mapping.At("one")
	if exception != nil {
		t.Fatalf("exception is not nil: %s", exception.Error())
	} else if *value != 1 {
		t.Fatal("value should not have changed")
	}
}

func TestMapAt(t *testing.T) {
	var mapping *collections.Map[string, int] = collections.NewMap[string, int]()

	var exception *gopolutils.Exception = mapping.Insert("one", 1)
	if exception != nil {
		t.Fatalf("unexpected exception: %s", exception.Error())
	}

	var value *int
	value, exception = mapping.At("one")

	if exception != nil {
		t.Fatal(exception)
	} else if *value != 1 {
		t.Fatalf("expected 1, got %d", *value)
	}
}

func TestMapAtEmpty(t *testing.T) {
	var mapping *collections.Map[string, int] = collections.NewMap[string, int]()

	var value *int
	var exception *gopolutils.Exception
	value, exception = mapping.At("one")

	if exception == nil {
		t.Fatal("expected ValueError")
	} else if value != nil {
		t.Fatal("expected nil value")
	}
}

func TestAtMissingKey(t *testing.T) {
	var mapping *collections.Map[string, int] = collections.NewMap[string, int]()

	var exception *gopolutils.Exception = mapping.Insert("one", 1)
	if exception != nil {
		t.Fatalf("unexpected exception: %s", exception.Error())
	}

	var value *int
	value, exception = mapping.At("two")

	if exception == nil {
		t.Fatal("expected KeyError")
	} else if value != nil {
		t.Fatal("expected nil value")
	}
}

func TestMapUpdate(t *testing.T) {
	var mapping *collections.Map[string, int] = collections.NewMap[string, int]()

	var exception *gopolutils.Exception = mapping.Insert("one", 1)
	if exception != nil {
		t.Fatalf("unexpected exception: %s", exception.Error())
	} else if exception = mapping.Update("one", 42); exception != nil {
		t.Fatal(exception)
	}
	var value *int
	value, exception = mapping.At("one")
	if exception != nil {
		t.Fatalf(exception.Error())
	} else if *value != 42 {
		t.Fatalf("expected 42, got %d", *value)
	}
}

func TestMapUpdateEmpty(t *testing.T) {
	var mapping *collections.Map[string, int] = collections.NewMap[string, int]()

	if mapping.Update("one", 1) == nil {
		t.Fatal("expected ValueError")
	}
}

func TestUpdateMissingKey(t *testing.T) {
	var mapping *collections.Map[string, int] = collections.NewMap[string, int]()

	var exception *gopolutils.Exception = mapping.Insert("one", 1)
	if exception != nil {
		t.Fatalf("unexpected exception: %s", exception.Error())
	} else if mapping.Update("two", 2) == nil {
		t.Fatal("expected KeyError")
	}
	var value *int
	value, exception = mapping.At("one")
	if exception != nil {
		t.Fatalf("exception is not nil: %s", exception.Error())
	} else if *value != 1 {
		t.Fatal("map should not have been modified")
	}
}

func TestKeys(t *testing.T) {
	var mapping *collections.Map[string, int] = collections.NewMap[string, int]()

	mapping.Insert("a", 1)
	mapping.Insert("b", 2)
	mapping.Insert("c", 3)

	var keys []string = mapping.Keys()

	slices.Sort(keys)

	var expected []string = []string{"a", "b", "c"}

	if !slices.Equal(keys, expected) {
		t.Fatalf("expected %v, got %v", expected, keys)
	}
}

func TestValues(t *testing.T) {
	var mapping *collections.Map[string, int] = collections.NewMap[string, int]()

	mapping.Insert("a", 3)
	mapping.Insert("b", 1)
	mapping.Insert("c", 2)

	var values []int = mapping.Values()

	slices.Sort(values)

	var expected []int = []int{1, 2, 3}

	if !slices.Equal(values, expected) {
		t.Fatalf("expected %v, got %v", expected, values)
	}
}

func TestMapRemove(t *testing.T) {
	var mapping *collections.Map[string, int] = collections.NewMap[string, int]()

	mapping.Insert("one", 1)
	mapping.Insert("two", 2)
	var exception *gopolutils.Exception = mapping.Remove("one")
	if exception != nil {
		t.Fatal(exception)
	} else if mapping.Size() != 1 {
		t.Fatalf("expected size 1, got %d", mapping.Size())
	} else if mapping.HasKey("one") {
		t.Fatal("key should have been removed")
	}
}

func TestMapRemoveEmpty(t *testing.T) {
	var mapping *collections.Map[string, int] = collections.NewMap[string, int]()

	if mapping.Remove("one") == nil {
		t.Fatal("expected ValueError")
	}
}

func TestRemoveMissingKey(t *testing.T) {
	var mapping *collections.Map[string, int] = collections.NewMap[string, int]()

	mapping.Insert("one", 1)

	if mapping.Remove("two") == nil {
		t.Fatal("expected KeyError")
	} else if !mapping.HasKey("one") {
		t.Fatal("map should not have been modified")
	}
}

func TestMapCollect(t *testing.T) {
	var mapping *collections.Map[string, int] = collections.NewMap[string, int]()

	mapping.Insert("a", 1)
	mapping.Insert("b", 2)
	mapping.Insert("c", 3)

	var pairs []collections.Pair[string, int] = mapping.Collect()

	if len(pairs) != 3 {
		t.Fatalf("expected 3 pairs, got %d", len(pairs))
	}

	var found map[string]int = map[string]int{}
	var i int
	for i = range pairs {
		var pair collections.Pair[string, int] = pairs[i]
		found[*pair.First()] = *pair.Second()
	}

	if found["a"] != 1 || found["b"] != 2 || found["c"] != 3 {
		t.Fatal("unexpected collected values")
	}
}

func TestMapIterator(t *testing.T) {
	var mapping *collections.Map[string, int] = collections.NewMap[string, int]()

	mapping.Insert("a", 1)
	mapping.Insert("b", 2)
	mapping.Insert("c", 3)

	var count int = 0

	mapping.Iterator().ForEach(func(pair collections.Pair[string, int]) {
		count++
	})

	if count != 3 {
		t.Fatalf("expected 3 pairs, got %d", count)
	}
}

func TestHasKey(t *testing.T) {
	var mapping *collections.Map[string, int] = collections.NewMap[string, int]()

	var exception *gopolutils.Exception = mapping.Insert("answer", 42)
	if exception != nil {
		t.Fatalf("unexpected exception: %s", exception.Error())
	} else if !mapping.HasKey("answer") {
		t.Fatal("expected key to exist")
	} else if mapping.HasKey("missing") {
		t.Fatal("did not expect key")
	}
}

func TestMapSize(t *testing.T) {
	var mapping *collections.Map[int, int] = collections.NewMap[int, int]()
	var i int
	for i = range 10 {
		mapping.Insert(i, i)
	}

	if mapping.Size() != 10 {
		t.Fatalf("expected size 10, got %d", mapping.Size())
	}
}

func TestMapIsEmpty(t *testing.T) {
	var mapping *collections.Map[int, int] = collections.NewMap[int, int]()

	if !mapping.IsEmpty() {
		t.Fatal("expected empty map")
	}

	var exception *gopolutils.Exception = mapping.Insert(1, 1)
	if exception != nil {
		t.Fatalf("unexpected exception: %s", exception.Error())
	} else if mapping.IsEmpty() {
		t.Fatal("expected non-empty map")
	}
}

func TestRemoveLastElement(t *testing.T) {
	var mapping *collections.Map[string, int] = collections.NewMap[string, int]()

	var exception *gopolutils.Exception = mapping.Insert("one", 1)
	if exception != nil {
		t.Fatalf("unexpected exception: %s", exception.Error())
	} else if exception := mapping.Remove("one"); exception != nil {
		t.Fatal(exception)
	} else if !mapping.IsEmpty() {
		t.Fatal("expected empty map")
	}
}

func TestInsertMany(t *testing.T) {
	var mapping *collections.Map[int, int] = collections.NewMap[int, int]()
	var i int
	for i = range 100 {
		if exception := mapping.Insert(i, i*i); exception != nil {
			t.Fatal(exception)
		}
	}

	if mapping.Size() != 100 {
		t.Fatalf("expected size 100, got %d", mapping.Size())
	}

	for i = range 100 {
		var value *int
		var exception *gopolutils.Exception
		value, exception = mapping.At(i)

		if exception != nil {
			t.Fatal(exception)
		} else if *value != i*i {
			t.Fatalf("incorrect value for key %d", i)
		}
	}
}
