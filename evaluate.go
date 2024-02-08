package cantor

func evaluate[T comparable](set IterableSet[T]) ExplicitSet[T] {
	result := explicitSet[T]{
		elements: make(map[T]struct{}),
	}

	set.Iter()(func(element T) (stop bool) {
		result.elements[element] = struct{}{}

		return true
	})

	return result
}
