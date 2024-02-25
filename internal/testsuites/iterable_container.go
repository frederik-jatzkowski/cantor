package testsuites

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
)

// RunTestsForDeduplicatingIterableContainer runs a test suite to check correct implementation of
// [pkg/github.com/frederik-jatzkowski/cantor.DeduplicatingIterableContainer].
func RunTestsForDeduplicatingIterableContainer(t *testing.T, constructor Constructor[byte, cantor.DeduplicatingIterableContainer[byte]]) {
	t.Run("Container", func(t *testing.T) {
		RunTestsForContainer(t, func(elements ...byte) cantor.Container[byte] {
			return constructor(elements...)
		})
	})

	t.Run("iterator on empty set", func(t *testing.T) {
		constructor().UniqueElements()(func(element byte) bool {
			t.Errorf("unexpected element during iteration: %d", element)

			return true
		})
	})

	t.Run("iterator on universe", func(t *testing.T) {
		expected := allBytes()
		container := constructor(expected...)

		found := cantor.NewHashSet[byte]()

		container.UniqueElements()(func(element byte) bool {
			if found.Contains(element) {
				t.Errorf("duplicate element: %d", element)
			}

			if !SliceContains(element, expected) {
				t.Errorf("unexpected element: %d", element)
			}

			found.Add(element)

			return true
		})

		if found.Size() != len(expected) {
			t.Errorf("only found %d elements but expected %d", found.Size(), len(expected))
		}
	})

	t.Run("iterator on some elements", func(t *testing.T) {
		expected := []byte{1, 2, 5, 8, 123}
		container := constructor(expected...)

		found := cantor.NewHashSet[byte]()

		container.UniqueElements()(func(element byte) bool {
			if found.Contains(element) {
				t.Errorf("duplicate element: %d", element)
			}

			if !SliceContains(element, expected) {
				t.Errorf("unexpected element: %d", element)
			}

			found.Add(element)

			return true
		})

		if found.Size() != len(expected) {
			t.Errorf("only found %d elements but expected %d", found.Size(), len(expected))
		}
	})

	t.Run("iterator with break", func(t *testing.T) {
		container := constructor(allBytes()...)

		sum := 0
		limit := 100
		container.UniqueElements()(func(element byte) bool {
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

	t.Run("no duplicate elements", func(t *testing.T) {
		expected := []byte{0, 0, 0, 1, 1, 1, 2, 2, 2, 5, 8, 8, 8, 123, 255, 255}
		container := constructor(expected...)

		found := cantor.NewHashSet[byte]()
		container.UniqueElements()(func(element byte) bool {
			if found.Contains(element) {
				t.Errorf("duplicate element: %d", element)
			}

			found.Add(element)

			return true
		})
	})
}
