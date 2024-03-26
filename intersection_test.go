package cantor_test

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
	"github.com/frederik-jatzkowski/cantor/internal/testsuites/sets"
)

func Test_intersection_ReadableSet(t *testing.T) {
	sets.RunTestsForReadableSet(t, func(elements ...byte) cantor.ReadableSet[byte] {
		a := cantor.NewHashSet(elements...)
		b := cantor.NewHashSet(elements...)

		return a.Intersect(b)
	})
}
