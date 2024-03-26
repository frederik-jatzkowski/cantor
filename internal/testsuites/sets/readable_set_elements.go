package sets

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
	"github.com/frederik-jatzkowski/cantor/internal/testsuites/testutils"
)

// RunTestsForReadableSet_Elements runs a test suite to check correct implementation of the Elements method on
// [pkg/github.com/frederik-jatzkowski/cantor.ReadableSet].
func RunTestsForReadableSet_Elements(t *testing.T, constructor Constructor[byte, cantor.ReadableSet[byte]]) {
	t.Run("Elements", func(t *testing.T) {
		t.Run("empty set", func(t *testing.T) {
			constructor().Elements()(func(element byte) bool {
				t.Errorf("unexpected element during iteration: %d", element)

				return true
			})
		})

		t.Run("universe", func(t *testing.T) {
			expected := testutils.AllBytes()
			container := constructor(expected...)

			found := cantor.NewHashSet[byte]()

			container.Elements()(func(element byte) bool {
				if found.Contains(element) {
					t.Errorf("duplicate element: %d", element)
				}

				if !testutils.SliceContains(element, expected) {
					t.Errorf("unexpected element: %d", element)
				}

				found.Add(element)

				return true
			})

			if found.Size() != len(expected) {
				t.Errorf("only found %d elements but expected %d", found.Size(), len(expected))
			}
		})

		t.Run("some elements", func(t *testing.T) {
			expected := []byte{1, 2, 5, 8, 123}
			container := constructor(expected...)

			found := cantor.NewHashSet[byte]()

			container.Elements()(func(element byte) bool {
				if found.Contains(element) {
					t.Errorf("duplicate element: %d", element)
				}

				if !testutils.SliceContains(element, expected) {
					t.Errorf("unexpected element: %d", element)
				}

				found.Add(element)

				return true
			})

			if found.Size() != len(expected) {
				t.Errorf("only found %d elements but expected %d", found.Size(), len(expected))
			}
		})

		t.Run("iteration with break", func(t *testing.T) {
			container := constructor(testutils.AllBytes()...)

			sum := 0
			limit := 100
			container.Elements()(func(element byte) bool {
				if sum > limit {
					t.Errorf("element yielded after break: %d", element)
				}

				if sum+int(element) > limit {
					return false
				}

				sum += int(element)

				return true
			})

			if sum > limit {
				t.Errorf("iteration did not stop when the yield function returned false")
			}
		})

		t.Run("deduplication", func(t *testing.T) {
			expected := []byte{0, 0, 0, 1, 1, 1, 2, 2, 2, 5, 8, 8, 8, 123, 255, 255}
			container := constructor(expected...)

			found := cantor.NewHashSet[byte]()
			container.Elements()(func(element byte) bool {
				if found.Contains(element) {
					t.Errorf("duplicate element: %d", element)
				}

				found.Add(element)

				return true
			})
		})
	})
}
