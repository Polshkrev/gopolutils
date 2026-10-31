package collections

import (
	"github.com/Polshkrev/gopolutils"
)

var _ Collection[any] = (*Deque[any])(nil)

// Representation of a double-ended queue.
type Deque[Type any] struct {
	*DoublyLinkedList[Type]
}

// Construct a new [Deque].
// Returns a new [Deque] of nodes with a head and tail initially pointing to nil.
func NewDeque[Type any]() *Deque[Type] {
	var list *Deque[Type] = new(Deque[Type])
	list.DoublyLinkedList = NewDoublyLinkedList[Type]()
	return list
}

// Link a given [SingleNode] to the front of the [Deque].
func (list *Deque[Type]) linkFront(node *Node[Type]) {
	node.previous = nil
	node.next = list.head
	if list.head == nil {
		list.tail = node
	} else {
		list.head.previous = node
	}
	list.head = node
}

// Move the element at the given index to the front of the list.
// If the [Deque] is empty, a [gopolutils.ValueError] is returned.
// If the given index is greater than or equal to the size of the list, an [gopolutils.IndexError] is returned.
func (list *Deque[Type]) Use(index gopolutils.Size) *gopolutils.Exception {
	if list.IsEmpty() {
		return gopolutils.NewNamedException(gopolutils.ValueError, "Can not move an element in an empty list at index %d.", index)
	} else if index >= list.size {
		return gopolutils.NewNamedException(gopolutils.IndexError, "Can not move an element in a list of size %d at index %d.", list.size, index)
	} else if index == 0 {
		return nil
	}
	var node *Node[Type] = list.traverse(index)
	list.unlink(node)
	list.linkFront(node)
	return nil
}
