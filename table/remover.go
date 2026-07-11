package table

import "github.com/Polshkrev/gopolutils"

// Representation of a database remover.
type Remover[Type any] interface {
	// Remove a given item from a database.
	// If the given item can not be removed, an [gopolutils.IOError] is returned.
	Remove(item Type) *gopolutils.Exception
}
