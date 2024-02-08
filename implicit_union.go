package cantor

type implicitUnion[T comparable] struct {
	args []Container[T]
}

func (set implicitUnion[T]) Contains(element T) bool {
	for _, arg := range set.args {
		if arg.Contains(element) {
			return true
		}
	}

	return false
}

func (set implicitUnion[T]) Union(other Container[T]) ImplicitSet[T] {
	return implicitUnion[T]{
		args: append(set.args, other),
	}
}

func (set implicitUnion[T]) Intersect(other Container[T]) ImplicitSet[T] {
	return implicitIntersection[T]{
		args: []Container[T]{set, other},
	}
}

func (set implicitUnion[T]) Complement() ImplicitSet[T] {
	return complement[T]{
		inner: set,
	}
}
