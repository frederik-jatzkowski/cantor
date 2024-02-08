package cantor_test

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
)

func Test_implicitIntersection_Union(t *testing.T) {
	set1 := cantor.NewImplicitSet(func(element int) bool {
		return element%2 == 0
	})
	set2 := cantor.NewImplicitSet(func(element int) bool {
		return element%3 == 0
	})
	set3 := cantor.NewImplicitSet(func(element int) bool {
		return element%5 == 0
	})
	set4 := cantor.NewImplicitSet(func(element int) bool {
		return element%7 == 0
	})

	intersection12 := set1.Intersect(set2)
	intersection34 := set3.Intersect(set4)

	union := intersection12.Union(intersection34)

	if !union.Contains(2 * 3) {
		t.Errorf("set should contain %d", 2*3)
	}

	if !union.Contains(5 * 7) {
		t.Errorf("set should contain %d", 2*3)
	}

	if union.Contains(2 * 5) {
		t.Errorf("set should not contain %d", 2*5)
	}
}

func Test_implicitIntersection_Contains(t *testing.T) {
	set1 := cantor.NewImplicitSet(func(element int) bool {
		return element%2 == 0
	})
	set2 := cantor.NewImplicitSet(func(element int) bool {
		return element%3 == 0
	})

	intersection := set1.Intersect(set2)

	if !intersection.Contains(2 * 3) {
		t.Errorf("set should contain %d", 2*3)
	}

	if intersection.Contains(2) {
		t.Errorf("set should not contain 2")
	}

	if intersection.Contains(3) {
		t.Errorf("set should not contain 3")
	}
}

func Test_implicitIntersection_Intersection(t *testing.T) {
	set1 := cantor.NewImplicitSet(func(element int) bool {
		return element%2 == 0
	})
	set2 := cantor.NewImplicitSet(func(element int) bool {
		return element%3 == 0
	})
	set3 := cantor.NewImplicitSet(func(element int) bool {
		return element%5 == 0
	})
	set4 := cantor.NewImplicitSet(func(element int) bool {
		return element%7 == 0
	})

	intersection12 := set1.Intersect(set2)
	intersection34 := set3.Intersect(set4)

	intersection := intersection12.Intersect(intersection34)

	if !intersection.Contains(2 * 3 * 5 * 7) {
		t.Errorf("set should contain %d", 2*3*5*7)
	}

	if intersection.Contains(2 * 3) {
		t.Errorf("set should not contain %d", 2*3)
	}

	if intersection.Contains(5 * 7) {
		t.Errorf("set should not contain %d", 5*7)
	}
}

func Test_implicitIntersection_Complement(t *testing.T) {
	set1 := cantor.NewImplicitSet(func(element int) bool {
		return element%2 == 0
	})
	set2 := cantor.NewImplicitSet(func(element int) bool {
		return element%3 == 0
	})

	intersection := set1.Intersect(set2)
	complement := intersection.Complement()

	if !complement.Contains(7) {
		t.Errorf("set should contain 7")
	}

	if complement.Contains(2 * 3) {
		t.Errorf("set should not contain %d", 2*3)
	}
}
