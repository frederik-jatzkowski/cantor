package cantor_test

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
	"github.com/frederik-jatzkowski/cantor/internal/testsuites"
)

func Test_union_IterableSet(t *testing.T) {
	testsuites.RunTestsForIterableSet(t, func(elements ...byte) cantor.IterableSet[byte] {
		a := cantor.NewHashSet(elements[:len(elements)/2]...)
		b := cantor.NewHashSet(elements[len(elements)/2:]...)

		return a.Union(b)
	})
}
