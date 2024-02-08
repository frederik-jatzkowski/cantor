package cantor

type implicitIntersection[T comparable] struct {
	args []Container[T]
}

func (set implicitIntersection[T]) Contains(element T) bool {
	for _, arg := range set.args {
		if !arg.Contains(element) {
			return false
		}
	}

	return true
}

func (set implicitIntersection[T]) Union(other Container[T]) ImplicitSet[T] {
	return implicitUnion[T]{
		args: []Container[T]{set, other},
	}
}

func (set implicitIntersection[T]) Intersect(other Container[T]) ImplicitSet[T] {
	return implicitIntersection[T]{
		args: append(set.args, other),
	}
}

func (set implicitIntersection[T]) Complement() ImplicitSet[T] {
	return complement[T]{
		inner: set,
	}
}
