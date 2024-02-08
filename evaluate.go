package cantor

func evaluate[T comparable](set EnumerableSet[T]) ExplicitSet[T] {
	result := explicitSet[T]{
		elements: make(map[T]struct{}),
	}

	set.Enumerate(func(element T) (stop bool) {
		result.elements[element] = struct{}{}

		return false
	})

	return result
}
