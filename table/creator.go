package table

import "github.com/Polshkrev/gopolutils"

// Representation of a database creator.
type Creator interface {
	// Create a database with a given name.
	// If the database can not be created, an [gopolutils.IOError] is returned.
	Create(name string) *gopolutils.Exception
}
