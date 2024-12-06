package collections

// Implementation of a python-like enumeration function iterator.
// Returns an enumerated function iterator.
func Enumerate[Type any](items View[Type]) func(func(uint64, Type) bool) {
	return func(yield func(uint64, Type) bool) {
		var i uint64
		for i = 0; i <= items.Size()-1; i++ {
			if !yield(i, items.Collect()[i]) {
				return
			}
		}
	}
}

// Determine if a given item is within a given slice of items.
// Returns true if the given item is present within the given slice of items.
func In[Type comparable](items View[Type], item Type) bool {
	var i Type
	for _, i = range items.Collect() {
		if i != item {
			continue
		}
		return true
	}
	return false
}
