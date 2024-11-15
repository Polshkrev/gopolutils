package collections

import (
	"fmt"
	"strings"

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
	if set.Contains(item) {
		return
	}
	set.items[item] = struct{}{}
	set.size++
}

func (set *Set[Type]) Extend(items []Type) {
	var item Type
	for _, item = range items {
		set.Append(item)
	}
}

func (set *Set[Type]) Remove(item Type) *gopolutils.Exception {
	var notFound *gopolutils.Exception = gopolutils.NewNamedException("KeyError", fmt.Sprintf("Item '%v' does not exist.", item))
	if !set.Contains(item) {
		return notFound
	}
	delete(set.items, item)
	set.size--
	return nil
}

func (set *Set[Type]) Discard(item Type) {
	if !set.Contains(item) {
		return
	}
	delete(set.items, item)
	set.size--
}

func (set Set[_]) Size() uint64 {
	return set.size
}

func (set Set[Type]) Contains(item Type) bool {
	var found bool
	_, found = set.items[item]
	return found
}

func (set Set[Type]) Items() *map[Type]struct{} {
	return &set.items
}

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

func (set Set[Type]) ToSlice() []Type {
	var list []Type = make([]Type, 0)
	var item Type
	for item = range *set.Items() {
		list = append(list, item)
	}
	return list
}

func (set Set[Type]) ToArray() *Array[Type] {
	var array *Array[Type] = NewArray[Type]()
	var list []Type = set.ToSlice()
	array.Extend(list)
	return array
}

func (set Set[Type]) ToString() string {
	var list []Type = set.ToSlice()
	var item Type
	var i int
	var buffer strings.Builder = strings.Builder{}
	buffer.WriteString("{")
	for i, item = range list {
		if i == len(list)-1 {
			buffer.WriteString(fmt.Sprintf("%v", item))
		} else {
			buffer.WriteString(fmt.Sprintf("%v,", item))
		}
	}
	buffer.WriteString("}")
	return buffer.String()
}
