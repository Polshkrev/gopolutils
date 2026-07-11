package table

import "github.com/Polshkrev/gopolutils/collections"

// Representation of a database table.
type Table[Type any] interface {
	Creater
	Setter
	Getter
	Inserter[Type]
	Accessor[Type]
	Remover[Type]
	Dropper
	Closer
	collections.Sized
}
