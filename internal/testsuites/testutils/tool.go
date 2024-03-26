package testutils

func SliceContains[T comparable](e T, elements []T) bool {
	for _, e2 := range elements {
		if e == e2 {
			return true
		}
	}

	return false
}

func AllBytes() (result []byte) {
	for i := 0; i <= 255; i++ {
		result = append(result, byte(i))
	}

	return result
}
