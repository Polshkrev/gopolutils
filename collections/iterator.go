package collections

import (
	"iter"
	"slices"

	"github.com/Polshkrev/gopolutils"
)

// Standardization of the iterator pattern.
type Iterator[Value any] struct {
	sequence iter.Seq[Value]
}

// Construct a new iterator from a given [View].
// Returns a new iterator from a given [View].
func IteratorFrom[Value any](source View[Value]) *Iterator[Value] {
	var iterator *Iterator[Value] = new(Iterator[Value])
	iterator.sequence = slices.Values(source.Collect())
	return iterator
}

// Map a given callback to each item in an iterator.
// Returns a new iterator with each of the items in the iterator transformed mapped to the callback.
func (iterator *Iterator[Value]) Map(callback func(Value) Value) *Iterator[Value] {
	var result *Iterator[Value] = new(Iterator[Value])
	result.sequence = func(yield func(Value) bool) {
		var value Value
		for value = range iterator.sequence {
			if !yield(callback(value)) {
				break
			}
		}
	}
	return result
}

// Filter an interator based on a given predicate.
// Returns a new iterator with each of the values of the iterator filtered by the given predicate.
func (iterator *Iterator[Value]) Filter(predicate func(Value) bool) *Iterator[Value] {
	var result *Iterator[Value] = new(Iterator[Value])
	result.sequence = func(yield func(Value) bool) {
		var value Value
		for value = range iterator.sequence {
			if !predicate(value) {
				continue
			} else if !yield(value) {
				return
			}
		}
	}
	return result
}

// Call a given callback on each of the values of the iterator.
func (iterator *Iterator[Value]) ForEach(callback func(Value)) {
	var value Value
	for value = range iterator.sequence {
		callback(value)
	}
}

// Collect the iterator into a slice.
// Returns a slice of the values of the iterator.
func (iterator *Iterator[Value]) Collect() []Value {
	return slices.Collect(iterator.sequence)
}

// Obtain the size of the iterator.
// Returns the size of the iterator.
func (iterator *Iterator[View]) Size() gopolutils.Size {
	var values []View = slices.Collect(iterator.sequence)
	var length int = len(values)
	return gopolutils.Size(length)
}

// Determine if the iterator is empty.
// Returns true if the size of the iterator is zero, else false.
func (iterator *Iterator[_]) IsEmpty() bool {
	return iterator.Size() == 0
}
