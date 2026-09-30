package collections

import (
	"fmt"

	"github.com/Polshkrev/gopolutils"
)

// Representation of a node within a singly-linked list.
type Node[Type any] struct {
	data Type
	next *Node[Type]
}

// Construct a new [Node] with given data for a singly-linked list.
// Returns a new [Node] with the given data.
func NewNode[Type any](data Type) *Node[Type] {
	var node *Node[Type] = new(Node[Type])
	node.data = data
	node.next = nil
	return node
}

// Set the data stored at the node.
func (node *Node[Type]) SetData(data Type) {
	node.data = data
}

// Link a given node to the node.
func (node *Node[Type]) SetNext(next *Node[Type]) {
	node.next = next
}

// Obtain the data of the node.
// Returns the data of the node.
func (node Node[Type]) Data() *Type {
	return &node.data
}

// Obtain the node with which the node has been linked.
// IF the
func (node Node[Type]) Next() (*Node[Type], *gopolutils.Exception) {
	if node.HasNext() {
		return nil, gopolutils.NewNamedException(gopolutils.ValueError, "No next node has been defined.")
	}
	return node.next, nil
}

// Determine if the node has a linked node.
// Returns true if the node has a linked node, else false.
func (node Node[Type]) HasNext() bool {
	return node.next != nil
}

// Obtain a string representation of the node.
// Returns a string representation of the data stored within the node.
func (node Node[Type]) String() string {
	return fmt.Sprintf("%+v", node.data)
}
