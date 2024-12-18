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
	except = mapping.Insert(0, 0)
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