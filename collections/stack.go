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
