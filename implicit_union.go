package cantor

type implicitUnion[T comparable] struct {
	args []Container[T]
}

func newImplicitUnion[T comparable](args ...Container[T]) ImplicitSet[T] {
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

func (set implicitUnion[T]) Union(other Container[T]) ImplicitSet[T] {
	return newImplicitUnion[T](append(set.args, other)...)
}

func (set implicitUnion[T]) Intersect(other Container[T]) ImplicitSet[T] {
	return newImplicitIntersection[T](set, other)
}

func (set implicitUnion[T]) Complement() ImplicitSet[T] {
	return newComplement[T](set)
}
