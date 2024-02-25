package testsuites

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
)

// RunTestsForDerivedSet runs a test suite to check correct implementation of
// [pkg/github.com/frederik-jatzkowski/cantor.DerivedSet].
func RunTestsForDerivedSet(t *testing.T, constructor Constructor[byte, cantor.DerivedSet[byte]]) {
	t.Run("DerivedSet", func(t *testing.T) {
		RunTestsForDeduplicatingIterableContainer(t, func(elements ...byte) cantor.DeduplicatingIterableContainer[byte] {
			return constructor(elements...)
		})

		RunTestsForDerivedSet_Union(t, constructor)
		RunTestsForDerivedSet_Intersection(t, constructor)
		RunTestsForDerivedSet_Complement(t, constructor)

		t.Run("fmt.Stringer", func(t *testing.T) {
			set := constructor(1, 2)
			str := set.String()

			switch str {
			case "{1, 2}", "{2, 1}":
			default:
				t.Errorf("invalid string: %s", str)
			}
		})
	})
}
