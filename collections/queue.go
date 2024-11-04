package collections

type Queue[Type any] struct {
	items []Type
	size  uint64
}

func NewQueue[Type any]() *Queue[Type] {
	var queue *Queue[Type] = new(Queue[Type])
	queue.items = make([]Type, 0)
	queue.size = 0
	return queue
}

func (queue *Queue[Type]) Append(item Type) {
	queue.items = append(queue.items, item)
	queue.size++
}
