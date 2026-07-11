package table

import "database/sql"

// Representation of a database getter.
type Getter interface {
	// Obtain the name of the database.
	Name() string
	// Obtain a handle to the database.
	Connection() *sql.DB
}
