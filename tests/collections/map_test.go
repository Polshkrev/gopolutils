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

func TestMapInsertSuccess(test *testing.T) {
	var mapping *collections.Map[int, int] = collections.NewMap[int, int]()
	var except *gopolutils.Exception = mapping.Insert(0, 0)
	except = mapping.Insert(0, 0)
	if except == nil {
		test.Errorf("Map insert returned nil: %+v %s.", *mapping, except.Error())
	}
}