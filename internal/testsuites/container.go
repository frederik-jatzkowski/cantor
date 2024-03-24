package testsuites

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
)

// RunTestsForContainer runs a test suite to check correct implementation of
// [pkg/github.com/frederik-jatzkowski/cantor.Container].
func RunTestsForContainer(t *testing.T, constructor Constructor[byte, cantor.Container[byte]]) {
	t.Run("Container", func(t *testing.T) {
		t.Run("Contains", func(t *testing.T) {
			t.Run("empty", func(t *testing.T) {
				actual := constructor()

				for i := byte(0); i < 255; i++ {
					if actual.Contains(i) {
						t.Errorf("contained %d but should not", i)
					}
				}
			})

			t.Run("some", func(t *testing.T) {
				expected := []byte{2, 5, 8, 123}
				actual := constructor(expected...)

				for i := byte(0); i < 255; i++ {
					if SliceContains(i, expected) && !actual.Contains(i) {
						t.Errorf("was expected to contain %d but did not", i)
					}

					if !SliceContains(i, expected) && actual.Contains(i) {
						t.Errorf("contained %d but should not", i)
					}
				}
			})

			t.Run("full", func(t *testing.T) {
				expected := AllBytes()
				actual := constructor(expected...)

				for _, i := range expected {
					if !actual.Contains(i) {
						t.Errorf("did not contain %d", i)
					}
				}
			})
		})
	})
}
