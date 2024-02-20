package cantor_test

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
)

func TestHashSet(t *testing.T) {
	RunAllSetTests(
		func(elements ...int) cantor.Set[int] {
			return cantor.NewHashSet[int](elements...)
		},
		t,
	)
}

func TestHashSet_Add(t *testing.T) {
	set := cantor.NewHashSet(1, 2, 3)

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

func TestHashSet_Remove(t *testing.T) {
	set := cantor.NewHashSet(1, 2, 3)

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
