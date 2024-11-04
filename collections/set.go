package collections

type Set[Type comparable] struct {
	items map[Type]struct{}
	size  uint64
}
