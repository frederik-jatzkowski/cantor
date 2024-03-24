package cantor_test

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
	"github.com/frederik-jatzkowski/cantor/internal/testsuites"
)

func TestImplicitSet_Contains(t *testing.T) {
	testsuites.RunTestsForContainer(t, func(elements ...byte) cantor.Container[byte] {
		return cantor.NewImplicitSet(func(element byte) bool {
			return testsuites.SliceContains(element, elements)
		})
	})
}

func TestImplicitSet_Union(t *testing.T) {
	t.Run("both empty", func(t *testing.T) {
		a := cantor.NewImplicitSet[byte](func(element byte) bool {
			return false
		})
		b := cantor.NewImplicitSet[byte](func(element byte) bool {
			return false
		})

		actual := a.Union(b)

		for i := byte(0); i < 255; i++ {
			if actual.Contains(i) {
				t.Errorf("contained %d but should not", i)
			}
		}
	})

	t.Run("disjunct", func(t *testing.T) {
		a := cantor.NewImplicitSet[byte](func(element byte) bool {
			return element == 1 || element == 2
		})
		b := cantor.NewImplicitSet[byte](func(element byte) bool {
			return element == 3 || element == 4
		})

		actual := a.Union(b)
		expected := []byte{1, 2, 3, 4}

		for i := byte(0); i < 255; i++ {
			if !actual.Contains(i) && testsuites.SliceContains(i, expected) {
				t.Errorf("was expected to contain %d but did not", i)
			}

			if actual.Contains(i) && !testsuites.SliceContains(i, expected) {
				t.Errorf("contained %d but should not", i)
			}
		}
	})

	t.Run("overlapping", func(t *testing.T) {
		a := cantor.NewImplicitSet[byte](func(element byte) bool {
			return element == 1 || element == 2
		})
		b := cantor.NewImplicitSet[byte](func(element byte) bool {
			return element == 2 || element == 3
		})

		actual := a.Union(b)
		expected := []byte{1, 2, 3}

		for i := byte(0); i < 255; i++ {
			if !actual.Contains(i) && testsuites.SliceContains(i, expected) {
				t.Errorf("was expected to contain %d but did not", i)
			}

			if actual.Contains(i) && !testsuites.SliceContains(i, expected) {
				t.Errorf("contained %d but should not", i)
			}
		}
	})

	t.Run("equal", func(t *testing.T) {
		a := cantor.NewImplicitSet[byte](func(element byte) bool {
			return element == 1 || element == 2
		})
		b := cantor.NewImplicitSet[byte](func(element byte) bool {
			return element == 1 || element == 2
		})

		actual := a.Union(b)
		expected := []byte{1, 2}

		for i := byte(0); i < 255; i++ {
			if !actual.Contains(i) && testsuites.SliceContains(i, expected) {
				t.Errorf("was expected to contain %d but did not", i)
			}

			if actual.Contains(i) && !testsuites.SliceContains(i, expected) {
				t.Errorf("contained %d but should not", i)
			}
		}
	})
}

func TestImplicitSet_Intersect(t *testing.T) {
	t.Run("one empty", func(t *testing.T) {
		a := cantor.NewImplicitSet[byte](func(element byte) bool {
			return false
		})
		b := cantor.NewImplicitSet[byte](func(element byte) bool {
			return element == 1 || element == 2 || element == 3
		})

		actual := a.Intersect(b)

		for i := byte(0); i < 255; i++ {
			if actual.Contains(i) {
				t.Errorf("contained %d but should not", i)
			}
		}
	})

	t.Run("nonempty disjunct", func(t *testing.T) {
		a := cantor.NewImplicitSet[byte](func(element byte) bool {
			return element == 1 || element == 2
		})
		b := cantor.NewImplicitSet[byte](func(element byte) bool {
			return element == 3 || element == 4
		})

		actual := a.Intersect(b)

		for i := byte(0); i < 255; i++ {
			if actual.Contains(i) {
				t.Errorf("contained %d but should not", i)
			}
		}
	})

	t.Run("overlapping", func(t *testing.T) {
		a := cantor.NewImplicitSet[byte](func(element byte) bool {
			return element == 1 || element == 2 || element == 3
		})
		b := cantor.NewImplicitSet[byte](func(element byte) bool {
			return element == 2 || element == 3 || element == 4
		})

		actual := a.Intersect(b)
		expected := []byte{2, 3}

		for i := byte(0); i < 255; i++ {
			if !actual.Contains(i) && testsuites.SliceContains(i, expected) {
				t.Errorf("was expected to contain %d but did not", i)
			}

			if actual.Contains(i) && !testsuites.SliceContains(i, expected) {
				t.Errorf("contained %d but should not", i)
			}
		}
	})

	t.Run("equal", func(t *testing.T) {
		a := cantor.NewImplicitSet[byte](func(element byte) bool {
			return element == 1 || element == 2
		})
		b := cantor.NewImplicitSet[byte](func(element byte) bool {
			return element == 1 || element == 2
		})

		actual := a.Intersect(b)
		expected := []byte{1, 2}

		for i := byte(0); i < 255; i++ {
			if !actual.Contains(i) && testsuites.SliceContains(i, expected) {
				t.Errorf("was expected to contain %d but did not", i)
			}

			if actual.Contains(i) && !testsuites.SliceContains(i, expected) {
				t.Errorf("contained %d but should not", i)
			}
		}
	})
}

func TestImplicitSet_Complement(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		actual := cantor.NewImplicitSet[byte](func(element byte) bool {
			return false
		}).Complement()

		for i := byte(0); i < 255; i++ {
			if !actual.Contains(i) {
				t.Errorf("was expected to contain %d but did not", i)
			}
		}
	})

	t.Run("some", func(t *testing.T) {
		unexpected := []byte{1, 2, 3}
		actual := cantor.NewImplicitSet[byte](func(element byte) bool {
			return testsuites.SliceContains(element, unexpected)
		}).Complement()

		for i := byte(0); i < 255; i++ {
			if !actual.Contains(i) && !testsuites.SliceContains(i, unexpected) {
				t.Errorf("was expected to contain %d but did not", i)
			}

			if actual.Contains(i) && testsuites.SliceContains(i, unexpected) {
				t.Errorf("contained %d but should not", i)
			}
		}
	})

	t.Run("full", func(t *testing.T) {
		unexpected := testsuites.AllBytes()
		actual := cantor.NewImplicitSet[byte](func(element byte) bool {
			return testsuites.SliceContains(element, unexpected)
		}).Complement()

		for i := byte(0); i < 255; i++ {
			if !actual.Contains(i) && !testsuites.SliceContains(i, unexpected) {
				t.Errorf("was expected to contain %d but did not", i)
			}

			if actual.Contains(i) && testsuites.SliceContains(i, unexpected) {
				t.Errorf("contained %d but should not", i)
			}
		}
	})
}
