package collections

import (
	"fmt"
	"sync"

	"github.com/Polshkrev/gopolutils"
)

// Implementation of a concurrent-safe queue data structure.
type SafeQueue[Type any] struct {
	lock  sync.RWMutex
	items []Type
	size  uint64
}

// Construct a new queue.
// Returns a pointer to a new queue.
func NewSafeQueue[Type any]() *SafeQueue[Type] {
	var queue *SafeQueue[Type] = new(SafeQueue[Type])
	queue.items = make([]Type, 0)
	queue.size = 0
	return queue
}

// Append an item to the queue.
func (queue *SafeQueue[Type]) Append(item Type) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	queue.items = append(queue.items, item)
	queue.size++
}

// Append multiple items to the queue.
func (queue *SafeQueue[Type]) Extend(items View[Type]) {
	var item Type
	for _, item = range items.Collect() {
		queue.Append(item)
	}
}

// Access the data stored in the queue at a given index.
// Returns a pointer to data stored in the queue at the given index.
// If the queue is empty, a ValueError is returned with a nil data pointer.
// If the index is greater than the size of the queue, an OutOfRangeError is returned with a nil data pointer.
func (queue *SafeQueue[Type]) At(index uint64) (*Type, *gopolutils.Exception) {
	queue.lock.RLock()
	defer queue.lock.RUnlock()
	if queue.IsEmpty() {
		return nil, gopolutils.NewNamedException(gopolutils.ValueError, fmt.Sprintf("Can not access an empty queue at index %d.", index))
	} else if index > queue.size {
		return nil, gopolutils.NewNamedException(gopolutils.OutOfRangeError, fmt.Sprintf("Can not access queue of size %d at index %d.", queue.size, index))
	}
	return &queue.items[index], nil
}

// Update a value within the queue.
// If the queue is empty, a ValueError is returned.
// If the given index is greater than the queue size, an OutOfRangeError is returned.
// If a non-nil Exception is returned, the queue is not modified.
func (queue *SafeQueue[Type]) Update(index uint64, value Type) *gopolutils.Exception {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	if queue.IsEmpty() {
		return gopolutils.NewNamedException(gopolutils.ValueError, fmt.Sprintf("Can not access an empty queue at index %d.", index))
	} else if index > queue.size {
		return gopolutils.NewNamedException(gopolutils.OutOfRangeError, fmt.Sprintf("Can not access queue of size %d at index %d.", queue.size, index))
	}
	queue.items[index] = value
	return nil
}

// Remove the data stored in the queue at a given index.
// If the queue is empty, a ValueError is returned.
// If the given index is greater than the size of the queue, an OutOfRangeError is returned.
// If a non-nil Exception is returned, the queue is not modified.
func (queue *SafeQueue[_]) Remove(index uint64) *gopolutils.Exception {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	if queue.IsEmpty() {
		return gopolutils.NewNamedException(gopolutils.ValueError, fmt.Sprintf("Can not remove from an empty queue at index %d.", index))
	} else if index > queue.size {
		return gopolutils.NewNamedException(gopolutils.OutOfRangeError, fmt.Sprintf("Can not remove element of queue of size %d at index %d.", queue.size, index))
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
// If the queue is evaluated to be empty, a ValueError is returned with a nil data pointer.
// If a non-nil Exception is returned, the queue is not modified.
func (queue *SafeQueue[Type]) Dequeue() (*Type, *gopolutils.Exception) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
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
// If the queue is evaluated to be empty, a ValueError is returned with a nil data pointer.
func (queue *SafeQueue[Type]) Peek() (*Type, *gopolutils.Exception) {
	queue.lock.RLock()
	defer queue.lock.RUnlock()
	if queue.IsEmpty() {
		return nil, gopolutils.NewNamedException(gopolutils.ValueError, "Can not peek into an empty queue.")
	}
	return &queue.items[0], nil
}

// Determine if the queue is empty.
// Returns true if the length of the underlying data and the size of the queue is equal to 0.
func (queue *SafeQueue[_]) IsEmpty() bool {
	queue.lock.RLock()
	defer queue.lock.RUnlock()
	return queue.size == 0 && len(queue.items) == 0
}

// Collect the data stored in the queue as a slice.
// Returns a view into the data stored in the queue.
func (queue *SafeQueue[Type]) Collect() []Type {
	queue.lock.RLock()
	defer queue.lock.RUnlock()
	return queue.items
}

// Get a pointer to the slice of the queue.
// Returns a mutable pointer to the underlying data within the queue.
func (queue *SafeQueue[Type]) Items() *[]Type {
	queue.lock.RLock()
	defer queue.lock.RUnlock()
	return &queue.items
}

// Access the size of the queue.
// Returns the size of the queue as an unsigned 64-bit integer.
func (queue *SafeQueue[_]) Size() uint64 {
	queue.lock.RLock()
	defer queue.lock.RUnlock()
	return queue.size
}
