package safe

import "github.com/Polshkrev/gopolutils/collections"

type Mapping[Key, Value any] interface {
	// Lock the internal mutex of the mapping for both reading and writing.
	Lock()
	// Unlock the internal mutex of the mapping for both reading and writing.
	Unlock()
	// Lock the internal mutex of the mapping for reading.
	RLock()
	// Unock the internal mutex of the mapping for reading.
	RUnlock()
	collections.Mapping[Key, Value]
}
