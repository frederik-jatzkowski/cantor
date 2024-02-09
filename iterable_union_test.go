package cantor_test

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
)

func Test_union_Union(t *testing.T) {
	set1 := cantor.NewExplicitSet(1, 2).Union(cantor.NewExplicitSet(3, 4))
	set2 := cantor.NewExplicitSet(3, 4).Union(cantor.NewExplicitSet(5, 6))

	union := set1.Union(set2)

	{
		expected := 6
		actual := union.Size()
		if expected != actual {
			t.Errorf("expected size to be %d but was %d", expected, actual)
		}
	}

	cantor.NewExplicitSet(1, 2, 3, 4, 5, 6).Iter()(func(element int) (next bool) {
		if !union.Contains(element) {
			t.Errorf("union did not contain %d", element)
		}

		return true
	})
}

func Test_union_Contains(t *testing.T) {
	set := cantor.NewExplicitSet(1, 2).Union(cantor.NewExplicitSet(3, 4))

	if !set.Contains(1) {
		t.Errorf("set should contain 1")
	}

	if set.Contains(0) {
		t.Errorf("set should not contain 0")
	}
}

func Test_union_Intersection(t *testing.T) {
	set1 := cantor.NewExplicitSet(1, 2).Union(cantor.NewExplicitSet(3, 4))
	set2 := cantor.NewExplicitSet(3, 4).Union(cantor.NewExplicitSet(5, 6))

	intersection := set1.Intersect(set2)

	{
		expected := 2
		actual := intersection.Size()
		if expected != actual {
			t.Errorf("expected size to be %d but was %d", expected, actual)
		}
	}

	cantor.NewExplicitSet(3, 4).Iter()(func(element int) (next bool) {
		if !intersection.Contains(element) {
			t.Errorf("intersection did not contain %d", element)
		}

		return false
	})
}

func Test_union_Complement(t *testing.T) {
	set := cantor.NewExplicitSet(1, 2).Union(cantor.NewExplicitSet(3, 4))
	complement := set.Complement()

	set.Iter()(func(element int) (next bool) {
		if complement.Contains(element) {
			t.Errorf("both set and complement contain %d", element)
		}

		return false
	})
}

func Test_union_Iter(t *testing.T) {
	set := cantor.NewExplicitSet(1, 2).Union(cantor.NewExplicitSet(3, 4))
	counter := 0

	set.Iter()(func(element int) (next bool) {
		counter++

		return true
	})

	if counter != 4 {
		t.Errorf("invalid enumeration count: %d", counter)
	}
}

func Test_union_IterBreak(t *testing.T) {
	set := cantor.NewExplicitSet(1, 2).Union(cantor.NewExplicitSet(3, 4))
	counter := 0

	set.Iter()(func(element int) (next bool) {
		counter++

		return counter < 3
	})

	if counter != 3 {
		t.Errorf("invalid enumeration count: %d", counter)
	}
}

func Test_union_Evaluate(t *testing.T) {
	set := cantor.NewExplicitSet(1, 2).Union(cantor.NewExplicitSet(3, 4))
	evaluated := set.Evaluate()

	{
		expected := set.Size()
		actual := evaluated.Size()
		if expected != actual {
			t.Errorf("expected size to be %d but was %d", expected, actual)
		}
	}
}
