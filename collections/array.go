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

func (array *Array[Type]) Append(item Type) {
	array.items = append(array.items, item)
	array.size++
}

func (array *Array[Type]) Extend(items []Type) {
	for _, item := range items {
		array.Append(item)
	}
}
