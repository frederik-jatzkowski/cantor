package cantor

type implicitUnion[T comparable] struct {
	args []Container[T]
}

func newImplicitUnion[T comparable](args ...Container[T]) DerivedImplicitSet[T] {
	return implicitUnion[T]{
		args: args,
	}
}

func (set implicitUnion[T]) Contains(element T) bool {
	for _, arg := range set.args {
		if arg.Contains(element) {
			return true
		}
	}

	return false
}

func (set implicitUnion[T]) Union(other Container[T]) DerivedImplicitSet[T] {
	return newImplicitUnion[T](append(set.args, other)...)
}

func (set implicitUnion[T]) Intersect(other Container[T]) DerivedImplicitSet[T] {
	return newImplicitIntersection[T](set, other)
}

func (set implicitUnion[T]) Complement() DerivedImplicitSet[T] {
	return newComplement[T](set)
}
