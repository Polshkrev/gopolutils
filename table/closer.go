package table

import "github.com/Polshkrev/gopolutils"

// Representation of a database closer.
type Closer interface {
	// Close a database.
	// If a database can not close, an [gopolutils.IOError] is returned.
	Close() *gopolutils.Exception
}
