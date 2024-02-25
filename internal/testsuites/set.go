package testsuites

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
)

// RunTestsForSet runs a test suite to check correct implementation of
// [pkg/github.com/frederik-jatzkowski/cantor.Set].
func RunTestsForSet(t *testing.T, constructor Constructor[byte, cantor.Set[byte]]) {
	t.Run("IterableContainer", func(t *testing.T) {
		RunTestsForDerivedSet(t, func(elements ...byte) cantor.DerivedSet[byte] {
			return constructor(elements...)
		})
	})

	t.Run("Size", func(t *testing.T) {
		t.Run("empty", func(t *testing.T) {
			expected := 0
			actual := constructor().Size()

			if actual != expected {
				t.Errorf("expected size %d but got %d", expected, actual)
			}
		})

		t.Run("some", func(t *testing.T) {
			expected := 4
			actual := constructor([]byte{2, 5, 8, 123}...).Size()

			if actual != expected {
				t.Errorf("expected size %d but got %d", expected, actual)
			}
		})

		t.Run("full", func(t *testing.T) {
			expected := 256
			actual := constructor(allBytes()...).Size()

			if actual != expected {
				t.Errorf("expected size %d but got %d", expected, actual)
			}
		})
	})
}
