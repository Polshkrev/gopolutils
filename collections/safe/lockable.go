package safe

// Standardization of a lockable collection.
type Lockable interface {
	// Lock the internal mutex of the collection for both reading and writing.
	Lock()
	// Lock the internal mutex of the collection for reading.
	RLock()
}
