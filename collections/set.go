package collections

import (
	"fmt"

	"github.com/Polshkrev/gopolutils"
)

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

func (set *Set[Type]) Remove(item Type) *gopolutils.Exception {
	var notFound *gopolutils.Exception = gopolutils.NewNamedException("KeyError", fmt.Sprintf("Item '%v' does not exist.", item))
	var found bool
	_, found = set.items[item]
	if !found {
		return notFound
	}
	delete(set.items, item)
	set.size--
	return nil
}

func (set *Set[Type]) Discard(item Type) {
	var found bool
	_, found = set.items[item]
	if !found {
		return
	}
	delete(set.items, item)
	set.size--
}
