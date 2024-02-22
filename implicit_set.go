package cantor

type implicitSet[T comparable] struct {
	predicate Predicate[T]
}

// NewImplicitSet returns an [ImplicitSet] representing all elements e where the given [Predicate] returns true.
// By using closures, the behaviour of the predicate might change, effectively changing this set.
// Such changes might be reflected in sets derived from this one.
func NewImplicitSet[T comparable](predicate Predicate[T]) ImplicitSet[T] {
	return implicitSet[T]{
		predicate: predicate,
	}
}

func (set implicitSet[T]) Contains(element T) bool {
	return set.predicate(element)
}

func (set implicitSet[T]) Union(other Container[T]) ImplicitSet[T] {
	return newImplicitUnion[T](set, other)
}

func (set implicitSet[T]) Intersect(other Container[T]) ImplicitSet[T] {
	return newImplicitIntersection[T](set, other)
}

func (set implicitSet[T]) Complement() ImplicitSet[T] {
	return newComplement[T](set)
}
