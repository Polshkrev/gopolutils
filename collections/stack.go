package collections

import (
	"fmt"

	"github.com/Polshkrev/gopolutils"
)

// Implementation of a stack.
type Stack[Type any] struct {
	items []Type
	size  uint64
}

// Construct a new stack.
// Returns a pointer to a new stack.
func NewStack[Type any]() *Stack[Type] {
	var stack *Stack[Type] = new(Stack[Type])
	stack.items = make([]Type, 0)
	stack.size = 0
	return stack
}

// Append an item to the stack.
func (stack *Stack[Type]) Append(item Type) {
	stack.items = append(stack.items, item)
	stack.size++
}

// Append multiple items to the stack.
func (stack *Stack[Type]) Extend(items View[Type]) {
	var item Type
	for _, item = range items.Collect() {
		stack.Append(item)
	}
}

// Access the data stored on the stack at a given index.
// Returns a pointer to the data stored on the stack at the given index.
// If the stack is evaluated to be empty, an IndexOutOfRangeError is returned with a nil data pointer.
// If the index is greater than the size of the stack, an IndexOutOfRangeError is returned with a nil data pointer.
func (stack Stack[Type]) At(index uint64) (*Type, *gopolutils.Exception) {
	if stack.IsEmpty() {
		return nil, gopolutils.NewNamedException("IndexOutOfRangeError", fmt.Sprintf("Can not access an empty stack at index %d.", index))
	} else if index > stack.size {
		return nil, gopolutils.NewNamedException("IndexOutOfRangeError", fmt.Sprintf("Can not access stack of size %d at index %d.", stack.size, index))
	}
	return &stack.items[index], nil
}

// Update a value within the stack.
// If the given index is greater than the stack size, an IndexOutOfRangeError is returned.
// If the stack is empty, an IndexOutOfRangeError is returned.
// If an IndexOutOfRangeError is returned, the stack is not modified.
func (stack *Stack[Type]) Update(index uint64, value Type) *gopolutils.Exception {
	if stack.IsEmpty() {
		return gopolutils.NewNamedException("IndexOutOfRangeError", fmt.Sprintf("Can not access an empty stack at index %d.", index))
	} else if index > stack.size {
		return gopolutils.NewNamedException("IndexOutOfRangeError", fmt.Sprintf("Can not access stack of size %d at index %d.", stack.size, index))
	}
	stack.items[index] = value
	return nil
}

// Remove the data stored on the stack at a given index.
// This method is currently not implemented.
// If the given index is greater than the size of the stack, an IndexOutOfRangeError is returned.
func (stack *Stack[_]) Remove(index uint64) *gopolutils.Exception {
	return gopolutils.NewNamedException("NotImplementedError", "Can not remove by index from a stack. Try using the pop method.")
}

// Pop the last appended item from the stack.
//
// This is the implementation of a "First In Last Out" data structure.
// As the name suggests, when the last item is popped off the stack, it is also removed from the stack.
// Returns a pointer to the last item in the stack.
// If the stack is evaluated to be empty, an Exception is returned with a nil data pointer.
// IF an Exception is returned, the stack is not modified.
func (stack *Stack[Type]) Pop() (*Type, *gopolutils.Exception) {
	if stack.IsEmpty() {
		return nil, gopolutils.NewException("Can not pop from an empty stack.")
	}
	var index int = len(stack.items) - 1
	var last Type
	last, stack.items = stack.items[index], stack.items[:index]
	stack.size--
	return &last, nil
}

// Peek at the last appended item from the stack.
//
// This is the implementation of a "First In Last Out" data structure.
// Unlike pop, this method accesses the data on the stack without modifying the stack itself.
// Returns a pointer to the last item in the stack.
// If the stack is evaluated to be empty, an Exception is returned with a nil data pointer.
func (stack *Stack[Type]) Peek() (*Type, *gopolutils.Exception) {
	if stack.IsEmpty() {
		return nil, gopolutils.NewException("Can not peek from an empty stack.")
	}
	var index int = len(stack.items) - 1
	return &stack.items[index], nil
}

// Determine if the stack is empty.
// Returns true if the length of the underlying data and the size of the stack is equal to 0.
func (stack Stack[_]) IsEmpty() bool {
	return len(stack.items) == 0 && stack.size == 0
}

// Collect the data stored in the stack as a slice.
// Returns a view into the data stored in the stack.
func (stack Stack[Type]) Collect() []Type {
	return stack.items
}

// Get a pointer to a slice of the data within stack.
// Returns a mutable pointer to the underlying data within the stack.
func (stack Stack[Type]) Items() *[]Type {
	return &stack.items
}

// Access the size of the stack.
// Returns the size of the stack as an unsigned 64-bit integer.
func (stack Stack[_]) Size() uint64 {
	return stack.size
}
