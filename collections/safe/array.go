package safe

import (
	"sync"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
)

// Implementation of a concurrent-safe array.
type Array[Type any] struct {
	lock  sync.RWMutex
	items []Type
	size  gopolutils.Size
}

// Construct a new array.
// Returns a pointer to a new empty array.
func NewArray[Type any]() *Array[Type] {
	var array *Array[Type] = new(Array[Type])
	array.items = make([]Type, 0)
	array.size = 0
	return array
}

// Append an item to the array.
func (array *Array[Type]) Append(item Type) {
	array.Lock()
	defer array.Unlock()
	array.items = append(array.items, item)
	array.size++
}

// Append multiple items to the array.
func (array *Array[Type]) Extend(items collections.View[Type]) {
	var i int
	for i = range items.Collect() {
		var item Type = items.Collect()[i]
		array.Append(item)
	}
}

// Access the data stored in the array at a given index.
// If the array is empty, a [gopolutils.ValueError] is returned with a nil data pointer.
// If the given index is greater than the size of the array, an [gopolutils.OutOfRangeError] is returned with a nil data pointer.
func (array *Array[Type]) At(index gopolutils.Size) (*Type, *gopolutils.Exception) {
	array.RLock()
	defer array.RUnlock()
	if array.IsEmpty() {
		return nil, gopolutils.NewNamedException(gopolutils.ValueError, "Can not access an empty array at index %d.", index)
	} else if index > array.size {
		return nil, gopolutils.NewNamedException(gopolutils.OutOfRangeError, "Can not access array of size %d at index %d.", array.size, index)
	}
	return &array.items[index], nil
}

// Update a value within the collection.
// If the collection is empty, a [gopolutils.ValueError] is returned.
// If the given index is greater than the collection size, an [gopolutils.OutOfRangeError] is returned.
// If a [gopolutils.ValueError] or an [gopolutils.OutOfRangeError] is returned, the collection is not modified.
func (array *Array[Type]) Update(index gopolutils.Size, value Type) *gopolutils.Exception {
	array.Lock()
	defer array.Unlock()
	if array.IsEmpty() {
		return gopolutils.NewNamedException(gopolutils.ValueError, "Can not access an empty array at index %d.", index)
	} else if index > array.size {
		return gopolutils.NewNamedException(gopolutils.OutOfRangeError, "Can not access array of size %d at index %d.", array.size, index)
	}
	array.items[index] = value
	return nil
}

// Remove the data stored in the array at a given index.
// If the array is empty, a [gopolutils.ValueError] is returned.
// If the given index is greater than the size of the array, an [gopolutils.OutOfRangeError] is returned.
// If a [gopolutils.ValueError] or an [gopolutils.OutOfRangeError] is returned, the array is not modified.
func (array *Array[Type]) Remove(index gopolutils.Size) *gopolutils.Exception {
	array.Lock()
	defer array.Unlock()
	if array.IsEmpty() {
		return gopolutils.NewNamedException(gopolutils.ValueError, "Can not remove from an empty array at index %d.", index)
	} else if index > array.size {
		return gopolutils.NewNamedException(gopolutils.OutOfRangeError, "Can not remove element of array of size %d at index %d.", array.size, index)
	}
	array.items = append(array.items[:index], array.items[index+1:]...)
	array.size--
	return nil
}

// Access the underlying data of the array.
// Returns a mutable pointer to the data stored in the array.
func (array *Array[Type]) Items() *[]Type {
	array.RLock()
	defer array.RUnlock()
	return &array.items
}

// Collect the data stored in the array as a slice.
// Returns a view into the data stored in the array.
func (array *Array[Type]) Collect() []Type {
	array.RLock()
	defer array.RUnlock()
	return array.items
}

// Access the size of the array.
// Returns the size of the array as an unsigned 64-bit integer.
func (array *Array[_]) Size() gopolutils.Size {
	array.RLock()
	defer array.RUnlock()
	return array.size
}

// Determine if the array is empty.
// Returns true if the length of the data and the size of the array are equal to 0.
func (array *Array[_]) IsEmpty() bool {
	array.RLock()
	defer array.RUnlock()
	return len(array.items) == 0 && array.size == 0
}

// Lock the internal mutex of the collection for both reading and writing.
func (array *Array[_]) Lock() {
	array.lock.Lock()
}

// Unlock the internal mutex of the collection for both reading and writing.
func (array *Array[_]) Unlock() {
	array.lock.Unlock()
}

// Lock the internal mutex of the collection for reading.
func (array *Array[_]) RLock() {
	array.lock.RLock()
}

// Unock the internal mutex of the collection for reading.
func (array *Array[_]) RUnlock() {
	array.lock.RUnlock()
}
