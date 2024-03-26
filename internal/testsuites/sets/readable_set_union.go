package sets

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
	"github.com/frederik-jatzkowski/cantor/internal/testsuites/testutils"
)

// RunTestsForReadableSet_Union runs a test suite to check correct implementation of the Union method on
// [pkg/github.com/frederik-jatzkowski/cantor.ReadableSet].
func RunTestsForReadableSet_Union(t *testing.T, constructor Constructor[byte, cantor.ReadableSet[byte]]) {
	t.Run("Union", func(t *testing.T) {
		t.Run("both empty", func(t *testing.T) {
			a := constructor()
			b := constructor()

			actual := a.Union(b)

			for i := byte(0); i < 255; i++ {
				if actual.Contains(i) {
					t.Errorf("contained %d but should not", i)
				}
			}
		})

		t.Run("disjunct", func(t *testing.T) {
			a := constructor(1, 2)
			b := constructor(3, 4)

			actual := a.Union(b)
			expected := []byte{1, 2, 3, 4}

			for i := byte(0); i < 255; i++ {
				if !actual.Contains(i) && testutils.SliceContains(i, expected) {
					t.Errorf("was expected to contain %d but did not", i)
				}

				if actual.Contains(i) && !testutils.SliceContains(i, expected) {
					t.Errorf("contained %d but should not", i)
				}
			}
		})

		t.Run("overlapping", func(t *testing.T) {
			a := constructor(1, 2)
			b := constructor(2, 3)

			actual := a.Union(b)
			expected := []byte{1, 2, 3}

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

			actual := a.Union(b)
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
