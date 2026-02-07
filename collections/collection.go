package collections

import "github.com/Polshkrev/gopolutils"

// Interface to standardize a linear data structure.
type Collection[Type any] interface {
	// Append an item to the collection.
	Append(item Type)
	// Append multiple items to the collection.
	Extend(items View[Type])
	// Access an item within the collection at a given index.
	// Returns a pointer to the data in the collection at the given index.
	// If the collection is empty, a [gopolutils.ValueError] is returned with a nil data pointer.
	// If the given index is greater than the size of the collection, an [gopolutils.OutOfRangeError] is returned with a nil data pointer.
	At(index gopolutils.Size) (*Type, *gopolutils.Exception)
	// Update a value within the collection.
	// If the collection is empty, an [gopolutils.OutOfRangeError] is returned.
	// If the given index is greater than the collection size, an [gopolutils.OutOfRangeError] is returned.
	// If an [gopolutils.OutOfRangeError] is returned, the collection is not modified.
	Update(index gopolutils.Size, value Type) *gopolutils.Exception
	// Remove the data in the collection at a given index.
	// If the collection is empty, a [gopolutils.ValueError] is returned.
	// If the given index is greater than the size of the collection, an [gopolutils.OutOfRangeError] is returned.
	// If an [gopolutils.OutOfRangeError] is returned, the collection is not modified.
	Remove(index gopolutils.Size) *gopolutils.Exception
	// Access a pointer to a slice of the data within the collection.
	// This method is called when the data stored in the collection is determined to be internally mutable, or a mutable pointer is needed to access the data.
	// Returns a mutable pointer to the underlying data within the collection.
	Items() *[]Type
	View[Type]
}
