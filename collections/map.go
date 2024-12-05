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
	var foundException *gopolutils.Exception = gopolutils.NewNamedException("KeyError", fmt.Sprintf("Key '%v' already exists.", key))
	if mapping.HasKey(key) {
		return foundException
	}
	mapping.items[key] = value
	mapping.size++
	return nil
}

func (mapping Map[Key, Value]) At(key Key) (*Value, *gopolutils.Exception) {
	var notFound *gopolutils.Exception = gopolutils.NewNamedException("KeyError", fmt.Sprintf("Key '%v' does not exist.", key))
	if !mapping.HasKey(key) {
		return nil, notFound
	}
	var value Value = mapping.items[key]
	return &value, nil
}

func (mapping Map[Key, _]) Keys() []Key {
	var keys []Key = make([]Key, 0)
	var key Key
	for key = range mapping.items {
		keys = append(keys, key)
	}
	return keys
}

func (mapping Map[_, Value]) Values() []Value {
	var values []Value = make([]Value, 0)
	var value Value
	for _, value = range mapping.items {
		values = append(values, value)
	}
	return values
}

func (mapping *Map[Key, _]) Remove(key Key) *gopolutils.Exception {
	var notFound *gopolutils.Exception = gopolutils.NewNamedException("KeyError", fmt.Sprintf("Key '%v' does not exist.", key))
	if !mapping.HasKey(key) {
		return notFound
	}
	delete(mapping.items, key)
	mapping.size--
	return nil
}

func (mapping Map[Key, _]) HasKey(key Key) bool {
	var found bool
	_, found = mapping.items[key]
	return found
}

func (mapping Map[_, _]) Size() uint64 {
	return mapping.size
}

// Determine if the map is empty.
// Returns true if the length of the underlying data and the size of the map is equal to 0.
func (mapping Map[_, _]) IsEmpty() bool {
	return len(mapping.items) == 0 && mapping.size == 0
}
