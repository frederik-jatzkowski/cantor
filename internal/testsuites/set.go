package testsuites

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
)

// RunTestsForSet runs a test suite to check correct implementation of
// [pkg/github.com/frederik-jatzkowski/cantor.Set].
func RunTestsForSet(t *testing.T, constructor Constructor[byte, cantor.ReadableSet[byte]]) {
	t.Run("Set", func(t *testing.T) {
		RunTestsForReadableSet(t, func(elements ...byte) cantor.ReadableSet[byte] {
			return constructor(elements...)
		})

	})
}
