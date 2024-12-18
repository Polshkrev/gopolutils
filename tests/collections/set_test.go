package tests

import (
	"testing"
	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
)

func TestSetConstructorIsNotNil(test *testing.T) {
	var nilSet *collections.Set[int] = collections.NewSet[int]()
	if nilSet == nil {
		test.Errorf("Set constructor is nil.\n")
	}
}

func TestSetAppend(test *testing.T) {
	var mock *collections.Set[int] = collections.NewSet[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	if !collections.In(mock, 2) {
		test.Errorf("Set did not append correctly.\n")
	}
}

func TestSetRemoveSuccess(test *testing.T) {
	var mock *collections.Set[int] = collections.NewSet[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var except *gopolutils.Exception = mock.Remove(0)
	if except != nil || collections.In(mock, 0) {
		test.Errorf("Set '%+v' did not remove at index '%d' correctly: %s.\n", *mock, 0, except.Error())
	}
}

func TestSetRemoveFail(test *testing.T) {
	var mock *collections.Set[int] = collections.NewSet[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var except *gopolutils.Exception = mock.Remove(10)
	if except == nil || !collections.In(mock, 0) {
		test.Errorf("Set '%+v' did not remove at index '%d' correctly.\n", *mock, 10)
	}
}

func TestSetDiscardSuccess(test *testing.T) {
	var mock *collections.Set[int] = collections.NewSet[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	mock.Discard(0)
	if collections.In(mock, 0) {
		test.Errorf("Set '%+v' did not discard at index '%d' correctly.\n", *mock, 0)
	}
}

func TestSetDiscardFail(test *testing.T) {
	var mock *collections.Set[int] = collections.NewSet[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	mock.Discard(10)
	if !collections.In(mock, 0) {
		test.Errorf("Set '%+v' did not remove at index '%d' correctly.\n", *mock, 10)
	}
}

func TestSetSizeSuccess(test *testing.T) {
	var mock *collections.Set[int] = collections.NewSet[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var size uint64 = mock.Size()
	if size != 3 {
		test.Errorf("Set size was not retuned correctly: '%d'.\n", size)
	}
}

func TestSetSizeFail(test *testing.T) {
	var mock *collections.Set[int] = collections.NewSet[int]()
	var size uint64 = mock.Size()
	if size != 0 {
		test.Errorf("Set size was not retuned correctly: '%d'.\n", size)
	}
}

func TestSetIsEmptySuccess(test *testing.T) {
	var mock *collections.Set[int] = collections.NewSet[int]()
	var result bool = mock.IsEmpty()
	if !result {
		test.Errorf("Set is empty was not retuned correctly: '%t'.\n", result)
	}
}

func TestSetIsEmptyFail(test *testing.T) {
	var mock *collections.Set[int] = collections.NewSet[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var result bool = mock.IsEmpty()
	if result {
		test.Errorf("Set is empty was not retuned correctly: '%t'.\n", result)
	}
}

func TestSetContainsSuccess(test *testing.T) {
	var mock *collections.Set[int] = collections.NewSet[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var result bool = mock.Contains(0)
	if !result {
		test.Errorf("Set contains was not retuned correctly: '%t'.\n", result)
	}
}

func TestSetContainsFail(test *testing.T) {
	var mock *collections.Set[int] = collections.NewSet[int]()
	var result bool = mock.Contains(0)
	if result {
		test.Errorf("Set contains was not retuned correctly: '%t'.\n", result)
	}
}

func TestSetDifferenceSuccess(test *testing.T) {
	var mockOne *collections.Set[int] = collections.NewSet[int]()
	mockOne.Append(0)
	mockOne.Append(1)
	mockOne.Append(2)
	var mockTwo *collections.Set[int] = collections.NewSet[int]()
	mockTwo.Append(0)
	mockTwo.Append(1)
	var difference *collections.Set[int] = mockOne.Difference(*mockTwo)
	var result bool = difference.Contains(2)
	if !result {
		test.Errorf("Set difference contains was not retuned correctly: '%t'.\n", result)
	}
}

func TestSetDifferenceFail(test *testing.T) {
	var mockOne *collections.Set[int] = collections.NewSet[int]()
	mockOne.Append(0)
	mockOne.Append(1)
	mockOne.Append(2)
	var mockTwo *collections.Set[int] = collections.NewSet[int]()
	mockTwo.Append(0)
	mockTwo.Append(1)
	mockTwo.Append(2)
	var difference *collections.Set[int] = mockOne.Difference(*mockTwo)
	var result bool = difference.Contains(2)
	if result {
		test.Errorf("Set difference contains was not retuned correctly: '%t'.\n", result)
	}
}

func TestSetIntersectionSuccess(test *testing.T) {
	var mockOne *collections.Set[int] = collections.NewSet[int]()
	mockOne.Append(0)
	mockOne.Append(1)
	mockOne.Append(2)
	var mockTwo *collections.Set[int] = collections.NewSet[int]()
	mockTwo.Append(0)
	var intersection *collections.Set[int] = mockOne.Intersection(*mockTwo)
	var result bool = intersection.Contains(0)
	if !result {
		test.Errorf("Set intersection contains was not retuned correctly: '%t'.\n", result)
	}
}

func TestSetIntersectionFail(test *testing.T) {
	var mockOne *collections.Set[int] = collections.NewSet[int]()
	mockOne.Append(0)
	mockOne.Append(1)
	mockOne.Append(2)
	var mockTwo *collections.Set[int] = collections.NewSet[int]()
	mockTwo.Append(0)
	mockTwo.Append(1)
	var intersection *collections.Set[int] = mockOne.Intersection(*mockTwo)
	var result bool = intersection.Contains(2)
	if result {
		test.Errorf("Set intersection contains was not retuned correctly: '%t'.\n", result)
	}
}

func TestSetItemsNotNil(test *testing.T) {
	var mock *collections.Set[int] = collections.NewSet[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	if mock.Items() == nil {
		test.Errorf("Set items is nil.\n")
	}
}