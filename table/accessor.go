package table

import (
	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
)

// Representation of a database accessor.
type Accessor[Type any] interface {
	// Obtain the item stored in a database at the given id.
	// Returns the item stored in a database at the given id.
	// If the item can not be obtained, a [gopolutils.ValueError] is returned with a nil data pointer.
	Get(id gopolutils.Size) (*Type, *gopolutils.Exception)
	// Obtain all the items stored in a database.
	// Returns a [collections.View] of all the items stored in a database.
	// If the items can not be obtained, an [gopolutils.IOError] is returned with a nil data pointer.
	GetAll() (collections.View[*Type], *gopolutils.Exception)
}
