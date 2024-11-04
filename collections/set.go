package collections

type Set[Type comparable] struct {
	items map[Type]struct{}
	size  uint64
}

func NewSet[Type comparable]() *Set[Type] {
	var set *Set[Type] = new(Set[Type])
	set.items = make(map[Type]struct{}, 0)
	set.size = 0
	return set
}

func (set *Set[Type]) Append(item Type) {
	var found bool
	_, found = set.items[item]
	if found {
		return
	}
	set.items[item] = struct{}{}
	set.size++
}
