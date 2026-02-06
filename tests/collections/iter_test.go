package tests

import (
	"testing"

	"github.com/Polshkrev/gopolutils/collections"
)

var (
	arrayInMock      *collections.Array[int]              = collections.NewArray[int]()
	enumerationMocks *collections.Array[Enumeration[int]] = collections.NewArray[Enumeration[int]]()
)

type Enumeration[Type any] struct {
	Index uint64
	Item  Type
}

func NewEnumeration[Type any](index uint64, item Type) *Enumeration[Type] {
	var enumeration *Enumeration[Type] = new(Enumeration[Type])
	enumeration.Index = index
	enumeration.Item = item
	return enumeration
}

func setupEnumMock(testing.TB) func(testing.TB) {
	var index uint64
	for index = range 10 {
		enumerationMocks.Append(*NewEnumeration[int](index, int(index)))
	}
	return func(testing.TB) {}
}

func setupArrayInMock(testing.TB) func(testing.TB) {
	arrayInMock.Append(0)
	arrayInMock.Append(1)
	arrayInMock.Append(2)
	return func(testing.TB) {}
}

func TestInSucces(test *testing.T) {
	var teardown func(testing.TB) = setupArrayInMock(test)
	defer teardown(test)
	var item int = 0
	var result bool = collections.In(arrayInMock, item)
	if !result {
		test.Errorf("Item '%d' is evaluated to not be in array '%+v'.", item, arrayInMock)
	}
}

func TestInFail(test *testing.T) {
	var teardown func(testing.TB) = setupArrayInMock(test)
	defer teardown(test)
	var item int = 5
	var result bool = collections.In(arrayInMock, item)
	if result {
		test.Errorf("Item '%d' is evaluated to be in array '%+v'.", item, arrayInMock)
	}
}

func TestEnumerationSucces(test *testing.T) {
	var enumTeardown func(testing.TB) = setupEnumMock(test)
	defer enumTeardown(test)
	var arrayTeardown func(testing.TB) = setupArrayInMock(test)
	defer arrayTeardown(test)
	var index uint64
	var found bool = false
	var find Enumeration[int] = *NewEnumeration[int](5, 5)
	for index = range collections.Enumerate(*enumerationMocks) {
		if *NewEnumeration[int](index, int(index)) != find {
			continue
		}
		found = true
	}
	if !found {
		test.Errorf("Enumeration not found.")
	}
}
