package collections

import (
	"fmt"
	"strings"

	"github.com/Polshkrev/gopolutils"
)

var _ Collection[any] = (*LinkedList[any])(nil)

// Representation of a linked list.
type LinkedList[Type any] struct {
	head *Node[Type]
	size gopolutils.Size
}

// Construct a new linked list.
// Returns a linked list of nodes with a head initially pointing to nil.
func NewLinkedList[Type any]() *LinkedList[Type] {
	var list *LinkedList[Type] = new(LinkedList[Type])
	list.head = nil
	list.size = 0
	return list
}

// Append a given item into the linked list.
func (list *LinkedList[Type]) Append(item Type) {
	if list.head == nil {
		list.head = NewNode(item)
		list.size++
		return
	}
	var last *Node[Type] = list.head
	for last.next != nil {
		last = last.next
	}
	last.next = NewNode(item)
	list.size++
}

// Append a [View] into the list.
func (list *LinkedList[Type]) Extend(items View[Type]) {
	var item Type
	for _, item = range items.Collect() {
		list.Append(item)
	}
}

// Obtain the data stored at the given index within the list.
// Returns a pointer to the data stored at the given idex of the list.
// If the list is evaluated to empty, a [gopolutils.ValueError] is returned with a nil data pointer.
// If the given index is greater than or equal to the size of the list, a [gopolutils.ValueError] is returned with a nil data pointer.
func (list LinkedList[Type]) At(index gopolutils.Size) (*Type, *gopolutils.Exception) {
	if list.IsEmpty() {
		return nil, gopolutils.NewNamedException(gopolutils.ValueError, "Can not access an empty list at index %d.", index)
	} else if index >= list.size {
		return nil, gopolutils.NewNamedException(gopolutils.IndexError, "Can not access list of size %d at index %d.", list.size, index)
	}

	var node *Node[Type] = list.head

	for range index {
		node = node.next
	}

	return node.Data(), nil
}

func (list LinkedList[Type]) Collect() []Type {
	var items []Type = make([]Type, 0, list.size)
	var node *Node[Type] = list.head

	for node != nil {
		items = append(items, node.data)
		node = node.next
	}

	return items
}

// Obtain a mutable pointer to the data stored within the list.
// Returns a mutable pointer to the data stored within the list.
func (list LinkedList[Type]) Items() *[]Type {
	var items []Type = make([]Type, 0, list.size)
	var node *Node[Type] = list.head

	for node != nil {
		items = append(items, node.data)
		node = node.next
	}

	return &items
}

// Update the data stored at the given index within the list with a given item.
// If the list is evaluated to empty, a [gopolutils.ValueError] is returned.
// If the given index is greater than or equal to the size of the list, a [gopolutils.ValueError] is returned.
func (list *LinkedList[Type]) Update(index gopolutils.Size, item Type) *gopolutils.Exception {
	if list.IsEmpty() {
		return gopolutils.NewNamedException(gopolutils.ValueError, "Can not update an empty list at index %d.", index)
	} else if index >= list.size {
		return gopolutils.NewNamedException(gopolutils.IndexError, "Can not update list of size %d at index %d.", list.size, index)
	}

	var node *Node[Type] = list.head

	for range index {
		node = node.next
	}

	node.SetData(item)

	return nil
}

// Remove a [Node] stored at the given index within the list.
// If the list is evaluated to empty, a [gopolutils.ValueError] is returned.
// If the given index is greater than or equal to the size of the list, a [gopolutils.ValueError] is returned.
func (list *LinkedList[Type]) Remove(index gopolutils.Size) *gopolutils.Exception {
	if list.IsEmpty() {
		return gopolutils.NewNamedException(gopolutils.ValueError, "Can not remove from an empty list at index %d.", index)
	} else if index >= list.size {
		return gopolutils.NewNamedException(gopolutils.IndexError, "Can not remove from list of size %d at index %d.", list.size, index)
	} else if index == 0 {
		list.head = list.head.next
		list.size--
		return nil
	}

	var previous *Node[Type] = list.head
	var i gopolutils.Size
	for i = 1; i < index; i++ {
		previous = previous.next
	}

	previous.next = previous.next.next
	list.size--

	return nil
}

// Obtain a string representation of the list.
// Returns a string representation of the list.
func (list LinkedList[Type]) String() string {
	if list.head == nil {
		return "[]"
	}
	var last *Node[Type] = list.head
	var buffer strings.Builder = strings.Builder{}
	buffer.WriteString("[")
	buffer.WriteString(fmt.Sprintf("%v", last.data))
	for last.next != nil {
		last = last.next
		fmt.Fprintf(&buffer, ", %v", last.data)
	}
	buffer.WriteString("]")
	return buffer.String()
}

// Determine if the list is empty.
// Returns true if the head of the list is nil and the size is equal to zero, else false.
func (list LinkedList[Type]) IsEmpty() bool {
	return list.head == nil && list.size == 0
}

// Obtain the size of the list.
// Returns the size of the list.
func (list LinkedList[Type]) Size() gopolutils.Size {
	return list.size
}
