package tests

import (
	"testing"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
)

func TestMapConstructorIsNotNil(test *testing.T) {
	var nilMap *collections.Map[int, int] = collections.NewMap[int, int]()
	if nilMap == nil {
		test.Errorf("Map constructor returned nil.")
	}
}

func TestMapInsertSuccess(test *testing.T) {
	var mapping *collections.Map[int, int] = collections.NewMap[int, int]()
	var except *gopolutils.Exception = mapping.Insert(0, 0)
	if except != nil {
		test.Errorf("Map insert did not return nil: %+v %s.", *mapping, except.Error())
	}
}

func TestMapInsertFail(test *testing.T) {
	var mapping *collections.Map[int, int] = collections.NewMap[int, int]()
	var except *gopolutils.Exception = mapping.Insert(0, 0)
	except = mapping.Insert(0, 1)
	if except == nil {
		test.Errorf("Map insert returned nil: %+v.", *mapping)
	}
}

func TestMapAtSuccess(test *testing.T) {
	var mapping *collections.Map[int, int] = collections.NewMap[int, int]()
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

func TestMapAtFail(test *testing.T) {
	var mapping *collections.Map[int, int] = collections.NewMap[int, int]()
	mapping.Insert(0, 0)
	mapping.Insert(1, 1)
	mapping.Insert(2, 2)
	var except *gopolutils.Exception
	_, except = mapping.At(10)
	if except == nil {
		test.Errorf("Map at returned nil: %+v.", *mapping)
	}
}

func TestMapUpdateSuccess(test *testing.T) {
	var mapping *collections.Map[int, int] = collections.NewMap[int, int]()
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

func TestMapUpdateFail(test *testing.T) {
	var mapping *collections.Map[int, int] = collections.NewMap[int, int]()
	var except *gopolutils.Exception = mapping.Update(0, 2)
	var exceptAt *gopolutils.Exception
	_, exceptAt = mapping.At(0)
	if except == nil || exceptAt == nil {
		test.Errorf("Map update did not return nil: %+v.", *mapping)
	}
}

func TestMapRemoveSuccess(test *testing.T) {
	var mapping *collections.Map[int, int] = collections.NewMap[int, int]()
	mapping.Insert(0, 0)
	mapping.Insert(1, 1)
	mapping.Insert(2, 2)
	var except *gopolutils.Exception = mapping.Remove(1)
	if except != nil {
		test.Errorf("Map remove did not return nil: %+v %s.", *mapping, except.Error())
	}
}

func TestMapRemoveFail(test *testing.T) {
	var mapping *collections.Map[int, int] = collections.NewMap[int, int]()
	mapping.Insert(0, 0)
	mapping.Insert(1, 1)
	mapping.Insert(2, 2)
	var except *gopolutils.Exception = mapping.Remove(10)
	if except == nil {
		test.Errorf("Map remove returned nil: %+v.", *mapping)
	}
}

func TestMapHasKeySuccess(test *testing.T) {
	var mapping *collections.Map[int, int] = collections.NewMap[int, int]()
	mapping.Insert(0, 0)
	mapping.Insert(1, 1)
	mapping.Insert(2, 2)
	var result bool = mapping.HasKey(1)
	if !result {
		test.Errorf("Map has key did not return correctly: %+v.", *mapping)
	}
}

func TestMapHasKeyFail(test *testing.T) {
	var mapping *collections.Map[int, int] = collections.NewMap[int, int]()
	mapping.Insert(0, 0)
	mapping.Insert(1, 1)
	mapping.Insert(2, 2)
	var result bool = mapping.HasKey(10)
	if result {
		test.Errorf("Map has key did not return correctly: %+v.", *mapping)
	}
}

func TestMapSizeSuccess(test *testing.T) {
	var mapping *collections.Map[int, int] = collections.NewMap[int, int]()
	mapping.Insert(0, 0)
	mapping.Insert(1, 1)
	mapping.Insert(2, 2)
	var size uint64 = mapping.Size()
	if size != 3 {
		test.Errorf("Array size was not retuned correctly: '%d'.\n", size)
	}
}

func TestMapSizeFail(test *testing.T) {
	var mapping *collections.Map[int, int] = collections.NewMap[int, int]()
	mapping.Insert(0, 0)
	mapping.Insert(1, 1)
	mapping.Insert(2, 2)
	mapping.Insert(4, 4)
	var size uint64 = mapping.Size()
	if size == 3 {
		test.Errorf("Array size was not retuned correctly: '%d'.\n", size)
	}
}

func TestMapIsEmptySuccess(test *testing.T) {
	var mapping *collections.Map[int, int] = collections.NewMap[int, int]()
	mapping.Insert(0, 0)
	mapping.Insert(1, 1)
	mapping.Insert(2, 2)
	var result bool = mapping.IsEmpty()
	if result {
		test.Errorf("Array is empty was not retuned correctly: '%t'.\n", result)
	}
}

func TestMapIsEmptyFail(test *testing.T) {
	var mapping *collections.Map[int, int] = collections.NewMap[int, int]()
	var result bool = mapping.IsEmpty()
	if !result {
		test.Errorf("Array is empty was not retuned correctly: '%t'.\n", result)
	}
}
