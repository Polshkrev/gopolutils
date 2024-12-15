package collections

import (
	"fmt"
	"os"
	"strings"

	"github.com/Polshkrev/gopolutils"
)

// Implementation of a set.
type Set[Type comparable] struct {
	items Mapping[Type, struct{}]
	size  uint64
}

// Construct a new set.
// Returns a pointer to a new empty set.
func NewSet[Type comparable]() *Set[Type] {
	var set *Set[Type] = new(Set[Type])
	set.items = NewMap[Type, struct{}]()
	set.size = 0
	return set
}

// Append an item to the set.
// If the set can not insert the item, this is a critical error that should not happen in most cicumstances, so — as a precaution — an error is printed to standard error and the programme exists.
func (set *Set[Type]) Append(item Type) {
	if set.Contains(item) {
		return
	}
	var except *gopolutils.Exception = set.items.Insert(item, struct{}{})
	if except != nil {
		fmt.Fprintln(os.Stderr, except.Error())
		os.Exit(1)
	}
	set.size++
}

// Append multiple items to the set.
// If the set can not insert the item, this is a critical error that should not happen in most cicumstances, so — as a precaution — an error is printed to standard error and the programme exists.
func (set *Set[Type]) Extend(items View[Type]) {
	var item Type
	for _, item = range items.Collect() {
		set.Append(item)
	}
}

// Access the data stored at a given index within the set.
// This method is not yet implemented.
// If the index is greater than the size of the set, an IndexOutOfRangeError is returned with a nil data pointer.
func (set Set[Type]) At(index uint64) (*Type, *gopolutils.Exception) {
	return nil, gopolutils.NewNamedException("NotImplementedError", "Can not access a set by index.")
}

// Remove an item in the set at a given index.
// If the set is evaluated to be empty, an IndexOutOfRangeError is returned.
// If the given index is greater than the size of the set, an IndexOutOfRangeError is returned.
// If no item can be found at the given index, an IndexError is returned.
// if an IndexError or an IndexOutOfRangError is returned, the set will not be modified.
func (set *Set[Type]) Remove(index uint64) *gopolutils.Exception {
	if set.IsEmpty() {
		return gopolutils.NewNamedException("IndexOutOfRangeError", fmt.Sprintf("Can not access an empty set at index %d.", index))
	} else if index > set.size {
		return gopolutils.NewNamedException("IndexOutOfRangeError", fmt.Sprintf("Can not access set of size %d at index %d.", set.size, index))
	}
	var i uint64
	var key Type
	for i, key = range Enumerate(set) {
		if i != index {
			continue
		}
		var except *gopolutils.Exception = set.items.Remove(key)
		if except != nil {
			return except
		}
		set.size--
		return nil
	}
	return gopolutils.NewNamedException("IndexError", fmt.Sprintf("No item at index %d exists.", index))
}

// Remove an item within the set without an exception.
// If the set is evaluated to be empty, the method will return without modifying the set.
// If the item is not in the set, the method will return without modifying the set.
func (set *Set[Type]) Discard(item Type) {
	if set.IsEmpty() {
		return
	} else if !set.Contains(item) {
		return
	}
	var except *gopolutils.Exception = set.items.Remove(item)
	if except != nil {
		// Critical error. This should be covered by the contains check, but just in case.
		fmt.Fprintln(os.Stderr, except.Error())
		os.Exit(1)
	}
	set.size--
}

// Access the size of the set.
// Returns the size of the set as an unsigned 64-bit integer.
func (set Set[_]) Size() uint64 {
	return set.size
}

// Determine if the given item is contained within the set.
// Returns true if the item is found within the set.
func (set Set[Type]) Contains(item Type) bool {
	return set.items.HasKey(item)
}

// Access the underlying data of the set.
// Returns a mutable pointer to a map representing the underlying data of the set.
func (set Set[Type]) Items() *[]Type {
	var keys []Type = set.items.Keys()
	return &keys
}

// Determine the difference between set and a given set operand.
// Returns a pointer to a new set that contains all the unique items that were contained within operand but not the original set.
func (set Set[Type]) Difference(other Set[Type]) *Set[Type] {
	var new *Set[Type] = NewSet[Type]()
	var item Type
	for _, item = range other.Collect() {
		if set.Contains(item) {
			continue
		}
		new.Append(item)
	}
	return new
}

// Determine the intersection between the set and a given set operand.
// Returns a pointer to a new set that contains all the unique items that were contained within both the original set and the operand.
func (set Set[Type]) Intersection(other Set[Type]) *Set[Type] {
	var new *Set[Type] = NewSet[Type]()
	var item Type
	for _, item = range other.Collect() {
		if !set.Contains(item) {
			continue
		}
		new.Append(item)
	}
	return new
}

// Determine if the set is empty.
// Returns true if the length of the underlying data stored in the set and the size of the set is equal to 0.
func (set Set[_]) IsEmpty() bool {
	return set.items.IsEmpty() && set.size == 0
}

// Access a slice of the data within the set.
// Returns a view of the data within the set.
func (set Set[Type]) Collect() []Type {
	return set.items.Keys()
}

// Transfer the data within the set to a linear array.
// Returns the set as an array.
func (set Set[Type]) ToArray() *Array[Type] {
	var array *Array[Type] = NewArray[Type]()
	array.Extend(set)
	return array
}

// Render a string representation of the set.
// Returns a string representation of the set.
func (set Set[Type]) ToString() string {
	var item Type
	var i uint64
	var buffer strings.Builder = strings.Builder{}
	buffer.WriteString("{")
	for i, item = range Enumerate(set) {
		if i == set.Size()-1 {
			buffer.WriteString(fmt.Sprintf("%v", item))
		} else {
			buffer.WriteString(fmt.Sprintf("%v,", item))
		}
	}
	buffer.WriteString("}")
	return buffer.String()
}
