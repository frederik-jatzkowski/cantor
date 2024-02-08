package cantor

type complement[T comparable] struct {
	inner Container[T]
}

func (set complement[T]) Contains(element T) bool {
	return !set.inner.Contains(element)
}

func (set complement[T]) Union(other Container[T]) ImplicitSet[T] {
	return implicitUnion[T]{
		args: []Container[T]{set, other},
	}
}

func (set complement[T]) Intersect(other Container[T]) ImplicitSet[T] {
	return implicitIntersection[T]{
		args: []Container[T]{set, other},
	}
}

func (set complement[T]) Complement() ImplicitSet[T] {
	return complement[T]{
		inner: set,
	}
}
