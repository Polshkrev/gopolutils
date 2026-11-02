package collections

import (
	"slices"

	"github.com/Polshkrev/gopolutils"
)

// Implementation of a python-like enumeration function iterator.
// Returns an enumerated function iterator.
func Enumerate[Type any](items View[Type]) func(func(gopolutils.Size, Type) bool) {
	return func(yield func(gopolutils.Size, Type) bool) {
		var i gopolutils.Size
		for i = 0; i <= items.Size()-1; i++ {
			if !yield(i, items.Collect()[i]) {
				return
			}
		}
	}
}

// Implementation of a reverse function iterator.
// Returns a function iterator of a reversed view.
func Reverse[Type any](items View[Type]) func(func(gopolutils.Size, Type) bool) {
	return func(yield func(gopolutils.Size, Type) bool) {
		var i gopolutils.Size
		for i = items.Size(); i > 0; i-- {
			if !yield(i-1, items.Collect()[i-1]) {
				return
			}
		}
	}
}

// Determine if a given item is within a given slice of items.
// Returns true if the given item is present within the given slice of items.
func In[Type comparable](items View[Type], item Type) bool {
	return slices.Contains(items.Collect(), item)
}

// Chain a variadic list of views together.
// Returns a new collection of the combination of each of the passed in views.
func Chain[Type any](views ...View[Type]) Collection[Type] {
	var result Collection[Type] = NewArray[Type]()
	var i int
	for i = range views {
		var view View[Type] = views[i]
		result.Extend(view)
	}
	return result
}

// Sum given variadic [gopolutils.Number] values.
// Returns a [gopolutils.Number] of the sum of the given values.
func Sum[Type gopolutils.Number](values ...Type) Type {
	var result Type = 0
	var i int
	for i = range values {
		result += values[i]
	}
	return result
}
