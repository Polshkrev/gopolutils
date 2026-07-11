package connect

import (
	"database/sql"
	"sync"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/fayl"
	"github.com/Polshkrev/gopolutils/table"
)

// Alias for a callback that returns a connection to database from a given path.
type Connector func(path *fayl.Path) (*sql.DB, *gopolutils.Exception)

var (
	connectorLock *sync.RWMutex = new(sync.RWMutex) // Connector concurrency lock.
	// Driver to connector mapping.
	connectors map[table.Driver]Connector = map[table.Driver]Connector{
		table.Sqlite: connectSqlite,
	}
)

// Connect to a database of a given driver at a given path.
// Returns a connection to a database at a given path.
// If the database can not connect, a [gopolutils.KeyError] is returned.
func Connect(driver table.Driver, path *fayl.Path) (*sql.DB, *gopolutils.Exception) {
	connectorLock.RLock()
	defer connectorLock.RUnlock()
	var connector Connector
	var ok bool
	connector, ok = connectors[driver]
	if !ok {
		return nil, gopolutils.NewNamedException(gopolutils.KeyError, "\"%s\" can not  be found in the driver mapping.", driver)
	}
	return connector(path)
}
