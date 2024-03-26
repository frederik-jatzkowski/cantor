package sets

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
)

// RunTestsForReadableSet_Equals runs a test suite to check correct implementation of the Equals method on
// [pkg/github.com/frederik-jatzkowski/cantor.ReadableSet].
func RunTestsForReadableSet_Equals(t *testing.T, constructor Constructor[byte, cantor.ReadableSet[byte]]) {
	t.Run("Equals", func(t *testing.T) {
		t.Run("subset", func(t *testing.T) {
			a := constructor(1, 2, 3)
			b := cantor.NewHashSet[byte](1, 2)

			if a.Equals(b) {
				t.Errorf("sets should not be equal but were")
			}
		})

		t.Run("overlapping", func(t *testing.T) {
			a := constructor(1, 2, 3)
			b := constructor(2, 3, 4)

			if a.Equals(b) {
				t.Errorf("sets should not be equal but were")
			}
		})

		t.Run("equal", func(t *testing.T) {
			a := constructor(1, 2)
			b := constructor(1, 2)

			if !a.Equals(b) {
				t.Errorf("sets should be equal but were not")
			}
		})
	})
}
