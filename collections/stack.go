package collections

import (
	"fmt"

	"github.com/Polshkrev/gopolutils"
)

type Stack[Type any] struct {
	items []Type
	size  uint64
}

func NewStack[Type any]() *Stack[Type] {
	var stack *Stack[Type] = new(Stack[Type])
	stack.items = make([]Type, 0)
	stack.size = 0
	return stack
}

func (stack *Stack[Type]) Append(item Type) {
	stack.items = append(stack.items, item)
	stack.size++
}

func (stack *Stack[Type]) Extend(items []Type) {
	var item Type
	for _, item = range items {
		stack.Append(item)
	}
}

func (stack Stack[Type]) At(index uint64) (*Type, *gopolutils.Exception) {
	var outOfRange *gopolutils.Exception = gopolutils.NewNamedException("IndexOutOfRangeError", fmt.Sprintf("Can not access stack of size %d at index %d.", stack.size, index))
	if index > stack.size {
		return nil, outOfRange
	}
	return &stack.items[index], nil
}
