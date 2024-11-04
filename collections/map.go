package collections

import (
	"fmt"

	"github.com/Polshkrev/gopolutils"
)

type Map[Key comparable, Value any] struct {
	items map[Key]Value
	size  uint64
}

func NewMap[Key comparable, Value any]() *Map[Key, Value] {
	var mapping *Map[Key, Value] = new(Map[Key, Value])
	mapping.items = make(map[Key]Value, 0)
	mapping.size = 0
	return mapping
}

func (mapping *Map[Key, Value]) Insert(key Key, value Value) *gopolutils.Exception {
	var foundException *gopolutils.Exception = gopolutils.NewNamedException("KeyError", fmt.Sprintf("Key %v already exists.", key))
	var found bool
	_, found = mapping.items[key]
	if found {
		return foundException
	}
	mapping.items[key] = value
	mapping.size++
	return nil
}
