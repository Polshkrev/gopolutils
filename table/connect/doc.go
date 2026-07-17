/*
Connect provides a standardization of connecting to a database.

The example below shows storing passwords in a database.

Example:

At first you will need to implement the concrete implementation of the [Table] interface.
In a package containing the type to store (in this case passwords):

	import "github.com/Polshkrev/gopolutils/table"

	var _ table.Table[password.Password] = (*Table)(nil)

	const (
		TableName string = "passwords"
	)

	var (
		fieldString = table.GetFields( // The variadic table column names )
	)

	type Table struct {
		// Implementation of the table...
	}

	// Methods for the table...

Setup for a database:

	import(
		"github.com/Polshkrev/gopolutils/fayl"
		"github.com/Polshkrev/gopolutils/table"
		"github.com/Polshkrev/gopolutils/table/connect"
	)

	var (
		path *fayl.Path = fayl.PathFrom("path/to/database.db")
	)

	func main() {
		var connection *sql.DB = gopolutils.Must(connect.Connect(driver, path))
		var database table.Table[password.Password] = password.NewTable(connection)
		database.Create(password.TableName)
	}
*/
package connect
