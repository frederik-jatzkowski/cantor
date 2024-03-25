package sets

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
	"github.com/frederik-jatzkowski/cantor/internal/testsuites/testutils"
)

// RunTestsForSet runs a test suite to check correct implementation of
// [pkg/github.com/frederik-jatzkowski/cantor.Set].
func RunTestsForSet(t *testing.T, constructor Constructor[byte, cantor.Set[byte]]) {
	t.Run("ReadableSet", func(t *testing.T) {
		RunTestsForReadableSet(t, func(elements ...byte) cantor.ReadableSet[byte] {
			return constructor(elements...)
		})

		t.Run("Add", func(t *testing.T) {
			t.Run("empty", func(t *testing.T) {
				set := constructor()

				for i := byte(0); i < 255; i++ {
					if !set.Add(i) {
						t.Errorf("set should have changed when adding %d", i)
					}

					if !set.Contains(i) {
						t.Errorf("set did not contain %d after adding it", i)
					}
				}
			})

			t.Run("some", func(t *testing.T) {
				initial := []byte{2, 5, 8, 123}
				set := constructor(initial...)

				for i := byte(0); i < 255; i++ {
					changed := set.Add(i)
					if changed && testutils.SliceContains(i, initial) {
						t.Errorf("set should not have changed when adding %d", i)
					}

					if !changed && !testutils.SliceContains(i, initial) {
						t.Errorf("set should have changed when adding %d", i)
					}

					if !set.Contains(i) {
						t.Errorf("set did not contain %d after adding it", i)
					}
				}
			})

			t.Run("full", func(t *testing.T) {
				set := constructor(testutils.AllBytes()...)

				for i := byte(0); i < 255; i++ {
					changed := set.Add(i)
					if changed {
						t.Errorf("set should not have changed when adding %d", i)
					}
				}
			})

			t.Run("repetitive", func(t *testing.T) {
				set := constructor()

				if !set.Add(1) {
					t.Errorf("set should have changed when adding %d", 1)
				}

				if set.Add(1) {
					t.Errorf("set should not have changed when adding %d again", 1)
				}
			})
		})

		t.Run("Remove", func(t *testing.T) {
			t.Run("empty", func(t *testing.T) {
				set := constructor()

				for i := byte(0); i < 255; i++ {
					if set.Remove(i) {
						t.Errorf("set should not have changed when removing %d", i)
					}
				}
			})

			t.Run("some", func(t *testing.T) {
				initial := []byte{2, 5, 8, 123}
				set := constructor(initial...)

				for i := byte(0); i < 255; i++ {
					changed := set.Remove(i)
					if !changed && testutils.SliceContains(i, initial) {
						t.Errorf("set should have changed when removing %d", i)
					}

					if changed && !testutils.SliceContains(i, initial) {
						t.Errorf("set should not have changed when removing %d", i)
					}

					if set.Contains(i) {
						t.Errorf("set did contain %d after removing it", i)
					}
				}
			})

			t.Run("full", func(t *testing.T) {
				set := constructor(testutils.AllBytes()...)

				for i := byte(0); i < 255; i++ {
					changed := set.Remove(i)
					if !changed {
						t.Errorf("set should have changed when removing %d", i)
					}

					if set.Contains(i) {
						t.Errorf("set did contain %d after removing it", i)
					}
				}
			})

			t.Run("repetitive", func(t *testing.T) {
				set := constructor(1)

				if !set.Remove(1) {
					t.Errorf("set should have changed when removing %d", 1)
				}

				if set.Remove(1) {
					t.Errorf("set should not have changed when removing %d again", 1)
				}
			})
		})
	})
}
