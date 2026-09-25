package collections

// Implementation of a wrapper around a collection.
type Wrapper[Type any] interface {
	View[Type]
}
