package table

import "database/sql"

// Representation of a database setter.
type Setter interface {
	// Set the name of the database.
	SetName(name string)
	// Set the connection of the database.
	SetConnection(connection *sql.DB)
}
