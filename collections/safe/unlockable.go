package safe

// Standardization of a unlockable collection.
type Unlockable interface {
	// Unlock the internal mutex of the collection for both reading and writing.
	Unlock()
	// Unock the internal mutex of the collection for reading.
	RUnlock()
}
