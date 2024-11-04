package collections

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
