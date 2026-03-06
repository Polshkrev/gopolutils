package safe

import "github.com/Polshkrev/gopolutils/collections"

// Representation of a collection of with an internal lock.
type Collection[Type any] interface {
	Lockable
	Unlockable
	collections.Collection[Type]
}
