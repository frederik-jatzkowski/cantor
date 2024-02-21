package cantor

type complement[T comparable] struct {
	arg Container[T]
}

func newComplement[T comparable](arg Container[T]) ImplicitSet[T] {
	return complement[T]{
		arg: arg,
	}
}

func (set complement[T]) Contains(element T) bool {
	return !set.arg.Contains(element)
}

func (set complement[T]) Union(other Container[T]) ImplicitSet[T] {
	return newImplicitUnion[T](set, other)
}

func (set complement[T]) Intersect(other Container[T]) ImplicitSet[T] {
	return newImplicitIntersection[T](set, other)
}

func (set complement[T]) Complement() ImplicitSet[T] {
	return newComplement[T](set)
}
