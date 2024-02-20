package cantor

type complement[T comparable] struct {
	arg Container[T]
}

func newComplement[T comparable](arg Container[T]) DerivedImplicitSet[T] {
	return complement[T]{
		arg: arg,
	}
}

func (set complement[T]) Contains(element T) bool {
	return !set.arg.Contains(element)
}

func (set complement[T]) Union(other Container[T]) DerivedImplicitSet[T] {
	return newImplicitUnion[T](set, other)
}

func (set complement[T]) Intersect(other Container[T]) DerivedImplicitSet[T] {
	return newImplicitIntersection[T](set, other)
}

func (set complement[T]) Complement() DerivedImplicitSet[T] {
	return newComplement[T](set)
}
