package collections

type Map[Key comparable, Value any] struct {
	items map[Key]Value
	size  uint64
}

func NewMap[Key comparable, Value any]() *Map[Key, Value] {
	var mapping *Map[Key, Value] = new(Map[Key, Value])
	mapping.items = make(map[Key]Value, 0)
	mapping.size = 0
	return mapping
}
