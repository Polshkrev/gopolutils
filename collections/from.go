package collections

// Standardization of a [Collection] that can be converted from another [Collection].
type From[Type any] interface {
	// Convert a collection into a wrapper.
	From(collection View[Type])
}
