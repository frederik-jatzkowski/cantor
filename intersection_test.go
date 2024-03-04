package cantor_test

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
	"github.com/frederik-jatzkowski/cantor/internal/testsuites"
)

func Test_intersection_IterableSet(t *testing.T) {
	testsuites.RunTestsForIterableSet(t, func(elements ...byte) cantor.IterableSet[byte] {
		a := cantor.NewHashSet(elements...)
		b := cantor.NewHashSet(elements...)

		return a.Intersect(b)
	})
}
