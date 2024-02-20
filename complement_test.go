package cantor_test

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
)

func Test_complement_Contains(t *testing.T) {
	set := cantor.NewHashSet(1, 2, 3).Complement()

	if !set.Contains(0) {
		t.Errorf("set should contain 0")
	}

	if set.Contains(1) {
		t.Errorf("set should not contain 1")
	}
}

func Test_complement_Union(t *testing.T) {
	set1 := cantor.NewHashSet(1, 2, 3).Complement()
	set2 := cantor.NewHashSet(3, 4, 5).Complement()

	union := set1.Union(set2)

	if !union.Contains(2) {
		t.Errorf("union of complements should contain %d", 2)
	}

	if union.Contains(3) {
		t.Errorf("union of complements should not contain %d", 3)
	}
}

func Test_complement_Intersection(t *testing.T) {
	set1 := cantor.NewHashSet(1, 2, 3).Complement()
	set2 := cantor.NewHashSet(3, 4, 5).Complement()

	intersection := set1.Intersect(set2)

	if intersection.Contains(2) {
		t.Errorf("intersection of complements should not contain %d", 2)
	}

	if intersection.Contains(4) {
		t.Errorf("intersection of complements should not contain %d", 4)
	}

	if !intersection.Contains(0) {
		t.Errorf("intersection of complements not contain %d", 0)
	}
}

func Test_complement_Complement(t *testing.T) {
	set := cantor.NewHashSet(1, 2, 3, 4, 5)
	complementComplement := set.Complement().Complement()

	set.Iter()(func(element int) (next bool) {
		if !complementComplement.Contains(element) {
			t.Errorf("complement of the complement does not contain %d", element)
		}

		return false
	})
}
