package collections

type Stack[Type any] struct {
	items []Type
	size  uint64
}
