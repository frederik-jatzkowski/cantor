package cantor

import "testing"

func Test_toString(t *testing.T) {
	set := NewExplicitSet(1, 2)
	str := toString[int](set)

	switch str {
	case "{1, 2}", "{2, 1}":
	default:
		t.Errorf("invalid string: %s", str)
	}
}
