package collections

// Interface to define a sized structure.
type Sized interface {
	// Retrieve the size of the structure.
	// Returns the size of the structure as an unsigned 64-bit integer.
	Size() uint64
	// Determine if the structure is empty.
	// Returns true if the length of the underlying data stored in the structure and the size of the structure is equal to 0.
	IsEmpty() bool
}
