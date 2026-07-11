package table

import (
	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
)

// Representation of a database inserter.
type Inserter[Type any] interface {
	// Insert a given record into the database.
	// If the given record can not be inserted, an [gopolutils.IOError] is returned.
	Insert(record Type) *gopolutils.Exception
	// Insert a [collections.View] of records into the database.
	// If the given records can not be inserted, an [gopolutils.IOError] is returned.
	InsertMany(records collections.View[Type]) *gopolutils.Exception
}
