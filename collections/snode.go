package collections

import (
	"fmt"

	"github.com/Polshkrev/gopolutils"
)

// Representation of a node within a singly-linked list.
type SingleNode[Type any] struct {
	data Type
	next *SingleNode[Type]
}

// Construct a new [SingleNode] with given data for a singly-linked list.
// Returns a new [SingleNode] with the given data.
func NewSingleNode[Type any](data Type) *SingleNode[Type] {
	var node *SingleNode[Type] = new(SingleNode[Type])
	node.data = data
	node.next = nil
	return node
}

// Set the data stored at the node.
func (node *SingleNode[Type]) SetData(data Type) {
	node.data = data
}

// Link a given node to the node.
func (node *SingleNode[Type]) SetNext(next *SingleNode[Type]) {
	node.next = next
}

// Obtain the data of the node.
// Returns the data of the node.
func (node SingleNode[Type]) Data() *Type {
	return &node.data
}

// Obtain the node with which the node has been linked.
// IF the
func (node SingleNode[Type]) Next() (*SingleNode[Type], *gopolutils.Exception) {
	if node.HasNext() {
		return nil, gopolutils.NewNamedException(gopolutils.ValueError, "No next node has been defined.")
	}
	return node.next, nil
}

// Determine if the node has a linked node.
// Returns true if the node has a linked node, else false.
func (node SingleNode[Type]) HasNext() bool {
	return node.next != nil
}

// Obtain a string representation of the node.
// Returns a string representation of the data stored within the node.
func (node SingleNode[Type]) String() string {
	return fmt.Sprintf("%+v", node.data)
}
