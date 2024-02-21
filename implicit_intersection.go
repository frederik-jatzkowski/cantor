package cantor

type implicitIntersection[T comparable] struct {
	args []Container[T]
}

func newImplicitIntersection[T comparable](args ...Container[T]) ImplicitSet[T] {
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

func (set implicitIntersection[T]) Union(other Container[T]) ImplicitSet[T] {
	return newImplicitUnion[T](set, other)
}

func (set implicitIntersection[T]) Intersect(other Container[T]) ImplicitSet[T] {
	return newImplicitIntersection[T](append(set.args, other)...)
}

func (set implicitIntersection[T]) Complement() ImplicitSet[T] {
	return newComplement[T](set)
}
