// Package testsuites contains test cases for all interfaces exported from [pkg/github.com/frederik-jatzkowski/cantor].
// These can be used to validate implementations of these interfaces.
package testsuites

type Constructor[E any, T any] func(elements ...E) T

func SliceContains[T comparable](e T, elements []T) bool {
	for _, e2 := range elements {
		if e == e2 {
			return true
		}
	}

	return false
}

func allBytes() (result []byte) {
	for i := byte(0); i < 255; i++ {
		result = append(result, i)
	}

	return result
}
