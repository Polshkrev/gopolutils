package collections

// Implementation of a wrapper around a collection.
type Wrapper[Type any] interface {
	From[Type]
	Into[Type]
	View[Type]
}
