# Table
Table is a package that defines utilities to aid in the creation, access, connection, insertion, and deletion &mdash; among other behaviour &mdash; of a database table.
The package defines interfaces, enums, and funtions to make interacting with a database table just *that* much easier.
## Interfaces
The interfaces defined in this package are as follows:
- Table: The interface to define each method related to interacting with a database table. Table is to implement a [Sized](/docs/en-UK/collections/sized.md) interface.

- Creator: For creating a table withing a database with a given name. This is used for defining the schema of the database.

- Setter: For defining public setters for the table name and connection.

- Getter: For defining public getters for the table name and connection.

- Inserter: For defining methods related to inserted data into the table.

- Accessor: For defining methods related to obtaining data from the table.

- Remover: For defining methods related to removing data from the table.

- Dropper: For defining how a table should be dropped.

- Closer: For defining how a database should close its connection.
## Field
Other types defined are enums, such as the `Field` type. This type is to be extended by any concrete implementation of a table. The only field defined by default is `Id` which is equal to the string value of `"id"`.
To extend the `Field` enum, as similar definition can be used as described:
```go
import "github.com/Polshkrev/gopolutils/table"

var (
    username Field = "username"
    password Field = "password"
    // Obtains the comma-seperated string representation of the fields as the columns to the database.
    fieldString string = GetFields(table.Id, username, password)
)
```
## Connect
To connect to a database, a function `Connect` is defined. This function takes in a [Path](/docs/en-UK/fayl/path.md) and a `Driver` enum, and returns either a handle to the database based on an internal concurrent mapping or an exception.