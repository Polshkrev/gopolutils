package collections

// Standardisation of a type that can be converted to a [Collection].
type Into[Type any] interface {
	// Convert a wrapper into a collection.
	// Returns a collection data type contatining all the elements from the wrapped collection.
	Into() Collection[Type]
}
