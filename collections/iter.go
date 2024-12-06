package collections

// Implementation of a python-like enumeration function iterator.
// Returns an enumerated function iterator.
func Enumerate[Type any](items []Type) func(func(uint64, Type) bool) {
	return func(yield func(uint64, Type) bool) {
		var i uint64
		for i = 0; i <= uint64(len(items)-1); i++ {
			if !yield(i, items[i]) {
				return
			}
		}
	}
}

// Determine if a given item is within a given slice of items.
// Returns true if the given item is present within the given slice of items.
func In[Type comparable](items []Type, item Type) bool {
	var i Type
	for _, i = range items {
		if i != item {
			continue
		}
		return true
	}
	return false
}
