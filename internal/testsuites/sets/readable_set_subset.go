package sets

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
)

// RunTestsForReadableSet_Subset runs a test suite to check correct implementation of the Subset method on
// [pkg/github.com/frederik-jatzkowski/cantor.ReadableSet].
func RunTestsForReadableSet_Subset(t *testing.T, constructor Constructor[byte, cantor.ReadableSet[byte]]) {
	t.Run("Subset", func(t *testing.T) {
		t.Run("subset", func(t *testing.T) {
			a := constructor(1, 2)
			b := cantor.NewHashSet[byte](1, 2, 3)

			if !a.Subset(b) {
				t.Errorf("a should be subset of b but was not")
			}
		})

		t.Run("overlapping", func(t *testing.T) {
			a := constructor(1, 2, 3)
			b := constructor(2, 3, 4)

			if a.Subset(b) {
				t.Errorf("a should not be subset of b but was")
			}
		})

		t.Run("equal", func(t *testing.T) {
			a := constructor(1, 2)
			b := constructor(1, 2)

			if !a.Subset(b) {
				t.Errorf("a should be subset of b but was not")
			}
		})
	})
}
