package collections

// Interface to define a read only sized view into a structure.
type View[Type any] interface {
	// Collect the structure into a view.
	// Returns a slice of the type of the structure.
	Collect() []Type
	Sized
}
