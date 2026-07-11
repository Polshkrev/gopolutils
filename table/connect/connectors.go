package connect

import (
	"database/sql"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/fayl"
	"github.com/Polshkrev/gopolutils/table"
)

// Connect to a sqlite database.
// If the database can not connect, an [gopolutils.IOError] is returned with a nil data pointer.
func connectSqlite(path *fayl.Path) (*sql.DB, *gopolutils.Exception) {
	var storage *sql.DB
	var openError error
	storage, openError = sql.Open(string(table.Sqlite), path.String())
	if openError != nil {
		return nil, gopolutils.NewNamedException(gopolutils.IOError, "%s", openError.Error())
	}
	return storage, nil
}
