package cantor

type implicitIntersection[T comparable] struct {
	args []Container[T]
}

func newImplicitIntersection[T comparable](args ...Container[T]) DerivedImplicitSet[T] {
	return implicitIntersection[T]{
		args: args,
	}
}

func (set implicitIntersection[T]) Contains(element T) bool {
	for _, arg := range set.args {
		if !arg.Contains(element) {
			return false
		}
	}

	return true
}

func (set implicitIntersection[T]) Union(other Container[T]) DerivedImplicitSet[T] {
	return newImplicitUnion[T](set, other)
}

func (set implicitIntersection[T]) Intersect(other Container[T]) DerivedImplicitSet[T] {
	return newImplicitIntersection[T](append(set.args, other)...)
}

func (set implicitIntersection[T]) Complement() DerivedImplicitSet[T] {
	return newComplement[T](set)
}
