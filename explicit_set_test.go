package cantor_test

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
)

func TestNewExplicitSet(t *testing.T) {
	set := cantor.NewExplicitSet(1, 2, 3)

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
}

func TestExplicitSet_Contains(t *testing.T) {
	set := cantor.NewExplicitSet(1, 2)

	if !set.Contains(2) {
		t.Errorf("set should contain 2")
	}

	if set.Contains(0) {
		t.Errorf("set should not contain 0")
	}
}

func TestExplicitSet_Add(t *testing.T) {
	set := cantor.NewExplicitSet(1, 2, 3)

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
}

func TestExplicitSet_Remove(t *testing.T) {
	set := cantor.NewExplicitSet(1, 2, 3)

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
}

func TestExplicitSet_Union(t *testing.T) {
	set1 := cantor.NewExplicitSet(1, 2, 3)
	set2 := cantor.NewExplicitSet(3, 4, 5)

	union := set1.Union(set2)

	{
		expected := 5
		actual := union.Size()
		if expected != actual {
			t.Errorf("expected size to be %d but was %d", expected, actual)
		}
	}

	cantor.NewExplicitSet(1, 2, 3, 4, 5).Iter()(func(element int) (next bool) {
		if !union.Contains(element) {
			t.Errorf("union did not contain %d", element)
		}

		return false
	})
}

func TestExplicitSet_Intersection(t *testing.T) {
	set1 := cantor.NewExplicitSet(1, 2, 3)
	set2 := cantor.NewExplicitSet(2, 3, 4)

	intersection := set1.Intersect(set2)

	{
		expected := 2
		actual := intersection.Size()
		if expected != actual {
			t.Errorf("expected size to be %d but was %d", expected, actual)
		}
	}

	cantor.NewExplicitSet(2, 3).Iter()(func(element int) (next bool) {
		if !intersection.Contains(element) {
			t.Errorf("intersection did not contain %d", element)
		}

		return false
	})
}

func TestExplicitSet_Complement(t *testing.T) {
	set := cantor.NewExplicitSet(1, 2, 3, 4, 5)
	complement := set.Complement()

	set.Iter()(func(element int) (next bool) {
		if complement.Contains(element) {
			t.Errorf("both set and complement contain %d", element)
		}

		return false
	})
}

func TestExplicitSet_Iter(t *testing.T) {
	counter := 0

	cantor.NewExplicitSet(1, 2, 3, 4, 5).Iter()(func(element int) (next bool) {
		counter++

		return true
	})

	if counter != 5 {
		t.Errorf("invalid enumeration count: %d", counter)
	}
}

func TestExplicitSet_IterBreak(t *testing.T) {
	counter := 0

	cantor.NewExplicitSet(1, 2, 3, 4, 5).Iter()(func(element int) (next bool) {
		counter++

		return counter < 3
	})

	if counter != 3 {
		t.Errorf("invalid enumeration count: %d", counter)
	}
}

func TestExplicitSet_Evaluate(t *testing.T) {
	set1 := cantor.NewExplicitSet(1, 2, 3)
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
}

func TestExplicitSet_String(t *testing.T) {
	set := cantor.NewExplicitSet(1, 2)
	str := set.String()

	switch str {
	case "{1, 2}", "{2, 1}":
	default:
		t.Errorf("invalid string: %s", str)
	}
}
