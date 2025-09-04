package collections

import (
	"fmt"

	"github.com/Polshkrev/gopolutils"
)

// Implementation of a classical dynamic array.
type Array[Type any] struct {
	items []Type
	size  uint64
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
	array.items = append(array.items, item)
	array.size++
}

// Append multiple items to the array.
func (array *Array[Type]) Extend(items View[Type]) {
	for _, item := range items.Collect() {
		array.Append(item)
	}
}

// Access the data stored in the array at a given index.
// If the array is empty, an IndexOutOfRangeError is returned with a nil data pointer.
// If the given index is greater than the size of the array, an IndexOutOfRangeError is returned with a nil data pointer.
func (array Array[Type]) At(index uint64) (*Type, *gopolutils.Exception) {
	if array.IsEmpty() {
		return nil, gopolutils.NewNamedException(gopolutils.OutOfRangeError, fmt.Sprintf("Can not access an empty array at index %d.", index))
	} else if index > array.size {
		return nil, gopolutils.NewNamedException(gopolutils.OutOfRangeError, fmt.Sprintf("Can not access array of size %d at index %d.", array.size, index))
	}
	return &array.items[index], nil
}

// Update a value within the collection.
// If the given index is greater than the collection size, an IndexOutOfRangeError is returned.
// If the collection is empty, an IndexOutOfRangeError is returned.
// If an IndexOutOfRangeError is returned, the collection is not modified.
func (array *Array[Type]) Update(index uint64, value Type) *gopolutils.Exception {
	if array.IsEmpty() {
		return gopolutils.NewNamedException(gopolutils.OutOfRangeError, fmt.Sprintf("Can not access an empty array at index %d.", index))
	} else if index > array.size {
		return gopolutils.NewNamedException(gopolutils.OutOfRangeError, fmt.Sprintf("Can not access array of size %d at index %d.", array.size, index))
	}
	array.items[index] = value
	return nil
}

// Remove the data stored in the array at a given index.
// If the array is empty, an IndexOutOfRangeError is returned.
// If the given index is greater than the size of the array, an IndexOutOfRangeError is returned.
// If an IndexOutOfRangeError is returned, the array is not modified.
func (array *Array[Type]) Remove(index uint64) *gopolutils.Exception {
	if array.IsEmpty() {
		return gopolutils.NewNamedException(gopolutils.OutOfRangeError, fmt.Sprintf("Can not remove from an empty array at index %d.", index))
	} else if index > array.size {
		return gopolutils.NewNamedException(gopolutils.OutOfRangeError, fmt.Sprintf("Can not remove element of array of size %d at index %d.", array.size, index))
	}
	array.items = append(array.items[:index], array.items[index+1:]...)
	array.size--
	return nil
}

// Access the underlying data of the array.
// Returns a mutable pointer to the data stored in the array.
func (array Array[Type]) Items() *[]Type {
	return &array.items
}

// Collect the data stored in the array as a slice.
// Returns a view into the data stored in the array.
func (array Array[Type]) Collect() []Type {
	return array.items
}

// Access the size of the array.
// Returns the size of the array as an unsigned 64-bit integer.
func (array Array[_]) Size() uint64 {
	return array.size
}

// Determine if the array is empty.
// Returns true if the length of the data and the size of the array are equal to 0.
func (array Array[_]) IsEmpty() bool {
	return len(array.items) == 0 && array.size == 0
}
