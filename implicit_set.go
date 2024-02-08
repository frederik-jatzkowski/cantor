package cantor

type implicitSet[T comparable] struct {
	predicate func(element T) bool
}

// NewImplicitSet returns an [ImplicitSet] representing all elements e where predicate(e) == true.
func NewImplicitSet[T comparable](predicate func(element T) bool) ImplicitSet[T] {
	return implicitSet[T]{
		predicate: predicate,
	}
}

func (set implicitSet[T]) Contains(element T) bool {
	return set.predicate(element)
}

func (set implicitSet[T]) Union(other Container[T]) ImplicitSet[T] {
	return implicitUnion[T]{
		args: []Container[T]{set, other},
	}
}

func (set implicitSet[T]) Intersect(other Container[T]) ImplicitSet[T] {
	return implicitIntersection[T]{
		args: []Container[T]{set, other},
	}
}

func (set implicitSet[T]) Complement() ImplicitSet[T] {
	return complement[T]{
		inner: set,
	}
}
