package collections

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
