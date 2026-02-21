package collections

import "github.com/Polshkrev/gopolutils"

// Implementation of a queue data structure.
type Queue[Type any] struct {
	items []Type
	size  gopolutils.Size
}

// Construct a new queue.
// Returns a pointer to a new queue.
func NewQueue[Type any]() *Queue[Type] {
	var queue *Queue[Type] = new(Queue[Type])
	queue.items = make([]Type, 0)
	queue.size = 0
	return queue
}

// Append an item to the queue.
func (queue *Queue[Type]) Append(item Type) {
	queue.items = append(queue.items, item)
	queue.size++
}

// Append multiple items to the queue.
func (queue *Queue[Type]) Extend(items View[Type]) {
	var item Type
	for _, item = range items.Collect() {
		queue.Append(item)
	}
}

// Access the data stored in the queue at a given index.
// Returns a pointer to data stored in the queue at the given index.
// If the queue is empty, a [gopolutils.ValueError] is returned with a nil data pointer.
// If the index is greater than the size of the queue, an [gopolutils.OutOfRangeError] is returned with a nil data pointer.
func (queue Queue[Type]) At(index gopolutils.Size) (*Type, *gopolutils.Exception) {
	if queue.IsEmpty() {
		return nil, gopolutils.NewNamedException(gopolutils.ValueError, "Can not access an empty queue at index %d.", index)
	} else if index > queue.size {
		return nil, gopolutils.NewNamedException(gopolutils.OutOfRangeError, "Can not access queue of size %d at index %d.", queue.size, index)
	}
	return &queue.items[index], nil
}

// Update a value within the queue.
// If the queue is empty, a [gopolutils.ValueError] is returned.
// If the given index is greater than the queue size, an [gopolutils.OutOfRangeError] is returned.
// If a [gopolutils.ValueError] or an [gopolutils.OutOfRangeError] is returned, the queue is not modified.
func (queue *Queue[Type]) Update(index gopolutils.Size, value Type) *gopolutils.Exception {
	if queue.IsEmpty() {
		return gopolutils.NewNamedException(gopolutils.ValueError, "Can not access an empty queue at index %d.", index)
	} else if index > queue.size {
		return gopolutils.NewNamedException(gopolutils.OutOfRangeError, "Can not access queue of size %d at index %d.", queue.size, index)
	}
	queue.items[index] = value
	return nil
}

// Remove the data stored in the queue at a given index.
// If the queue is empty, a [gopolutils.ValueError] is returned.
// If the given index is greater than the size of the queue, an [gopolutils.OutOfRangeError] is returned.
// If a [gopolutils.ValueError] or an [gopolutils.OutOfRangeError] is returned, the queue is not modified.
func (queue *Queue[_]) Remove(index gopolutils.Size) *gopolutils.Exception {
	if queue.IsEmpty() {
		return gopolutils.NewNamedException(gopolutils.ValueError, "Can not remove from an empty queue at index %d.", index)
	} else if index > queue.size {
		return gopolutils.NewNamedException(gopolutils.OutOfRangeError, "Can not remove element of queue of size %d at index %d.", queue.size, index)
	}
	queue.items = append(queue.items[:index], queue.items[index+1:]...)
	queue.size--
	return nil
}

// Dequeue the first item in the queue.
//
// This is the implementation of a "Fist In First Out" data structure.
// Returns a pointer to the first item in the queue.
// Like the name suggests, when an item is dequeued, the item is removed from the queue.
// If the queue is evaluated to be empty, a [gopolutils.ValueError] is returned with a nil data pointer.
// If a [gopolutils.ValueError] is returned, the queue is not modified.
func (queue *Queue[Type]) Dequeue() (*Type, *gopolutils.Exception) {
	if queue.IsEmpty() {
		return nil, gopolutils.NewNamedException(gopolutils.ValueError, "Can not dequeue from an empty queue.")
	}
	var first Type
	first, queue.items = queue.items[0], queue.items[1:]
	queue.size--
	return &first, nil
}

// Access the first element in the queue.
// Returns a pointer to the first item in the queue.
// If the queue is evaluated to be empty, a [gopolutils.ValueError] is returned with a nil data pointer.
func (queue *Queue[Type]) Peek() (*Type, *gopolutils.Exception) {
	if queue.IsEmpty() {
		return nil, gopolutils.NewNamedException(gopolutils.ValueError, "Can not peek into an empty queue.")
	}
	return &queue.items[0], nil
}

// Determine if the queue is empty.
// Returns true if the length of the underlying data and the size of the queue is equal to 0.
func (queue Queue[_]) IsEmpty() bool {
	return queue.size == 0 && len(queue.items) == 0
}

// Collect the data stored in the queue as a slice.
// Returns a view into the data stored in the queue.
func (queue Queue[Type]) Collect() []Type {
	return queue.items
}

// Get a pointer to the slice of the queue.
// Returns a mutable pointer to the underlying data within the queue.
func (queue Queue[Type]) Items() *[]Type {
	return &queue.items
}

// Access the size of the queue.
// Returns the size of the queue as an unsigned 64-bit integer.
func (queue Queue[_]) Size() gopolutils.Size {
	return queue.size
}
