package cantor_test

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
	"github.com/frederik-jatzkowski/cantor/internal/testsuites/sets"
)

func Test_union_ReadableSet(t *testing.T) {
	sets.RunTestsForReadableSet(t, func(elements ...byte) cantor.ReadableSet[byte] {
		a := cantor.NewHashSet(elements[:len(elements)/2]...)
		b := cantor.NewHashSet(elements[len(elements)/2:]...)

		return a.Union(b)
	})
}
