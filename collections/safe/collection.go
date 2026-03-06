package safe

import "github.com/Polshkrev/gopolutils/collections"

// Representation of a concurrent-safe collection.
type Collection[Type any] interface {
	Lockable
	Unlockable
	collections.Collection[Type]
}
