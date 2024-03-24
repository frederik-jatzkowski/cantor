package testsuites

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
)

// RunTestsForReadableSet runs a test suite to check correct implementation of
// [pkg/github.com/frederik-jatzkowski/cantor.ReadableSet].
func RunTestsForReadableSet(t *testing.T, constructor Constructor[byte, cantor.ReadableSet[byte]]) {
	t.Run("ReadableSet", func(t *testing.T) {
		RunTestsForContainer(t, func(elements ...byte) cantor.Container[byte] {
			return constructor(elements...)
		})

		RunTestsForReadableSet_Elements(t, constructor)
		RunTestsForReadableSet_Size(t, constructor)
		RunTestsForReadableSet_Union(t, constructor)
		RunTestsForReadableSet_Intersection(t, constructor)
		RunTestsForReadableSet_Complement(t, constructor)

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
