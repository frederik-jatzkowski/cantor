package cantor

import (
	"fmt"
	"strings"
)

func toString[T comparable](set DeduplicatingIterableContainer[T]) string {
	var elements []string

	set.UniqueKeys()(func(element T) (next bool) {
		elements = append(elements, fmt.Sprint(element))

		return true
	})

	return fmt.Sprintf("{%s}", strings.Join(elements, ", "))
}
