package collections

import "github.com/Polshkrev/gopolutils"

// Interface to standardize a key-value pair mapping.
type Mapping[Key, Value any] interface {
	// Insert a key-value pair into the mapping.
	// If the key is already in the mapping, instead of just quietly not inserting into the mapping, a [gopolutils.KeyError] is returned.
	// If a [gopolutils.KeyError] is returned, the mapping is not modified.
	Insert(key Key, value Value) *gopolutils.Exception
	// Access an element at a given key within the mapping.
	// Returns a pointer to the data stored at the given key.
	// If the mapping is empty, a [gopolutils.ValueError] is returned with a nil data pointer.
	// If the key is not in the mapping, a [gopolutils.KeyError] is returned with a nil data pointer.
	At(key Key) (*Value, *gopolutils.Exception)
	// Update a value within the mapping.
	// If the mapping is empty, a [gopolutils.ValueError] is returned with a nil data pointer.
	// If the key does not exist in the mapping, a [gopolutils.KeyError] is returned.
	// If a [gopolutils.ValueError] or a [gopolutils.KeyError] is returned, the mapping is not modified.
	Update(key Key, value Value) *gopolutils.Exception
	// Access a slice of unique keys within the mapping.
	// Returns a slice of unique keys within the mapping.
	Keys() []Key
	// Access a slice of unique values within the mapping.
	// Returns a slice of unique values within the mapping.
	Values() []Value
	// Remove an item stored at a given key within the mapping.
	// If the map is empty, a [gopolutils.ValueError] is returned.
	// If the given key is not stored in the map, a [gopolutils.KeyError] is returned.
	// If a [gopolutils.ValueError] or a [gopolutils.KeyError] is returned, the mapping is not modified.
	Remove(key Key) *gopolutils.Exception
	// Determine if a given key is located within the mapping.
	// Returns true if the given key is found.
	HasKey(key Key) bool
	View[Pair[Key, Value]]
}
