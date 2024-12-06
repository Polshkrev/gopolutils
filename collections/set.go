package collections

import (
	"fmt"
	"strings"

	"github.com/Polshkrev/gopolutils"
)

// Implementation of a set.
type Set[Type comparable] struct {
	items map[Type]struct{}
	size  uint64
}

// Construct a new set.
// Returns a pointer to a new empty set.
func NewSet[Type comparable]() *Set[Type] {
	var set *Set[Type] = new(Set[Type])
	set.items = make(map[Type]struct{}, 0)
	set.size = 0
	return set
}

// Append an item to the set.
func (set *Set[Type]) Append(item Type) {
	if set.Contains(item) {
		return
	}
	set.items[item] = struct{}{}
	set.size++
}

// Append multiple items to the set.
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

// Remove an item in the set.
// if the item is not in the set, a KeyError is returned.
// if a KeyError is returned, the set will not be modified.
func (set *Set[Type]) Remove(item Type) *gopolutils.Exception {
	var notFound *gopolutils.Exception = gopolutils.NewNamedException("KeyError", fmt.Sprintf("Item '%v' does not exist.", item))
	if !set.Contains(item) {
		return notFound
	}
	delete(set.items, item)
	set.size--
	return nil
}

// Remove an item within the set without an exception.
// If the item is not in the set, the method will return without modifying the set.
func (set *Set[Type]) Discard(item Type) {
	if !set.Contains(item) {
		return
	}
	delete(set.items, item)
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
	var found bool
	_, found = set.items[item]
	return found
}

// Access the underlying data of the set.
// Returns a mutable pointer to a map representing the underlying data of the set.
func (set Set[Type]) Items() *map[Type]struct{} {
	return &set.items
}

// Determine the difference between set and a given set operand.
// Returns a pointer to a new set that contains all the unique items that were contained within operand but not the original set.
func (set Set[Type]) Difference(other Set[Type]) *Set[Type] {
	var new *Set[Type] = NewSet[Type]()
	var item Type
	for item = range *other.Items() {
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
	for item = range *other.Items() {
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
	return len(set.items) == 0 && set.size == 0
}

// Access a slice of the data within the set.
// Returns a view of the data within the set.
func (set Set[Type]) Collect() []Type {
	var list []Type = make([]Type, 0)
	var item Type
	for item = range *set.Items() {
		list = append(list, item)
	}
	return list
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
	var i int
	var buffer strings.Builder = strings.Builder{}
	buffer.WriteString("{")
	for i, item = range set.Collect() {
		if i == int(set.Size()-1) {
			buffer.WriteString(fmt.Sprintf("%v", item))
		} else {
			buffer.WriteString(fmt.Sprintf("%v,", item))
		}
	}
	buffer.WriteString("}")
	return buffer.String()
}
