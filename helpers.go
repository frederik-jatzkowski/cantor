package cantor

import (
	"fmt"
	"strings"
)

func toString[T comparable](set ReadableSet[T]) string {
	var elements []string

	set.Elements()(func(element T) (next bool) {
		elements = append(elements, fmt.Sprint(element))

		return true
	})

	return fmt.Sprintf("{%s}", strings.Join(elements, ", "))
}

func count[T comparable](iterator Iterator[T]) (result int) {
	iterator(func(element T) (next bool) {
		result++

		return true
	})

	return result
}
