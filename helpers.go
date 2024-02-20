package cantor

import (
	"fmt"
	"strings"
)

func evaluate[T comparable](set ReadableSet[T]) Set[T] {
	result := HashSet[T]{
		elements: make(map[T]struct{}),
	}

	set.IterateDistinct()(func(element T) (next bool) {
		result.elements[element] = struct{}{}

		return true
	})

	return result
}

func toString[T comparable](set ReadableSet[T]) string {
	var elements []string

	set.IterateDistinct()(func(element T) (next bool) {
		elements = append(elements, fmt.Sprint(element))

		return true
	})

	return fmt.Sprintf("{%s}", strings.Join(elements, ", "))
}
