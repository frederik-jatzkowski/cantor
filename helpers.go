package cantor

import (
	"fmt"
	"strings"
)

func evaluate[T comparable](set IterableSet[T]) ExplicitSet[T] {
	result := explicitSet[T]{
		elements: make(map[T]struct{}),
	}

	set.Iter()(func(element T) (next bool) {
		result.elements[element] = struct{}{}

		return true
	})

	return result
}

func toString[T comparable](set IterableSet[T]) string {
	var elements []string

	set.Iter()(func(element T) (next bool) {
		elements = append(elements, fmt.Sprint(element))

		return true
	})

	return fmt.Sprintf("{%s}", strings.Join(elements, ", "))
}
