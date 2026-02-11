package safe

import "github.com/Polshkrev/gopolutils/collections"

// Representation of a collection of with an internal lock.
type SafeCollection[Type any] interface {
	// Lock the internal mutex of the collection for both reading and writing.
	Lock()
	// Unlock the internal mutex of the collection for both reading and writing.
	Unlock()
	// Lock the internal mutex of the collection for reading.
	RLock()
	// Unock the internal mutex of the collection for reading.
	RUnlock()
	collections.Collection[Type]
}
