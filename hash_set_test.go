package cantor_test

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
	"github.com/frederik-jatzkowski/cantor/internal/testsuites"
)

func TestNewHashSet(t *testing.T) {
	testsuites.RunTestsForSet(t, func(elements ...byte) cantor.Set[byte] {
		return cantor.NewHashSet(elements...)
	})
}

func TestNewHashSetFromIterator(t *testing.T) {
	testsuites.RunTestsForSet(t, func(elements ...byte) cantor.Set[byte] {
		return cantor.NewHashSetFromIterator(func(yield func(element byte) (next bool)) {
			for _, element := range elements {
				if !yield(element) {
					break
				}
			}
		})
	})
}
