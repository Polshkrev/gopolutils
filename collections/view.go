package collections

// Interface to define a read only sized view into a structure.
type View[Type any] interface {
	// Collect the structure into a view.
	Collect() []Type
	Sized
}
