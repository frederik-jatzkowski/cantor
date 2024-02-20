package cantor_test

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
)

func TestHashSet(t *testing.T) {
	RunAllSetTests(
		func(elements ...int) cantor.Set[int] {
			return cantor.NewHashSet[int](elements...)
		},
		t,
	)
}
