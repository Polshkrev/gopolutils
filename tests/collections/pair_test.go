package tests

import (
	"testing"

	"github.com/Polshkrev/gopolutils/collections"
)

func TestPairFirstConstructedCorrectly(test *testing.T) {
	var pair *collections.Pair[int, int] = collections.NewPair[int, int](1, 2)
	if *pair.First() != 1 {
		test.Errorf("Pair was constructed with an incorrect first argument: %d. Expected: %d\n", *pair.First(), 1)
	}
}

func TestPairSecondConstructedCorrectly(test *testing.T) {
	var pair *collections.Pair[int, int] = collections.NewPair[int, int](1, 2)
	if *pair.Second() != 2 {
		test.Errorf("Pair was constructed with an incorrect second argument: %d. Expected: %d\n", *pair.Second(), 2)
	}
}

func TestSwapPairHasSameInitialFirstValue(test *testing.T) {
	var pair *collections.Pair[int, int] = collections.NewPair[int, int](1, 2)
	var operand *collections.Pair[int, int] = collections.NewPair[int, int](3, 4)
	pair.Swap(operand)
	if *pair.First() != 3 {
		test.Errorf("Swap had changed the first value of pair incorrectly: %d. Expected: %d\n", *pair.First(), 3)
	}
}

func TestSwapPairHasSameInitialSecondValue(test *testing.T) {
	var pair *collections.Pair[int, int] = collections.NewPair[int, int](1, 2)
	var operand *collections.Pair[int, int] = collections.NewPair[int, int](3, 4)
	pair.Swap(operand)
	if *pair.Second() != 4 {
		test.Errorf("Swap had changed the second value of pair incorrectly: %d. Expected: %d\n", *pair.Second(), 4)
	}
}

func TestSwapOperandHasSameInitialFirstValue(test *testing.T) {
	var pair *collections.Pair[int, int] = collections.NewPair[int, int](1, 2)
	var operand *collections.Pair[int, int] = collections.NewPair[int, int](3, 4)
	pair.Swap(operand)
	if *operand.First() != 1 {
		test.Errorf("Swap had changed the first value of operand incorrectly: %d. Expected: %d\n", *operand.First(), 1)
	}
}

func TestSwapOperandHasSameInitialSecondValue(test *testing.T) {
	var pair *collections.Pair[int, int] = collections.NewPair[int, int](1, 2)
	var operand *collections.Pair[int, int] = collections.NewPair[int, int](3, 4)
	pair.Swap(operand)
	if *operand.Second() != 2 {
		test.Errorf("Swap had changed the first value of operand incorrectly: %d. Expected: %d\n", *operand.Second(), 2)
	}
}
