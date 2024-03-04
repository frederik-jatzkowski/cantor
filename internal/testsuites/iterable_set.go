package testsuites

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
)

// RunTestsForIterableSet runs a test suite to check correct implementation of
// [pkg/github.com/frederik-jatzkowski/cantor.IterableSet].
func RunTestsForIterableSet(t *testing.T, constructor Constructor[byte, cantor.IterableSet[byte]]) {
	t.Run("IterableSet", func(t *testing.T) {
		RunTestsForContainer(t, func(elements ...byte) cantor.Container[byte] {
			return constructor(elements...)
		})

		RunTestsForIterableSet_Elements(t, constructor)
		RunTestsForIterableSet_Size(t, constructor)
		RunTestsForIterableSet_Union(t, constructor)
		RunTestsForIterableSet_Intersection(t, constructor)
		RunTestsForIterableSet_Complement(t, constructor)

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
