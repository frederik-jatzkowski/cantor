package cantor_test

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
	"github.com/frederik-jatzkowski/cantor/internal/testsuites"
)

func Test_complement_ImplicitSet(t *testing.T) {
	testsuites.RunTestsForImplicitSet(t, func(elements ...byte) cantor.ImplicitSet[byte] {
		return cantor.NewImplicitSet(func(element byte) bool {
			return !testsuites.SliceContains(element, elements)
		}).Complement()
	})
}
