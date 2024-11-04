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

func (stack *Stack[Type]) Remove(index uint64) *gopolutils.Exception {
	var notImplemented *gopolutils.Exception = gopolutils.NewNamedException("NotImplementedError", "Can not remove by index from a stack. Try using the pop method.")
	return notImplemented
}

func (stack *Stack[Type]) Pop() (*Type, *gopolutils.Exception) {
	var empty *gopolutils.Exception = gopolutils.NewException("Can not pop from an empty stack.")
	if len(stack.items) == 0 {
		return nil, empty
	}
	var index int = len(stack.items) - 1
	var last Type
	last, stack.items = stack.items[index], stack.items[:index]
	stack.size--
	return &last, nil
}

func (stack *Stack[Type]) Peek() (*Type, *gopolutils.Exception) {
	var empty *gopolutils.Exception = gopolutils.NewException("Can not peek from an empty stack.")
	if len(stack.items) == 0 {
		return nil, empty
	}
	var index int = len(stack.items) - 1
	return &stack.items[index], nil
}
