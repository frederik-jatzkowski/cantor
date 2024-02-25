package cantor_test

import (
	"testing"

	"github.com/frederik-jatzkowski/cantor"
	"github.com/frederik-jatzkowski/cantor/internal/testsuites"
)

func TestHashSet_Set(t *testing.T) {
	testsuites.RunTestsForSet(t, func(elements ...byte) cantor.Set[byte] {
		return cantor.NewHashSet(elements...)
	})
}

func TestNewHashSetFromIterator(t *testing.T) {
	testsuites.RunTestsForContainer(t, func(elements ...byte) cantor.Container[byte] {
		return cantor.NewHashSetFromIterator(func(yield func(element byte) (next bool)) {
			for _, element := range elements {
				if !yield(element) {
					break
				}
			}
		})
	})
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
