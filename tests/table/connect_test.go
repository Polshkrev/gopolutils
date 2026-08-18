package tests

import (
	"database/sql"
	"fmt"
	"sync"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/fayl"
	"github.com/Polshkrev/gopolutils/table"
	"github.com/Polshkrev/gopolutils/table/connect"
)

func TestConnectSqlite(t *testing.T) {
	var path *fayl.Path = fayl.PathFrom(t.TempDir()).JoinAs("database.db")

	var database *sql.DB
	var exception *gopolutils.Exception

	database, exception = connect.Connect(table.Sqlite, path)
	if exception != nil {
		t.Fatal(exception)
	}
	defer database.Close()

	if pingError := database.Ping(); pingError != nil {
		t.Fatal(pingError)
	}
}

func TestConnectInvalidDriver(t *testing.T) {
	var path *fayl.Path = fayl.PathFrom(t.TempDir()).JoinAs("database.db")

	var database *sql.DB
	var exception *gopolutils.Exception

	database, exception = connect.Connect(table.Driver("invalid"), path)
	if exception == nil {
		t.Fatal("expected exception")
	} else if database != nil {
		t.Fatal("expected nil database")
	} else if !exception.Is(gopolutils.KeyError) {
		t.Fatalf("expected KeyError, got %s", exception.Name())
	}
}

func TestConnectNilPath(t *testing.T) {
	var database *sql.DB
	var exception *gopolutils.Exception

	database, exception = connect.Connect(table.Sqlite, nil)
	if exception == nil {
		t.Fatal("expected exception")
	} else if database != nil {
		t.Fatal("expected nil database")
	}
}

func TestConnectMultipleConnections(t *testing.T) {
	var path *fayl.Path = fayl.PathFrom(t.TempDir()).JoinAs("database.db")
	for range 5 {
		var database *sql.DB
		var exception *gopolutils.Exception
		var pingError error
		database, exception = connect.Connect(table.Sqlite, path)
		if exception != nil {
			t.Fatal(exception)
		} else if pingError = database.Ping(); pingError != nil {
			database.Close()
			t.Fatal(pingError)
		}

		database.Close()
	}
}

func TestConnectConcurrent(t *testing.T) {
	var path *fayl.Path = fayl.PathFrom(t.TempDir()).JoinAs("database.db")

	const workers int = 32

	var wait sync.WaitGroup
	wait.Add(workers)
	for range workers {
		go func() {
			defer wait.Done()
			var database *sql.DB
			var exception *gopolutils.Exception
			var pingError error
			database, exception = connect.Connect(table.Sqlite, path)
			if exception != nil {
				t.Error(exception)
				return
			}
			defer database.Close()

			if pingError = database.Ping(); pingError != nil {
				t.Error(pingError)
			}
		}()
	}

	wait.Wait()
}

// func TestConnectCreatesDatabaseFile(t *testing.T) {
// 	var path *fayl.Path = fayl.PathFrom(t.TempDir()).JoinAs("database.db")

// 	if path.Exists() {
// 		t.Fatal("database should not exist")
// 	}
// 	var database *sql.DB
// 	var exception *gopolutils.Exception
// 	database, exception = connect.Connect(table.Sqlite, path)
// 	if exception != nil {
// 		t.Fatal(exception)
// 	}
// 	database.Close()

// 	if !path.Exists() {
// 		t.Fatal("database file was not created")
// 	}
// }

func TestConnectDifferentDatabases(t *testing.T) {
	var i int
	for i = range 5 {
		var path *fayl.Path = fayl.PathFrom(t.TempDir()).JoinAs(fmt.Sprintf("database%d.db", i))
		var database *sql.DB
		var exception *gopolutils.Exception
		var pingError error
		database, exception = connect.Connect(table.Sqlite, path)
		if exception != nil {
			t.Fatal(exception)
		} else if pingError = database.Ping(); pingError != nil {
			database.Close()
			t.Fatal(pingError)
		}

		database.Close()
	}
}

func TestConnectInvalidDriverConcurrent(t *testing.T) {
	const workers int = 16

	var wait sync.WaitGroup
	wait.Add(workers)

	for range workers {
		go func() {
			defer wait.Done()
			var database *sql.DB
			var exception *gopolutils.Exception
			database, exception = connect.Connect(
				table.Driver("invalid"),
				fayl.PathFrom("database.db"),
			)

			if exception == nil {
				t.Error("expected exception")
			} else if database != nil {
				t.Error("expected nil database")
			}
		}()
	}

	wait.Wait()
}
