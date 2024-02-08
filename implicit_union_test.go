package cantor_test

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
)

func Test_implicitUnion_Union(t *testing.T) {
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

	union12 := set1.Union(set2)
	union34 := set3.Union(set4)

	union := union12.Union(union34)

	if !union.Contains(4) {
		t.Errorf("set should contain 4")
	}

	if !union.Contains(9) {
		t.Errorf("set should contain 9")
	}

	if !union.Contains(25) {
		t.Errorf("set should contain 25")
	}

	if !union.Contains(49) {
		t.Errorf("set should contain 49")
	}

	if union.Contains(1) {
		t.Errorf("set should not contain 1")
	}
}

func Test_implicitUnion_Contains(t *testing.T) {
	set1 := cantor.NewImplicitSet(func(element int) bool {
		return element%2 == 0
	})
	set2 := cantor.NewImplicitSet(func(element int) bool {
		return element%3 == 0
	})

	union := set1.Union(set2)

	if !union.Contains(2) {
		t.Errorf("set should contain 2")
	}

	if union.Contains(1) {
		t.Errorf("set should not contain 1")
	}

	if !union.Contains(-3) {
		t.Errorf("set should contain -2")
	}

	if union.Contains(-1) {
		t.Errorf("set should not contain -1")
	}
}

func Test_implicitUnion_Intersection(t *testing.T) {
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

	union12 := set1.Union(set2)
	union34 := set3.Union(set4)

	intersection := union12.Intersect(union34)

	if !intersection.Contains(2 * 3 * 5 * 7) {
		t.Errorf("set should contain %d", 2*3*5*7)
	}

	if intersection.Contains(3) {
		t.Errorf("set should not contain 3")
	}
}

func Test_implicitUnion_Complement(t *testing.T) {
	set1 := cantor.NewImplicitSet(func(element int) bool {
		return element%2 == 0
	})
	set2 := cantor.NewImplicitSet(func(element int) bool {
		return element%3 == 0
	})

	union := set1.Union(set2)
	complement := union.Complement()

	if !complement.Contains(7) {
		t.Errorf("set should contain 7")
	}

	if complement.Contains(2 * 3) {
		t.Errorf("set should not contain %d", 2*3)
	}
}
