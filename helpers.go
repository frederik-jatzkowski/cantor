package cantor

import (
	"fmt"
	"strings"
)

func evaluate[T comparable](set DeduplicatingIterableContainer[T]) HashSet[T] {
	result := HashSet[T]{
		elements: make(map[T]struct{}),
	}

	set.UniqueElements()(func(element T) (next bool) {
		result.elements[element] = struct{}{}

		return true
	})

	return result
}

func toString[T comparable](set DeduplicatingIterableContainer[T]) string {
	var elements []string

	set.UniqueElements()(func(element T) (next bool) {
		elements = append(elements, fmt.Sprint(element))

		return true
	})

	return fmt.Sprintf("{%s}", strings.Join(elements, ", "))
}
