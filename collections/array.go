package collections

import (
	"fmt"

	"github.com/Polshkrev/gopolutils"
)

type Array[Type any] struct {
	items []Type
	size  uint64
}

func NewArray[Type any]() *Array[Type] {
	var array *Array[Type] = new(Array[Type])
	array.items = make([]Type, 0)
	array.size = 0
	return array
}

func (array *Array[Type]) Append(item Type) {
	array.items = append(array.items, item)
	array.size++
}

func (array *Array[Type]) Extend(items []Type) {
	for _, item := range items {
		array.Append(item)
	}
}

func (array Array[Type]) At(index uint64) (*Type, *gopolutils.Exception) {
	var outOfRange *gopolutils.Exception = gopolutils.NewNamedException("IndexOutOfRangeError", fmt.Sprintf("Can not access array of size %d at index %d.", array.size, index))
	if index > array.size {
		return nil, outOfRange
	}
	return &array.items[index], nil
}

func (array *Array[Type]) Remove(index uint64) *gopolutils.Exception {
	var outOfRange *gopolutils.Exception = gopolutils.NewNamedException("IndexOutOfRangeError", fmt.Sprintf("Can not access array of size %d at index %d.", array.size, index))
	if index > array.size {
		return outOfRange
	}
	array.items = append(array.items[:index], array.items[index+1:]...)
	array.size--
	return nil
}

func (array Array[Type]) Items() *[]Type {
	return &array.items
}

func (array Array[_]) Size() uint64 {
	return array.size
}

func (array Array[_]) IsEmpty() bool {
	return len(array.items) == 0 && array.size == 0
}
