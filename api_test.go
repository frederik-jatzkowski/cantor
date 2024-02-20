package cantor_test

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
)

// RunAllSetTests runs a suite of tests against an implementation of [Set]. For this, a constructor has to be given.
func RunAllSetTests(constructor func(elements ...int) cantor.Set[int], t *testing.T) {
	t.Run("New", func(t *testing.T) {
		set := constructor(1, 2, 3)

		{
			expected := 3
			actual := set.Size()
			if expected != actual {
				t.Errorf("expected size to be %d but was %d", expected, actual)
			}
		}

		if !set.Contains(1) {
			t.Errorf("set does not contain %d", 1)
		}

		if !set.Contains(2) {
			t.Errorf("set does not contain %d", 2)
		}

		if !set.Contains(3) {
			t.Errorf("set does not contain %d", 3)
		}
	})

	t.Run("Contains", func(t *testing.T) {
		set := constructor(1, 2)

		if !set.Contains(2) {
			t.Errorf("set should contain 2")
		}

		if set.Contains(0) {
			t.Errorf("set should not contain 0")
		}
	})

	t.Run("Add", func(t *testing.T) {
		set := constructor(1, 2, 3)

		if set.Contains(4) {
			t.Error("Contains(4) should return false")
		}

		if !set.Add(4) {
			t.Error("Add(4) should return true")
		}

		if !set.Contains(4) {
			t.Error("Contains(4) should return true now")
		}

		if set.Add(4) {
			t.Error("Add(4) should return false now")
		}
	})

	t.Run("Remove", func(t *testing.T) {
		set := constructor(1, 2, 3)

		if !set.Contains(3) {
			t.Error("Contains(3) should return true")
		}

		if !set.Remove(3) {
			t.Error("Remove(3) should return true")
		}

		if set.Contains(3) {
			t.Error("Contains(3) should return false now")
		}

		if set.Remove(3) {
			t.Error("Remove(3) should return false now")
		}
	})

	t.Run("Union", func(t *testing.T) {
		set1 := constructor(1, 2, 3)
		set2 := constructor(3, 4, 5)

		union := set1.Union(set2)

		{
			expected := 5
			actual := union.Size()
			if expected != actual {
				t.Errorf("expected size to be %d but was %d", expected, actual)
			}
		}

		for _, element := range []int{1, 2, 3, 4, 5} {
			if !union.Contains(element) {
				t.Errorf("union did not contain %d", element)
			}
		}
	})

	t.Run("Intersection", func(t *testing.T) {
		set1 := constructor(1, 2, 3)
		set2 := constructor(2, 3, 4)

		intersection := set1.Intersect(set2)

		{
			expected := 2
			actual := intersection.Size()
			if expected != actual {
				t.Errorf("expected size to be %d but was %d", expected, actual)
			}
		}

		for _, element := range []int{2, 3} {
			if !intersection.Contains(element) {
				t.Errorf("intersection did not contain %d", element)
			}
		}
	})

	t.Run("Complement", func(t *testing.T) {
		set := constructor(1, 2, 3, 4, 5)
		complement := set.Complement()

		for _, element := range []int{1, 2, 3, 4, 5} {
			if complement.Contains(element) {
				t.Errorf("both set and complement contain %d", element)
			}
		}

		for _, element := range []int{0, 6, 1000} {
			if !complement.Contains(element) {
				t.Errorf("complement should contain %d", element)
			}
		}
	})

	t.Run("Iter", func(t *testing.T) {
		reference := cantor.NewHashSet(1, 2, 3, 4, 5)
		found := cantor.NewHashSet[int]()

		constructor(1, 2, 3, 4, 5).IterateDistinct()(func(element int) (next bool) {
			if !found.Add(element) {
				t.Errorf("duplicate element: %d", element)
			}

			if !reference.Contains(element) {
				t.Errorf("unexpected element: %d", element)
			}

			return true
		})

		if found.Size() != reference.Size() {
			t.Errorf("invalid number of elements: %d", found.Size())
		}
	})

	t.Run("Iter with break", func(t *testing.T) {
		counter := 0

		cantor.NewHashSet(1, 2, 3, 4, 5).IterateDistinct()(func(element int) (next bool) {
			counter++

			return counter < 3
		})

		if counter != 3 {
			t.Errorf("invalid enumeration count: %d", counter)
		}
	})

	t.Run("Evaluate", func(t *testing.T) {
		set1 := constructor(1, 2, 3)
		set2 := set1.Evaluate()

		{
			expected := set1.Size()
			actual := set2.Size()
			if expected != actual {
				t.Errorf("expected size to be %d but was %d", expected, actual)
			}
		}

		set2.Add(4)

		if set1.Contains(4) {
			t.Errorf("evaluate should create new independent sets")
		}
	})

	t.Run("String", func(t *testing.T) {
		set := constructor(1, 2)
		str := set.String()

		switch str {
		case "{1, 2}", "{2, 1}":
		default:
			t.Errorf("invalid string: %s", str)
		}
	})
}
