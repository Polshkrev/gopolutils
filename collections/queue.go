package collections

import (
	"fmt"

	"github.com/Polshkrev/gopolutils"
)

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

func (queue *Queue[Type]) Extend(items []Type) {
	var item Type
	for _, item = range items {
		queue.Append(item)
	}
}

func (queue Queue[Type]) At(index uint64) (*Type, *gopolutils.Exception) {
	var outOfRange *gopolutils.Exception = gopolutils.NewNamedException("IndexOutOfRangeError", fmt.Sprintf("Can not access queue of size %d at index %d.", queue.size, index))
	if index > queue.size {
		return nil, outOfRange
	}
	return &queue.items[index], nil
}

func (queue *Queue[Type]) Remove(index uint64) *gopolutils.Exception {
	var notImplemented *gopolutils.Exception = gopolutils.NewNamedException("NotImplementedError", "Can not remove by index in a queue. Try using the dequeue method.")
	return notImplemented
}

func (queue *Queue[Type]) Dequeue() (*Type, *gopolutils.Exception) {
	var empty *gopolutils.Exception = gopolutils.NewException("Can not dequeue from an empty queue.")
	if queue.size == 0 {
		return nil, empty
	}
	return &queue.items[0], nil
}
