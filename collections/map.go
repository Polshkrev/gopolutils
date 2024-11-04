package collections

type Map[Key comparable, Value any] struct {
	items map[Key]Value
	size  uint64
}
