package tests

import (
	"testing"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections/safe"
)

func TestSafeMapConstructorIsNotNil(test *testing.T) {
	var nilMap *safe.Map[int, int] = safe.NewMap[int, int]()
	if nilMap == nil {
		test.Errorf("Map constructor returned nil.")
	}
}

func TestSafeMapInsertSuccess(test *testing.T) {
	var mapping *safe.Map[int, int] = safe.NewMap[int, int]()
	var except *gopolutils.Exception = mapping.Insert(0, 0)
	if except != nil {
		test.Errorf("Map insert did not return nil: %+v %s.", *mapping, except.Error())
	}
}

func TestSafeMapInsertFail(test *testing.T) {
	var mapping *safe.Map[int, int] = safe.NewMap[int, int]()
	var except *gopolutils.Exception = mapping.Insert(0, 0)
	except = mapping.Insert(0, 1)
	if except == nil {
		test.Errorf("Map insert returned nil: %+v.", *mapping)
	}
}

func TestSafeMapAtSuccess(test *testing.T) {
	var mapping *safe.Map[int, int] = safe.NewMap[int, int]()
	mapping.Insert(0, 0)
	mapping.Insert(1, 1)
	mapping.Insert(2, 2)
	var item *int
	var except *gopolutils.Exception
	item, except = mapping.At(1)
	if except != nil || *item != 1 {
		test.Errorf("Map at did not return nil: %+v %s.", *mapping, except.Error())
	}
}

func TestSafeMapAtFail(test *testing.T) {
	var mapping *safe.Map[int, int] = safe.NewMap[int, int]()
	mapping.Insert(0, 0)
	mapping.Insert(1, 1)
	mapping.Insert(2, 2)
	var except *gopolutils.Exception
	_, except = mapping.At(10)
	if except == nil {
		test.Errorf("Map at returned nil: %+v.", *mapping)
	}
}

func TestSafeMapUpdateSuccess(test *testing.T) {
	var mapping *safe.Map[int, int] = safe.NewMap[int, int]()
	mapping.Insert(0, 0)
	mapping.Insert(1, 1)
	mapping.Insert(2, 2)
	var except *gopolutils.Exception = mapping.Update(1, 2)
	var result *int
	var exceptAt *gopolutils.Exception
	result, exceptAt = mapping.At(1)
	if except != nil || exceptAt != nil || *result != 2 {
		test.Errorf("Map update did not return nil: %+v %s, %s.", *mapping, except.Error(), exceptAt.Error())
	}
}

func TestSafeMapUpdateFail(test *testing.T) {
	var mapping *safe.Map[int, int] = safe.NewMap[int, int]()
	var except *gopolutils.Exception = mapping.Update(0, 2)
	var exceptAt *gopolutils.Exception
	_, exceptAt = mapping.At(0)
	if except == nil || exceptAt == nil {
		test.Errorf("Map update did not return nil: %+v.", *mapping)
	}
}

func TestSafeMapRemoveSuccess(test *testing.T) {
	var mapping *safe.Map[int, int] = safe.NewMap[int, int]()
	mapping.Insert(0, 0)
	mapping.Insert(1, 1)
	mapping.Insert(2, 2)
	var except *gopolutils.Exception = mapping.Remove(1)
	if except != nil {
		test.Errorf("Map remove did not return nil: %+v %s.", *mapping, except.Error())
	}
}

func TestSafeMapRemoveFail(test *testing.T) {
	var mapping *safe.Map[int, int] = safe.NewMap[int, int]()
	mapping.Insert(0, 0)
	mapping.Insert(1, 1)
	mapping.Insert(2, 2)
	var except *gopolutils.Exception = mapping.Remove(10)
	if except == nil {
		test.Errorf("Map remove returned nil: %+v.", *mapping)
	}
}

func TestSafeMapHasKeySuccess(test *testing.T) {
	var mapping *safe.Map[int, int] = safe.NewMap[int, int]()
	mapping.Insert(0, 0)
	mapping.Insert(1, 1)
	mapping.Insert(2, 2)
	var result bool = mapping.HasKey(1)
	if !result {
		test.Errorf("Map has key did not return correctly: %+v.", *mapping)
	}
}

func TestSafeMapHasKeyFail(test *testing.T) {
	var mapping *safe.Map[int, int] = safe.NewMap[int, int]()
	mapping.Insert(0, 0)
	mapping.Insert(1, 1)
	mapping.Insert(2, 2)
	var result bool = mapping.HasKey(10)
	if result {
		test.Errorf("Map has key did not return correctly: %+v.", *mapping)
	}
}

func TestSafeMapSizeSuccess(test *testing.T) {
	var mapping *safe.Map[int, int] = safe.NewMap[int, int]()
	mapping.Insert(0, 0)
	mapping.Insert(1, 1)
	mapping.Insert(2, 2)
	var size gopolutils.Size = mapping.Size()
	if size != 3 {
		test.Errorf("Array size was not retuned correctly: '%d'.\n", size)
	}
}

func TestSafeMapSizeFail(test *testing.T) {
	var mapping *safe.Map[int, int] = safe.NewMap[int, int]()
	mapping.Insert(0, 0)
	mapping.Insert(1, 1)
	mapping.Insert(2, 2)
	mapping.Insert(4, 4)
	var size gopolutils.Size = mapping.Size()
	if size == 3 {
		test.Errorf("Array size was not retuned correctly: '%d'.\n", size)
	}
}

func TestSafeMapIsEmptySuccess(test *testing.T) {
	var mapping *safe.Map[int, int] = safe.NewMap[int, int]()
	mapping.Insert(0, 0)
	mapping.Insert(1, 1)
	mapping.Insert(2, 2)
	var result bool = mapping.IsEmpty()
	if result {
		test.Errorf("Array is empty was not retuned correctly: '%t'.\n", result)
	}
}

func TestSafeMapIsEmptyFail(test *testing.T) {
	var mapping *safe.Map[int, int] = safe.NewMap[int, int]()
	var result bool = mapping.IsEmpty()
	if !result {
		test.Errorf("Array is empty was not retuned correctly: '%t'.\n", result)
	}
}
