package testsuites

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
)

// RunTestsForDerivedSet runs a test suite to check correct implementation of
// [pkg/github.com/frederik-jatzkowski/cantor.DerivedSet].
func RunTestsForDerivedSet(t *testing.T, constructor Constructor[byte, cantor.DerivedSet[byte]]) {
	t.Run("IterableContainer", func(t *testing.T) {
		RunTestsForIterableContainer(t, func(elements ...byte) cantor.IterableContainer[byte] {
			return constructor(elements...)
		})
	})

	RunTestsForDerivedSet_Union(t, constructor)
	RunTestsForDerivedSet_Intersection(t, constructor)
	RunTestsForDerivedSet_Complement(t, constructor)

	t.Run("IntoHashSet", func(t *testing.T) {
		t.Run("empty", func(t *testing.T) {
			actual := constructor().IntoHashSet()

			for i := byte(0); i < 255; i++ {
				if actual.Contains(i) {
					t.Errorf("contained %d but should not", i)
				}
			}
		})

		t.Run("some", func(t *testing.T) {
			expected := []byte{2, 5, 8, 123}
			actual := constructor(expected...).IntoHashSet()

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
			expected := allBytes()
			actual := constructor(expected...).IntoHashSet()

			for _, i := range expected {
				if !actual.Contains(i) {
					t.Errorf("did not contain %d", i)
				}
			}
		})

		t.Run("base set is independent", func(t *testing.T) {
			expected := []byte{2, 5, 8, 123}
			base := constructor(expected...)
			actual := base.IntoHashSet()
			actual.Add(1)

			if base.Contains(1) {
				t.Error("base set is not independent from evaluated set")
			}
		})
	})

	t.Run("fmt.Stringer", func(t *testing.T) {
		set := constructor(1, 2)
		str := set.String()

		switch str {
		case "{1, 2}", "{2, 1}":
		default:
			t.Errorf("invalid string: %s", str)
		}
	})
}
