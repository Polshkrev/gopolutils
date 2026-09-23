package tests

import (
	"testing"

	"github.com/Polshkrev/gopolutils/collections"
)

func TestNewPair(t *testing.T) {
	var pair *collections.Pair[int, string] = collections.NewPair(1, "hello")

	if pair == nil {
		t.Fatal("expected pair")
	} else if pair.First() != 1 {
		t.Fatalf("expected first to be 1, got %d", pair.First())
	} else if pair.Second() != "hello" {
		t.Fatalf("expected second to be hello, got %s", pair.Second())
	}
}

func TestFirst(t *testing.T) {
	var pair *collections.Pair[int, string] = collections.NewPair(42, "value")

	var first int = pair.First()

	if first != 42 {
		t.Fatalf("expected 42, got %d", first)
	}
}

func TestSecond(t *testing.T) {
	var pair *collections.Pair[string, int] = collections.NewPair("key", 10)

	var second int = pair.Second()

	if second != 10 {
		t.Fatalf("expected 10, got %d", second)
	}
}

func TestSetFirst(t *testing.T) {
	var pair *collections.Pair[int, int] = collections.NewPair(1, 2)

	pair.SetFirst(10)

	if pair.First() != 10 {
		t.Fatalf("expected 10, got %d", pair.First())
	} else if pair.Second() != 2 {
		t.Fatal("second should not have changed")
	}
}

func TestSetSecond(t *testing.T) {
	var pair *collections.Pair[int, int] = collections.NewPair(1, 2)

	pair.SetSecond(20)

	if pair.Second() != 20 {
		t.Fatalf("expected 20, got %d", pair.Second())
	} else if pair.First() != 1 {
		t.Fatal("first should not have changed")
	}
}

func TestSet(t *testing.T) {
	var pair *collections.Pair[int, int] = collections.NewPair(1, 2)

	pair.Set(10, 20)

	if pair.First() != 10 {
		t.Fatalf("expected 10, got %d", pair.First())
	} else if pair.Second() != 20 {
		t.Fatalf("expected 20, got %d", pair.Second())
	}
}

func TestSwap(t *testing.T) {
	var first *collections.Pair[int, string] = collections.NewPair(1, "one")
	var second *collections.Pair[int, string] = collections.NewPair(2, "two")

	first.Swap(second)

	if first.First() != 2 || first.Second() != "two" {
		t.Fatal("first pair not swapped correctly")
	} else if second.First() != 1 || second.Second() != "one" {
		t.Fatal("second pair not swapped correctly")
	}
}

func TestSwapSelf(t *testing.T) {
	var pair *collections.Pair[int, string] = collections.NewPair(1, "one")

	pair.Swap(pair)

	if pair.First() != 1 {
		t.Fatal("unexpected first value")
	} else if pair.Second() != "one" {
		t.Fatal("unexpected second value")
	}
}

func TestFlip(t *testing.T) {
	var pair *collections.Pair[int, string] = collections.NewPair(42, "hello")

	var flipped *collections.Pair[string, int] = pair.Flip()

	if flipped.First() != "hello" {
		t.Fatalf("expected hello, got %s", flipped.First())
	} else if flipped.Second() != 42 {
		t.Fatalf("expected 42, got %d", flipped.Second())
	}
}

func TestFlipDoesNotModifyOriginal(t *testing.T) {
	var pair *collections.Pair[int, int] = collections.NewPair(1, 2)

	var flipped *collections.Pair[int, int] = pair.Flip()

	if pair.First() != 1 || pair.Second() != 2 {
		t.Fatal("original pair was modified")
	}

	flipped.Set(3, 4)

	if pair.First() != 1 || pair.Second() != 2 {
		t.Fatal("original pair changed after modifying flipped pair")
	}
}

func TestItems(t *testing.T) {
	var pair *collections.Pair[int, string] = collections.NewPair(5, "five")
	var first *int
	var second *string
	first, second = pair.Items()

	if *first != 5 {
		t.Fatalf("expected 5, got %d", *first)
	} else if *second != "five" {
		t.Fatalf("expected five, got %s", *second)
	}

	*first = 10
	*second = "ten"

	if pair.First() != 10 {
		t.Fatal("expected modification through Items()")
	} else if pair.Second() != "ten" {
		t.Fatal("expected modification through Items()")
	}
}

func TestPairDifferentTypes(t *testing.T) {
	type Person struct {
		Name string
	}

	var person Person = Person{Name: "Alice"}

	var pair *collections.Pair[Person, []int] = collections.NewPair(person, []int{1, 2, 3})

	if pair.First().Name != "Alice" {
		t.Fatal("unexpected first value")
	} else if len(pair.Second()) != 3 {
		t.Fatal("unexpected second value")
	}
}
func TestSwapMultipleTimes(t *testing.T) {
	var left *collections.Pair[int, int] = collections.NewPair(1, 2)
	var right *collections.Pair[int, int] = collections.NewPair(3, 4)

	left.Swap(right)
	left.Swap(right)

	if left.First() != 1 || left.Second() != 2 {
		t.Fatal("left pair incorrect")
	} else if right.First() != 3 || right.Second() != 4 {
		t.Fatal("right pair incorrect")
	}
}
