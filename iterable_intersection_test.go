package cantor_test

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
)

func Test_intersection_Union(t *testing.T) {
	set1 := cantor.NewHashSet(1, 3).Intersect(cantor.NewHashSet(2, 3))
	set2 := cantor.NewHashSet(4, 5).Intersect(cantor.NewHashSet(4, 6))

	union := set1.Union(set2)

	{
		expected := 2
		actual := union.Size()
		if expected != actual {
			t.Errorf("expected size to be %d but was %d", expected, actual)
		}
	}

	cantor.NewHashSet(3, 4).Iter()(func(element int) (next bool) {
		if !union.Contains(element) {
			t.Errorf("union did not contain %d", element)
		}

		return false
	})
}

func Test_intersection_Contains(t *testing.T) {
	set := cantor.NewHashSet(1, 2).Intersect(cantor.NewHashSet(2, 3))

	if !set.Contains(2) {
		t.Errorf("set should contain 2")
	}

	if set.Contains(1) {
		t.Errorf("set should not contain 1")
	}
}

func Test_intersection_Intersection(t *testing.T) {
	set1 := cantor.NewHashSet(1, 2).Intersect(cantor.NewHashSet(2, 3))
	set2 := cantor.NewHashSet(2, 5).Intersect(cantor.NewHashSet(2, 6))

	intersection := set1.Intersect(set2)

	{
		expected := 1
		actual := intersection.Size()
		if expected != actual {
			t.Errorf("expected size to be %d but was %d", expected, actual)
		}
	}

	cantor.NewHashSet(2).Iter()(func(element int) (next bool) {
		if !intersection.Contains(element) {
			t.Errorf("intersection did not contain %d", element)
		}

		return false
	})
}

func Test_intersection_Complement(t *testing.T) {
	set := cantor.NewHashSet(1, 2).Intersect(cantor.NewHashSet(2, 3))
	complement := set.Complement()

	set.Iter()(func(element int) (next bool) {
		if complement.Contains(element) {
			t.Errorf("both set and complement contain %d", element)
		}

		return false
	})
}

func Test_intersection_Iter(t *testing.T) {
	set := cantor.NewHashSet(1, 2, 3).Intersect(cantor.NewHashSet(2, 3, 4))
	counter := 0

	set.Iter()(func(element int) (next bool) {
		counter++

		return true
	})

	if counter != 2 {
		t.Errorf("invalid enumeration count: %d", counter)
	}
}

func Test_intersection_IterBreak(t *testing.T) {
	set := cantor.NewHashSet(1, 2, 3).Intersect(cantor.NewHashSet(2, 3, 4))
	counter := 0

	set.Iter()(func(element int) (next bool) {
		counter++

		return counter < 1
	})

	if counter != 1 {
		t.Errorf("invalid enumeration count: %d", counter)
	}
}

func Test_intersection_Evaluate(t *testing.T) {
	set := cantor.NewHashSet(1, 2).Intersect(cantor.NewHashSet(3, 4))
	evaluated := set.Evaluate()

	{
		expected := set.Size()
		actual := evaluated.Size()
		if expected != actual {
			t.Errorf("expected size to be %d but was %d", expected, actual)
		}
	}
}

func Test_intersection_String(t *testing.T) {
	set := cantor.NewHashSet(0, 1, 2).Intersect(cantor.NewHashSet(1, 2, 3))
	str := set.String()

	switch str {
	case "{1, 2}", "{2, 1}":
	default:
		t.Errorf("invalid string: %s", str)
	}
}
