package sets

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
	"github.com/frederik-jatzkowski/cantor/internal/testsuites/testutils"
)

// RunTestsForReadableSet_Complement runs a test suite to check correct implementation of the Complement method on
// [pkg/github.com/frederik-jatzkowski/cantor.ReadableSet].
func RunTestsForReadableSet_Complement(t *testing.T, constructor Constructor[byte, cantor.ReadableSet[byte]]) {
	t.Run("Intersection", func(t *testing.T) {
		t.Run("empty", func(t *testing.T) {
			unexpected := []byte{}
			actual := constructor(unexpected...).Complement()

			for i := byte(0); i < 255; i++ {
				if !actual.Contains(i) && !testutils.SliceContains(i, unexpected) {
					t.Errorf("was expected to contain %d but did not", i)
				}

				if actual.Contains(i) && testutils.SliceContains(i, unexpected) {
					t.Errorf("contained %d but should not", i)
				}
			}
		})

		t.Run("some", func(t *testing.T) {
			unexpected := []byte{1, 2, 3}
			actual := constructor(unexpected...).Complement()

			for i := byte(0); i < 255; i++ {
				if !actual.Contains(i) && !testutils.SliceContains(i, unexpected) {
					t.Errorf("was expected to contain %d but did not", i)
				}

				if actual.Contains(i) && testutils.SliceContains(i, unexpected) {
					t.Errorf("contained %d but should not", i)
				}
			}
		})

		t.Run("full", func(t *testing.T) {
			unexpected := testutils.AllBytes()
			actual := constructor(unexpected...).Complement()

			for i := byte(0); i < 255; i++ {
				if !actual.Contains(i) && !testutils.SliceContains(i, unexpected) {
					t.Errorf("was expected to contain %d but did not", i)
				}

				if actual.Contains(i) && testutils.SliceContains(i, unexpected) {
					t.Errorf("contained %d but should not", i)
				}
			}
		})
	})
}
