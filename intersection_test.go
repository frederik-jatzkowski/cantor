package cantor_test

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
	"github.com/frederik-jatzkowski/cantor/internal/testsuites"
)

func Test_intersection_DerivedSet(t *testing.T) {
	testsuites.RunTestsForDerivedSet(t, func(elements ...byte) cantor.DerivedSet[byte] {
		a := cantor.NewHashSet(elements...)
		b := cantor.NewHashSet(elements...)

		return a.Intersect(b)
	})
}
