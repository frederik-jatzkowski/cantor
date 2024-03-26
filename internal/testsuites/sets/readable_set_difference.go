package sets

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
)

// RunTestsForReadableSet_Difference runs a test suite to check correct implementation of the Difference method on
// [pkg/github.com/frederik-jatzkowski/cantor.ReadableSet].
func RunTestsForReadableSet_Difference(t *testing.T, constructor Constructor[byte, cantor.ReadableSet[byte]]) {
	t.Run("Difference", func(t *testing.T) {
		t.Run("argument is universal set", func(t *testing.T) {
			a := constructor(1, 2, 3)
			b := cantor.NewImplicitSet[byte](func(element byte) bool { return true })

			actual := a.Difference(b)

			for i := byte(0); i < 255; i++ {
				if actual.Contains(i) {
					t.Errorf("contained %d but should not", i)
				}
			}
		})

		t.Run("argument is subset", func(t *testing.T) {
			a := constructor(1, 2, 3)
			b := cantor.NewHashSet[byte](1, 2)

			actual := a.Difference(b)
			expected := cantor.NewHashSet[byte](3)

			for i := byte(0); i < 255; i++ {
				if actual.Contains(i) && !expected.Contains(i) {
					t.Errorf("contained %d but should not", i)
				}

				if !actual.Contains(i) && expected.Contains(i) {
					t.Errorf("was expected to contain %d but did not", i)
				}
			}
		})

		t.Run("overlapping", func(t *testing.T) {
			a := constructor(1, 2, 3)
			b := constructor(2, 3, 4)

			actual := a.Difference(b)
			expected := cantor.NewHashSet[byte](1)

			for i := byte(0); i < 255; i++ {
				if actual.Contains(i) && !expected.Contains(i) {
					t.Errorf("contained %d but should not", i)
				}

				if !actual.Contains(i) && expected.Contains(i) {
					t.Errorf("was expected to contain %d but did not", i)
				}
			}
		})

		t.Run("equal", func(t *testing.T) {
			a := constructor(1, 2)
			b := constructor(1, 2)

			actual := a.Difference(b)

			for i := byte(0); i < 255; i++ {
				if actual.Contains(i) {
					t.Errorf("contained %d but should not", i)
				}
			}
		})
	})
}
