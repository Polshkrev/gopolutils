package collections

// Representation of a pair — or de facto tuple — structure.
type Pair[First any, Second any] struct {
	first  First
	second Second
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
func (pair Pair[First, _]) First() *First {
	return &pair.first
}

// Return a pointer to the second property of the pair.
func (pair Pair[_, Second]) Second() *Second {
	return &pair.second
}

// Swap two pairs with the same types.
// Both the original pair and the operand passed into the function will be modified.
func (pair *Pair[First, Second]) Swap(operand *Pair[First, Second]) {
	var newPair *Pair[First, Second] = NewPair(*operand.First(), *operand.Second())
	*operand = *pair
	*pair = *newPair
	newPair = nil
}

// Flip the values of a pair.
// Returns a new pair where the first value of the original pair is second and the second value of the original pair is first.
func (pair *Pair[First, Second]) Flip() *Pair[Second, First] {
	return NewPair(pair.second, pair.first)
}

// Get a tuple of each of the properties in the pair.
// Returns a pointer to each of the properties in the pair.
func (pair Pair[First, Second]) Items() (*First, *Second) {
	return &pair.first, &pair.second
}
