package cantor_test

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
	"github.com/frederik-jatzkowski/cantor/internal/testsuites"
)

func Test_union_DerivedSet(t *testing.T) {
	testsuites.RunTestsForDerivedSet(t, func(elements ...byte) cantor.DerivedSet[byte] {
		a := cantor.NewHashSet(elements[:len(elements)/2]...)
		b := cantor.NewHashSet(elements[len(elements)/2:]...)

		return a.Union(b)
	})
}
