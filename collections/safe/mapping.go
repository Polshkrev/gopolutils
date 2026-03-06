package safe

import "github.com/Polshkrev/gopolutils/collections"

// Interface to standardize a concurrent-safe key-value pair mapping.
type Mapping[Key, Value any] interface {
	Lockable
	Unlockable
	collections.Mapping[Key, Value]
}
