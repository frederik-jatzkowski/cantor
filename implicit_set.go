package cantor

type implicitSet[T comparable] struct {
	predicate func(element T) bool
}

// NewImplicitSet returns an [ImplicitSet] representing all elements e where predicate(e) is true.
// By using closures, the behaviour of the predicate might change, effectively changing this set.
// Such changes might be reflected in sets derived from this one.
func NewImplicitSet[T comparable](predicate func(element T) bool) ImplicitSet[T] {
	return implicitSet[T]{
		predicate: predicate,
	}
}

func (set implicitSet[T]) Contains(element T) bool {
	return set.predicate(element)
}

func (set implicitSet[T]) Union(other Container[T]) DerivedImplicitSet[T] {
	return newImplicitUnion[T](set, other)
}

func (set implicitSet[T]) Intersect(other Container[T]) DerivedImplicitSet[T] {
	return newImplicitIntersection[T](set, other)
}

func (set implicitSet[T]) Complement() DerivedImplicitSet[T] {
	return newComplement[T](set)
}
