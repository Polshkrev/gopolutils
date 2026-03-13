package safe

import "sync"

// Representation of a pair — or de facto tuple — structure.
type Pair[First any, Second any] struct {
	firstLock  sync.RWMutex
	first      First
	secondLock sync.RWMutex
	second     Second
}

// Construct a new pair given two arguments of each respective type variables.
// This constructor resturns a pointer to a newly constructed pair.
func NewPair[First any, Second any](first First, second Second) *Pair[First, Second] {
	var pair *Pair[First, Second] = new(Pair[First, Second])
	pair.first = first
	pair.second = second
	return pair
}

// Return a pointer to the first property of the pair.
func (pair *Pair[First, _]) First() *First {
	pair.RLock()
	defer pair.RUnlock()
	return &pair.first
}

// Return a pointer to the second property of the pair.
func (pair *Pair[_, Second]) Second() *Second {
	pair.RLock()
	defer pair.RUnlock()
	return &pair.second
}

// Set the first property of the pair.
func (pair *Pair[First, _]) SetFirst(first First) {
	pair.Lock()
	defer pair.Unlock()
	pair.first = first
}

// Set the second property of the pair.
func (pair *Pair[_, Second]) SetSecond(second Second) {
	pair.Lock()
	defer pair.Unlock()
	pair.second = second
}

// Set each of the properties of the pair.
func (pair *Pair[First, Second]) Set(first First, second Second) {
	pair.Lock()
	defer pair.Unlock()
	pair.first = first
	pair.second = second
}

// Swap two pairs with the same types.
// Both the original pair and the operand passed into the function will be modified.
func (pair *Pair[First, Second]) Swap(operand *Pair[First, Second]) {
	operand.Lock()
	defer operand.Unlock()
	pair.Lock()
	defer pair.Unlock()
	var newPair *Pair[First, Second] = NewPair(*operand.First(), *operand.Second())
	newPair.RLock()
	defer newPair.RUnlock()
	operand.Set(pair.first, pair.second)
	pair.Set(newPair.first, newPair.second)
}

// Flip the values of a pair.
// Returns a new pair where the first value of the original pair is second and the second value of the original pair is first.
func (pair *Pair[First, Second]) Flip() *Pair[Second, First] {
	pair.RLock()
	defer pair.RUnlock()
	return NewPair(pair.second, pair.first)
}

// Get a tuple of each of the properties in the pair.
// Returns a pointer to each of the properties in the pair.
func (pair *Pair[First, Second]) Items() (*First, *Second) {
	pair.RLock()
	defer pair.RUnlock()
	return &pair.first, &pair.second
}

// Lock the internal mutex of the pair for both reading and writing.
func (pair *Pair[_, _]) Lock() {
	pair.firstLock.Lock()
	pair.secondLock.Lock()
}

// Unlock the internal mutex of the pair for both reading and writing.
func (pair *Pair[_, _]) Unlock() {
	pair.firstLock.Unlock()
	pair.secondLock.Unlock()
}

// Lock the internal mutex of the pair for reading.
func (pair *Pair[_, _]) RLock() {
	pair.firstLock.RLock()
	pair.secondLock.RLock()
}

// Unock the internal mutex of the pair for reading.
func (pair *Pair[_, _]) RUnlock() {
	pair.firstLock.RUnlock()
	pair.secondLock.RUnlock()
}
