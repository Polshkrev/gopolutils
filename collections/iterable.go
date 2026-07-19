package collections

// Representation of a collection over which can be iterated.
type Iterable[Type any] interface {
	// Obtain an iterator over the data of the collection.
	// Returns an iterator the data of the collection.
	Iterator() Iterator[Type]
}
