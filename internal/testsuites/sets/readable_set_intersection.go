package sets

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
	"github.com/frederik-jatzkowski/cantor/internal/testsuites/testutils"
)

// RunTestsForReadableSet_Intersection runs a test suite to check correct implementation of the Intersection method on
// [pkg/github.com/frederik-jatzkowski/cantor.ReadableSet].
func RunTestsForReadableSet_Intersection(t *testing.T, constructor Constructor[byte, cantor.ReadableSet[byte]]) {
	t.Run("Intersection", func(t *testing.T) {
		t.Run("one empty", func(t *testing.T) {
			a := constructor()
			b := constructor(1, 2, 3)

			actual := a.Intersect(b)

			for i := byte(0); i < 255; i++ {
				if actual.Contains(i) {
					t.Errorf("contained %d but should not", i)
				}
			}
		})

		t.Run("nonempty disjunct", func(t *testing.T) {
			a := constructor(1, 2)
			b := constructor(3, 4)

			actual := a.Intersect(b)

			for i := byte(0); i < 255; i++ {
				if actual.Contains(i) {
					t.Errorf("contained %d but should not", i)
				}
			}
		})

		t.Run("overlapping", func(t *testing.T) {
			a := constructor(1, 2, 3)
			b := constructor(2, 3, 4)

			actual := a.Intersect(b)
			expected := []byte{2, 3}

			for i := byte(0); i < 255; i++ {
				if !actual.Contains(i) && testutils.SliceContains(i, expected) {
					t.Errorf("was expected to contain %d but did not", i)
				}

				if actual.Contains(i) && !testutils.SliceContains(i, expected) {
					t.Errorf("contained %d but should not", i)
				}
			}
		})

		t.Run("equal", func(t *testing.T) {
			a := constructor(1, 2)
			b := constructor(1, 2)

			actual := a.Intersect(b)
			expected := []byte{1, 2}

			for i := byte(0); i < 255; i++ {
				if !actual.Contains(i) && testutils.SliceContains(i, expected) {
					t.Errorf("was expected to contain %d but did not", i)
				}

				if actual.Contains(i) && !testutils.SliceContains(i, expected) {
					t.Errorf("contained %d but should not", i)
				}
			}
		})
	})
}
