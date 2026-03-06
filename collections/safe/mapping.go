package safe

import "github.com/Polshkrev/gopolutils/collections"

type Mapping[Key, Value any] interface {
	Lockable
	Unlockable
	collections.Mapping[Key, Value]
}
