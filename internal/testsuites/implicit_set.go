package testsuites

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
)

// RunTestsForImplicitSet runs a test suite to check correct implementation of
// [pkg/github.com/frederik-jatzkowski/cantor.ImplicitSet].
func RunTestsForImplicitSet(t *testing.T, constructor Constructor[byte, cantor.ImplicitSet[byte]]) {
	t.Run("Container", func(t *testing.T) {
		RunTestsForContainer(t, func(elements ...byte) cantor.Container[byte] {
			return constructor(elements...)
		})
	})

	RunTestsForImplicitSet_Union(t, constructor)
	RunTestsForImplicitSet_Intersection(t, constructor)
	RunTestsForImplicitSet_Complement(t, constructor)
}
