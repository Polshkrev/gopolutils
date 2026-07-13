package table

import "github.com/Polshkrev/gopolutils"

// Representation of a database dropper.
type Dropper interface {
	// Drop a database with a given name.
	// If the database can not be dropped, an [gopolutils.IOError] is returned.
	Drop(name string) *gopolutils.Exception
}
