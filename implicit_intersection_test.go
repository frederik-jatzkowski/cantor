package cantor_test

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
	"github.com/frederik-jatzkowski/cantor/internal/testsuites"
)

func Test_implicitIntersection_ImplicitSet(t *testing.T) {
	testsuites.RunTestsForImplicitSet(t, func(elements ...byte) cantor.ImplicitSet[byte] {
		a := cantor.NewImplicitSet(func(element byte) bool {
			return testsuites.SliceContains(element, elements)
		})

		b := cantor.NewImplicitSet(func(element byte) bool {
			return testsuites.SliceContains(element, elements)
		})

		return a.Intersect(b)
	})
}
