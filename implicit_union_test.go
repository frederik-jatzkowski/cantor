package cantor_test

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
	"github.com/frederik-jatzkowski/cantor/internal/testsuites"
)

func Test_implicitUnion_ImplicitSet(t *testing.T) {
	testsuites.RunTestsForImplicitSet(t, func(elements ...byte) cantor.ImplicitSet[byte] {
		a := cantor.NewImplicitSet(func(element byte) bool {
			return testsuites.SliceContains(element, elements[:len(elements)/2])
		})

		b := cantor.NewImplicitSet(func(element byte) bool {
			return testsuites.SliceContains(element, elements[len(elements)/2:])
		})

		return a.Union(b)
	})
}
