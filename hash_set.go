package cantor

// HashSet implements [Set] using an underlying map.
type HashSet[T comparable] struct {
	elements map[T]struct{}
}

func NewHashSet[T comparable](elements ...T) HashSet[T] {
	result := HashSet[T]{
		elements: make(map[T]struct{}, len(elements)),
	}

	for _, element := range elements {
		result.elements[element] = struct{}{}
	}

	return result
}

// Add implements [HashSet.Add].
func (set HashSet[T]) Add(element T) bool {
	before := len(set.elements)
	set.elements[element] = struct{}{}

	return before < len(set.elements)
}

func (set HashSet[T]) Remove(element T) bool {
	before := len(set.elements)
	delete(set.elements, element)

	return before > len(set.elements)
}

func (set HashSet[T]) Contains(element T) bool {
	_, contains := set.elements[element]

	return contains
}

func (set HashSet[T]) Union(other IterableSet[T]) IterableSet[T] {
	return newIterableUnion[T](set, other)
}

func (set HashSet[T]) Intersect(other Container[T]) IterableSet[T] {
	return newIterableIntersection[T](set, other)
}

func (set HashSet[T]) Complement() ImplicitSet[T] {
	return newComplement[T](set)
}

func (set HashSet[T]) Iter() (rangefunc func(yield func(element T) (next bool))) {
	return func(yield func(element T) (next bool)) {
		for element := range set.elements {
			if !yield(element) {
				break
			}
		}
	}
}

func (set HashSet[T]) Size() int {
	return len(set.elements)
}

func (set HashSet[T]) Evaluate() Set[T] {
	return evaluate[T](set)
}

func (set HashSet[T]) String() string {
	return toString[T](set)
}
