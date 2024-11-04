package collections

import "github.com/Polshkrev/gopolutils"

type Mapping[Key, Value any] interface {
	Insert(key Key, value Value) *gopolutils.Exception
	At(key Key) (*Value, *gopolutils.Exception)
	Keys() []Key
	Values() []Value
	Remove(key Key) *gopolutils.Exception
	HasKey(key Key) bool
	Size() uint64
}
