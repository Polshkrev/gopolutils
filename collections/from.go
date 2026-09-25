package collections

type From[Type any] interface {
	// Convert a collection into a wrapper.
	From(collection View[Type])
}
