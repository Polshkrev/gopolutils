package collections

type Pair[First any, Second any] struct {
	first  First
	second Second
}

func NewPair[First any, Second any](first First, second Second) *Pair[First, Second] {
	var pair *Pair[First, Second] = new(Pair[First, Second])
	pair.first = first
	pair.second = second
	return pair
}
