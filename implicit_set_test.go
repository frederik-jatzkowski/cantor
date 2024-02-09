package cantor_test

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
)

func Test_implicitSet_Union(t *testing.T) {
	set1 := cantor.NewImplicitSet(func(element int) bool {
		return element%2 == 0
	})
	set2 := cantor.NewImplicitSet(func(element int) bool {
		return element%3 == 0
	})

	union := set1.Union(set2)

	cantor.NewExplicitSet(2, 3, 4, 6).Iter()(func(element int) (next bool) {
		if !union.Contains(element) {
			t.Errorf("union did not contain %d", element)
		}

		return false
	})
}

func Test_implicitSet_Contains(t *testing.T) {
	set := cantor.NewImplicitSet(func(element int) bool {
		return element%2 == 0
	})

	if !set.Contains(2) {
		t.Errorf("set should contain 2")
	}

	if set.Contains(1) {
		t.Errorf("set should not contain 1")
	}

	if !set.Contains(-2) {
		t.Errorf("set should contain -2")
	}

	if set.Contains(-1) {
		t.Errorf("set should not contain -1")
	}
}

func Test_implicitSet_Intersection(t *testing.T) {
	set1 := cantor.NewImplicitSet(func(element int) bool {
		return element%2 == 0
	})
	set2 := cantor.NewImplicitSet(func(element int) bool {
		return element%3 == 0
	})

	intersection := set1.Intersect(set2)

	if !intersection.Contains(6) {
		t.Errorf("set should contain 6")
	}

	if intersection.Contains(2) {
		t.Errorf("set should not contain 2")
	}
}

func Test_implicitSet_Complement(t *testing.T) {
	set := cantor.NewImplicitSet(func(element int) bool {
		return element%2 == 0
	})

	complement := set.Complement()

	if !complement.Contains(1) {
		t.Errorf("set should contain 1")
	}

	if complement.Contains(2) {
		t.Errorf("set should not contain 2")
	}
}
