package collections

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