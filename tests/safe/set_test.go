package tests

import (
	"testing"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
	"github.com/Polshkrev/gopolutils/safe"
)

func TestSafeSetConstructorIsNotNil(test *testing.T) {
	var nilSet *safe.Set[int] = safe.NewSet[int]()
	if nilSet == nil {
		test.Errorf("Set constructor is nil.\n")
	}
}

func TestSafeSetAppend(test *testing.T) {
	var mock *safe.Set[int] = safe.NewSet[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	if !collections.In(mock, 2) {
		test.Errorf("Set did not append correctly.\n")
	}
}

func TestSafeSetDiscardSuccess(test *testing.T) {
	var mock *safe.Set[int] = safe.NewSet[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	mock.Discard(0)
	if collections.In(mock, 0) {
		test.Errorf("Set '%+v' did not discard at index '%d' correctly.\n", *mock, 0)
	}
}

func TestSafeSetDiscardFail(test *testing.T) {
	var mock *safe.Set[int] = safe.NewSet[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	mock.Discard(10)
	if !collections.In(mock, 0) {
		test.Errorf("Set '%+v' did not remove at index '%d' correctly.\n", *mock, 10)
	}
}

func TestSafeSetSizeSuccess(test *testing.T) {
	var mock *safe.Set[int] = safe.NewSet[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var size gopolutils.Size = mock.Size()
	if size != 3 {
		test.Errorf("Set size was not retuned correctly: '%d'.\n", size)
	}
}

func TestSafeSetSizeFail(test *testing.T) {
	var mock *safe.Set[int] = safe.NewSet[int]()
	var size gopolutils.Size = mock.Size()
	if size != 0 {
		test.Errorf("Set size was not retuned correctly: '%d'.\n", size)
	}
}

func TestSafeSetIsEmptySuccess(test *testing.T) {
	var mock *safe.Set[int] = safe.NewSet[int]()
	var result bool = mock.IsEmpty()
	if !result {
		test.Errorf("Set is empty was not retuned correctly: '%t'.\n", result)
	}
}

func TestSafeSetIsEmptyFail(test *testing.T) {
	var mock *safe.Set[int] = safe.NewSet[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var result bool = mock.IsEmpty()
	if result {
		test.Errorf("Set is empty was not retuned correctly: '%t'.\n", result)
	}
}

func TestSafeSetContainsSuccess(test *testing.T) {
	var mock *safe.Set[int] = safe.NewSet[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	var result bool = mock.Contains(0)
	if !result {
		test.Errorf("Set contains was not retuned correctly: '%t'.\n", result)
	}
}

func TestSafeSetContainsFail(test *testing.T) {
	var mock *safe.Set[int] = safe.NewSet[int]()
	var result bool = mock.Contains(0)
	if result {
		test.Errorf("Set contains was not retuned correctly: '%t'.\n", result)
	}
}

func TestSafeSetDifferenceSuccess(test *testing.T) {
	var mockOne *safe.Set[int] = safe.NewSet[int]()
	mockOne.Append(0)
	mockOne.Append(1)
	mockOne.Append(2)
	var mockTwo *safe.Set[int] = safe.NewSet[int]()
	mockTwo.Append(0)
	mockTwo.Append(2)
	var difference *safe.Set[int] = mockOne.Difference(*mockTwo)
	var result bool = difference.Contains(2)
	if result {
		test.Errorf("Set difference contains was not retuned correctly: '%t'.\n", result)
	}
}

func TestSafeSetDifferenceFail(test *testing.T) {
	var mockOne *safe.Set[int] = safe.NewSet[int]()
	mockOne.Append(0)
	mockOne.Append(1)
	mockOne.Append(2)
	var mockTwo *safe.Set[int] = safe.NewSet[int]()
	mockTwo.Append(0)
	mockTwo.Append(1)
	mockTwo.Append(2)
	var difference *safe.Set[int] = mockOne.Difference(*mockTwo)
	var result bool = difference.Contains(2)
	if result {
		test.Errorf("Set difference contains was not retuned correctly: '%t'.\n", result)
	}
}

func TestSafeSetIntersectionSuccess(test *testing.T) {
	var mockOne *safe.Set[int] = safe.NewSet[int]()
	mockOne.Append(0)
	mockOne.Append(1)
	mockOne.Append(2)
	var mockTwo *safe.Set[int] = safe.NewSet[int]()
	mockTwo.Append(0)
	var intersection *safe.Set[int] = mockOne.Intersection(*mockTwo)
	var result bool = intersection.Contains(0)
	if !result {
		test.Errorf("Set intersection contains was not retuned correctly: '%t'.\n", result)
	}
}

func TestSafeSetIntersectionFail(test *testing.T) {
	var mockOne *safe.Set[int] = safe.NewSet[int]()
	mockOne.Append(0)
	mockOne.Append(1)
	mockOne.Append(2)
	var mockTwo *safe.Set[int] = safe.NewSet[int]()
	mockTwo.Append(0)
	mockTwo.Append(1)
	var intersection *safe.Set[int] = mockOne.Intersection(*mockTwo)
	var result bool = intersection.Contains(2)
	if result {
		test.Errorf("Set intersection contains was not retuned correctly: '%t'.\n", result)
	}
}

func TestSafeSetItemsNotNil(test *testing.T) {
	var mock *safe.Set[int] = safe.NewSet[int]()
	mock.Append(0)
	mock.Append(1)
	mock.Append(2)
	if mock.Items() == nil {
		test.Errorf("Set items is nil.\n")
	}
}
