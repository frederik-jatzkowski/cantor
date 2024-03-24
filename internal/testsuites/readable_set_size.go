package testsuites

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
)

// RunTestsForReadableSet_Size runs a test suite to check correct implementation of the Size method on
// [pkg/github.com/frederik-jatzkowski/cantor.ReadableSet].
func RunTestsForReadableSet_Size(t *testing.T, constructor Constructor[byte, cantor.ReadableSet[byte]]) {
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
			actual := constructor(AllBytes()...).Size()

			if actual != expected {
				t.Errorf("expected size %d but got %d", expected, actual)
			}
		})
	})
}
