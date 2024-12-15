package collections

import (
	"fmt"

	"github.com/Polshkrev/gopolutils"
)

// Implementation of a HashMap.
type Map[Key comparable, Value any] struct {
	items map[Key]Value
	size  uint64
}

// Consruct a new map.
// Returns a pointer to a new empty map.
func NewMap[Key comparable, Value any]() *Map[Key, Value] {
	var mapping *Map[Key, Value] = new(Map[Key, Value])
	mapping.items = make(map[Key]Value, 0)
	mapping.size = 0
	return mapping
}

// Insert a key-value pair into the map.
// If the key is already in the map, instead of just quietly not inserting into the map, a KeyEror is retruned.
func (mapping *Map[Key, Value]) Insert(key Key, value Value) *gopolutils.Exception {
	if mapping.HasKey(key) {
		return gopolutils.NewNamedException("KeyError", fmt.Sprintf("Key '%v' already exists.", key))
	}
	mapping.items[key] = value
	mapping.size++
	return nil
}

// Access an element at a given key within the map.
// Returns a pointer to the data stored at the given key.
// If the key is not in the map, a KeyError is returned with a nil data pointer.
func (mapping Map[Key, Value]) At(key Key) (*Value, *gopolutils.Exception) {
	if mapping.IsEmpty() {
		return nil, gopolutils.NewNamedException("KeyError", fmt.Sprintf("Can not access an empty map at key '%+v'.", key))
	} else if !mapping.HasKey(key) {
		return nil, gopolutils.NewNamedException("KeyError", fmt.Sprintf("Key '%+v' does not exist.", key))
	}
	var value Value = mapping.items[key]
	return &value, nil
}

// Access a slice of unique keys within the map.
// Returns a slice of unique keys within the map.
func (mapping Map[Key, _]) Keys() []Key {
	var keys []Key = make([]Key, 0)
	var key Key
	for key = range mapping.items {
		keys = append(keys, key)
	}
	return keys
}

// Access a slice of unique values within the map.
// Returns a slice of unique values within the map.
func (mapping Map[_, Value]) Values() []Value {
	var values []Value = make([]Value, 0)
	var value Value
	for _, value = range mapping.items {
		values = append(values, value)
	}
	return values
}

// Remove an item stored at a given key within the map.
// If the given key is not stored in the map, a KeyError is returned.
func (mapping *Map[Key, _]) Remove(key Key) *gopolutils.Exception {
	if mapping.IsEmpty() {
		return gopolutils.NewNamedException("KeyError", fmt.Sprintf("Can not remove from an empty map at key '%+v'", key))
	} else if !mapping.HasKey(key) {
		return gopolutils.NewNamedException("KeyError", fmt.Sprintf("Key '%+v' does not exist.", key))
	}
	delete(mapping.items, key)
	mapping.size--
	return nil
}

// Determine if a given key is stored in the map.
// Returns true if the key is stored in the map.
func (mapping Map[Key, _]) HasKey(key Key) bool {
	var found bool
	_, found = mapping.items[key]
	return found
}

// Acces the size of the map.
// Returns the size of the map as an unsigned 64-bit integer.
func (mapping Map[_, _]) Size() uint64 {
	return mapping.size
}

// Determine if the map is empty.
// Returns true if the length of the underlying data and the size of the map is equal to 0.
func (mapping Map[_, _]) IsEmpty() bool {
	return len(mapping.items) == 0 && mapping.size == 0
}
