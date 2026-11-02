package collections

import (
	"fmt"
	"strings"

	"github.com/Polshkrev/gopolutils"
)

var _ Collection[any] = (*DoublyLinkedList[any])(nil)
var _ Iterable[any] = (*DoublyLinkedList[any])(nil)

// Representation of a doubly-linked list.
type DoublyLinkedList[Type any] struct {
	head *Node[Type]
	tail *Node[Type]
	size gopolutils.Size
}

// Construct a new doubly-linked list.
// Returns a doubly-linked list of nodes with a head and tail initially pointing to nil.
func NewDoublyLinkedList[Type any]() *DoublyLinkedList[Type] {
	var list *DoublyLinkedList[Type] = new(DoublyLinkedList[Type])
	list.head = nil
	list.tail = nil
	list.size = 0
	return list
}

// Obtain a pointer to the head of the list.
// Returns a pointer to the head of the list.
func (list DoublyLinkedList[Type]) Head() Node[Type] {
	return *list.head
}

// Determine if the list has a valid head.
// Returns true if the head of the list is not equal to nil, else false.
func (list DoublyLinkedList[Type]) HasHead() bool {
	return list.head != nil
}

// Set the head of the list.
func (list *DoublyLinkedList[Type]) SetHead(node *Node[Type]) {
	list.head = node
}

// Obtain a pointer to the tail of the list.
// Returns a pointer to the tail of the list.
func (list DoublyLinkedList[Type]) Tail() Node[Type] {
	return *list.tail
}

// Determine if the list has a valid tail.
// Returns true if the tail of the list is not equal to nil, else false.
func (list DoublyLinkedList[Type]) HasTail() bool {
	return list.tail != nil
}

// Set the tail of the list.
func (list *DoublyLinkedList[Type]) SetTail(node *Node[Type]) {
	list.tail = node
}

// Append a given item into the doubly-linked list.
func (list *DoublyLinkedList[Type]) Append(item Type) {
	var node *Node[Type] = NewNode(item)
	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		node.previous = list.tail
		list.tail.next = node
		list.tail = node
	}
	list.size++
}

// Append a [View] into the list.
func (list *DoublyLinkedList[Type]) Extend(items View[Type]) {
	var collected []Type = items.Collect()
	var i int
	for i = range collected {
		list.Append(collected[i])
	}
}

// Traverse the list from the head to the given index.
// Returns A pointer to the [Node] at the given index.
func (list DoublyLinkedList[Type]) traverseForward(index gopolutils.Size) *Node[Type] {
	var node *Node[Type] = list.head
	for range index {
		node = node.next
	}
	return node
}

// Traverse the list from the tail to the given index.
// Returns a pointer to the [Node] at the given index.
func (list DoublyLinkedList[Type]) traverseBackward(index gopolutils.Size) *Node[Type] {
	var node *Node[Type] = list.tail
	var i gopolutils.Size
	for i = list.size - 1; i > index; i-- {
		node = node.previous
	}

	return node
}

// Traverse the list until the given index.
// Returns a pointer to the [Node] at the given index.
func (list DoublyLinkedList[Type]) traverse(index gopolutils.Size) *Node[Type] {
	if index < list.size/2 {
		return list.traverseForward(index)
	}
	return list.traverseBackward(index)
}

// Obtain the data stored at the given index within the list.
// Returns a pointer to the data stored at the given idex of the list.
// If the list is evaluated to empty, a [gopolutils.ValueError] is returned with a nil data pointer.
// If the given index is greater than or equal to the size of the list, a [gopolutils.ValueError] is returned with a nil data pointer.
func (list *DoublyLinkedList[Type]) At(index gopolutils.Size) (*Type, *gopolutils.Exception) {
	if list.IsEmpty() {
		return nil, gopolutils.NewNamedException(gopolutils.ValueError, "Can not access an empty list at index %d.", index)
	} else if index >= list.size {
		return nil, gopolutils.NewNamedException(gopolutils.IndexError, "Can not access list of size %d at index %d.", list.size, index)
	}
	var node *Node[Type] = list.traverse(index)
	return &node.data, nil
}

func (list DoublyLinkedList[Type]) Collect() []Type {
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
func (list *DoublyLinkedList[Type]) Items() *[]Type {
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
func (list *DoublyLinkedList[Type]) Update(index gopolutils.Size, item Type) *gopolutils.Exception {
	if list.IsEmpty() {
		return gopolutils.NewNamedException(gopolutils.ValueError, "Can not update an empty list at index %d.", index)
	} else if index >= list.size {
		return gopolutils.NewNamedException(gopolutils.IndexError, "Can not update list of size %d at index %d.", list.size, index)
	}
	var node *Node[Type] = list.traverse(index)
	node.SetData(item)
	return nil
}

// Replace the previous link of the node with its predecessor.
func (list *DoublyLinkedList[Type]) unlinkPrevious(node *Node[Type]) {
	if node.previous == nil {
		list.head = node.next
		return
	}
	node.previous.next = node.next
}

// Replace the next link of the [Node] with its successor.
func (list *DoublyLinkedList[Type]) unlinkNext(node *Node[Type]) {
	if node.next == nil {
		list.tail = node.previous
		return
	}
	node.next.previous = node.previous
}

// Unlink a [Node] from the list.
func (list *DoublyLinkedList[Type]) unlink(node *Node[Type]) {
	list.unlinkPrevious(node)
	list.unlinkNext(node)
	node.previous = nil
	node.next = nil
}

// Remove a [Node] stored at the given index within the list.
// If the list is evaluated to empty, a [gopolutils.ValueError] is returned.
// If the given index is greater than or equal to the size of the list, a [gopolutils.ValueError] is returned.
func (list *DoublyLinkedList[Type]) Remove(index gopolutils.Size) *gopolutils.Exception {
	if list.IsEmpty() {
		return gopolutils.NewNamedException(gopolutils.ValueError, "Can not access an empty list at index %d.", index)
	} else if index >= list.size {
		return gopolutils.NewNamedException(gopolutils.IndexError, "Can not access list of size %d at index %d.", list.size, index)
	}
	list.unlink(list.traverse(index))
	list.size--
	return nil
}

// Determine if the list is empty.
// Returns true if the head of the list is nil and the size is equal to zero, else false.
func (list DoublyLinkedList[Type]) IsEmpty() bool {
	return list.size == 0
}

// Obtain the size of the list.
// Returns the size of the list.
func (list DoublyLinkedList[Type]) Size() gopolutils.Size {
	return list.size
}

// Obtain an [Iterator] over the data of the collection.
// Returns an [Iterator] the data of the collection.
func (list DoublyLinkedList[Type]) Iterator() *Iterator[Type] {
	return IteratorFrom(list)
}

// Obtain a string representation of the list.
// Returns a string representation of the list.
func (list DoublyLinkedList[Type]) String() string {
	if list.head == nil {
		return "[]"
	}
	var last *Node[Type] = list.head
	var buffer strings.Builder = strings.Builder{}
	buffer.WriteRune('[')
	fmt.Fprint(&buffer, last.data)
	for last.next != nil {
		last = last.next
		fmt.Fprintf(&buffer, ", %v", last.data)
	}
	buffer.WriteRune(']')
	return buffer.String()
}
