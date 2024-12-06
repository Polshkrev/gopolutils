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
// If the given index is greater than the size of the array, an IndexOutOfRangeError is returned with a nil data pointer.
func (array Array[Type]) At(index uint64) (*Type, *gopolutils.Exception) {
	var outOfRange *gopolutils.Exception = gopolutils.NewNamedException("IndexOutOfRangeError", fmt.Sprintf("Can not access array of size %d at index %d.", array.size, index))
	if index > array.size {
		return nil, outOfRange
	}
	return &array.items[index], nil
}

// Remove the data stored in the array at a given index.
// If the given index is greater than the size of the array, an IndexOutOfRangeError is returned.
// If an IndexOutOfRangeError is returned, the array is not modified.
func (array *Array[Type]) Remove(index uint64) *gopolutils.Exception {
	var outOfRange *gopolutils.Exception = gopolutils.NewNamedException("IndexOutOfRangeError", fmt.Sprintf("Can not access array of size %d at index %d.", array.size, index))
	if index > array.size {
		return outOfRange
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
