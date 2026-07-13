package table

import "github.com/Polshkrev/gopolutils"

// Representation of a database driver.
type Driver gopolutils.StringEnum

const (
	Sqlite Driver = "sqlite3" // An sqlite driver.
)
