package collections

type Queue[Type any] struct {
	items []Type
	size  uint64
}
