package collections

// Implementation of a wrapper around a collection.
type Wrapper[Type any] interface {
	// Convert a wrapper into a collection.
	// Returns a collection data type contatining all the elements from the wrapped collection.
	Into() Collection[Type]
	// Convert a collection into a wrapper.
	From(collection View[Type])
	View[Type]
}
