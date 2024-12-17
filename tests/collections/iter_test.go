package tests

import (
	"testing"

	"github.com/Polshkrev/gopolutils/collections"
)

var (
	arrayMock *collections.Array[int] = collections.NewArray[int]()
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

func setupArrayMock(testing.TB) func(testing.TB) {
	arrayMock.Append(0)
	arrayMock.Append(1)
	arrayMock.Append(2)
	return func(testing.TB) {}
}

func TestInSucces(test *testing.T) {
	var teardown func(testing.TB) = setupArrayMock(test)
	defer teardown(test)
	var item int = 0
	var result bool = collections.In(arrayMock, item)
	if !result {
		test.Errorf("Item '%d' is evaluated to not be in array '%+v'.", item, arrayMock)
	}
}

func TestInFail(test *testing.T) {
	var teardown func(testing.TB) = setupArrayMock(test)
	defer teardown(test)
	var item int = 5
	var result bool = collections.In(arrayMock, item)
	if result {
		test.Errorf("Item '%d' is evaluated to be in array '%+v'.", item, arrayMock)
	}
}

func TestEnumerationSucces(test *testing.T) {
	var enumTeardown func(testing.TB) = setupEnumMock(test)
	defer enumTeardown(test)
	var arrayTeardown func(testing.TB) = setupArrayMock(test)
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
