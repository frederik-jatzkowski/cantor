package cantor_test

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
)

func Test_intersection_Union(t *testing.T) {
	set1 := cantor.NewExplicitSet(1, 3).Intersect(cantor.NewExplicitSet(2, 3))
	set2 := cantor.NewExplicitSet(4, 5).Intersect(cantor.NewExplicitSet(4, 6))

	union := set1.Union(set2)

	{
		expected := 2
		actual := union.Size()
		if expected != actual {
			t.Errorf("expected size to be %d but was %d", expected, actual)
		}
	}

	cantor.NewExplicitSet(3, 4).Enumerate(func(element int) (stop bool) {
		if !union.Contains(element) {
			t.Errorf("union did not contain %d", element)
		}

		return false
	})
}

func Test_intersection_Contains(t *testing.T) {
	set := cantor.NewExplicitSet(1, 2).Intersect(cantor.NewExplicitSet(2, 3))

	if !set.Contains(2) {
		t.Errorf("set should contain 2")
	}

	if set.Contains(1) {
		t.Errorf("set should not contain 1")
	}
}

func Test_intersection_Intersection(t *testing.T) {
	set1 := cantor.NewExplicitSet(1, 2).Intersect(cantor.NewExplicitSet(2, 3))
	set2 := cantor.NewExplicitSet(2, 5).Intersect(cantor.NewExplicitSet(2, 6))

	intersection := set1.Intersect(set2)

	{
		expected := 1
		actual := intersection.Size()
		if expected != actual {
			t.Errorf("expected size to be %d but was %d", expected, actual)
		}
	}

	cantor.NewExplicitSet(2).Enumerate(func(element int) (stop bool) {
		if !intersection.Contains(element) {
			t.Errorf("intersection did not contain %d", element)
		}

		return false
	})
}

func Test_intersection_Complement(t *testing.T) {
	set := cantor.NewExplicitSet(1, 2).Intersect(cantor.NewExplicitSet(2, 3))
	complement := set.Complement()

	set.Enumerate(func(element int) (stop bool) {
		if complement.Contains(element) {
			t.Errorf("both set and complement contain %d", element)
		}

		return false
	})
}

func Test_intersection_Enumerate(t *testing.T) {
	set := cantor.NewExplicitSet(1, 2, 3).Intersect(cantor.NewExplicitSet(2, 3, 4))
	counter := 0

	set.Enumerate(func(element int) (stop bool) {
		counter++

		return false
	})

	if counter != 2 {
		t.Errorf("invalid enumeration count: %d", counter)
	}
}

func Test_intersection_EnumerateStop(t *testing.T) {
	set := cantor.NewExplicitSet(1, 2, 3).Intersect(cantor.NewExplicitSet(2, 3, 4))
	counter := 0

	set.Enumerate(func(element int) (stop bool) {
		counter++

		return counter >= 1
	})

	if counter != 1 {
		t.Errorf("invalid enumeration count: %d", counter)
	}
}

func Test_intersection_Evaluate(t *testing.T) {
	set := cantor.NewExplicitSet(1, 2).Intersect(cantor.NewExplicitSet(3, 4))
	evaluated := set.Evaluate()

	{
		expected := set.Size()
		actual := evaluated.Size()
		if expected != actual {
			t.Errorf("expected size to be %d but was %d", expected, actual)
		}
	}
}
