package collections

import "github.com/Polshkrev/gopolutils"

type Collection[Type any] interface {
	Append(item Type)
	Extend(items []Type)
	At(index uint64) (*Type, *gopolutils.Exception)
	Remove(index uint64) *gopolutils.Exception
}
