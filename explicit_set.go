package cantor

type explicitSet[T comparable] struct {
	elements map[T]struct{}
}

func NewExplicitSet[T comparable](elements ...T) ExplicitSet[T] {
	result := explicitSet[T]{
		elements: make(map[T]struct{}, len(elements)),
	}

	for _, element := range elements {
		result.elements[element] = struct{}{}
	}

	return result
}

func (set explicitSet[T]) Add(element T) bool {
	before := len(set.elements)
	set.elements[element] = struct{}{}

	return before < len(set.elements)
}

func (set explicitSet[T]) Remove(element T) bool {
	before := len(set.elements)
	delete(set.elements, element)

	return before > len(set.elements)
}

func (set explicitSet[T]) Contains(element T) bool {
	_, contains := set.elements[element]

	return contains
}

func (set explicitSet[T]) Union(other EnumerableSet[T]) EnumerableSet[T] {
	return union[T]{
		args: []EnumerableSet[T]{set, other},
	}
}

func (set explicitSet[T]) Intersect(other Container[T]) EnumerableSet[T] {
	return intersection[T]{
		arg:  set,
		args: []Container[T]{other},
	}
}

func (set explicitSet[T]) Complement() ImplicitSet[T] {
	return complement[T]{
		inner: set,
	}
}

func (set explicitSet[T]) Enumerate(callback func(element T) (stop bool)) {
	for element := range set.elements {
		if callback(element) {
			break
		}
	}
}

func (set explicitSet[T]) Size() int {
	return len(set.elements)
}

func (set explicitSet[T]) Evaluate() ExplicitSet[T] {
	return evaluate[T](set)
}
