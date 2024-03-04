package cantor

// [ImplicitSet] represents a set, which is only defined by an arbitrary predicate, its Contains-method.
// Thus, an [ImplicitSet] can represent an infinite amount of elements without performance or memory overhead.
// Due to its unconstrained nature, this type of set can only be used in places,
// where no iteration of elements is required.
type ImplicitSet[T comparable] struct {
	predicate Predicate[T]
}

// [NewImplicitSet] returns an [ImplicitSet] representing all elements e where the given [Predicate] returns true.
// By using closures, the behaviour of the predicate might change, effectively changing this set.
// Such changes might be reflected in sets derived from this one.
func NewImplicitSet[T comparable](predicate Predicate[T]) ImplicitSet[T] {
	return ImplicitSet[T]{
		predicate: predicate,
	}
}

func (set ImplicitSet[T]) Contains(element T) bool {
	return set.predicate(element)
}

// Union returns an [ImplicitSet] set representing the set union of its arguments.
func (set ImplicitSet[T]) Union(other Container[T]) ImplicitSet[T] {
	return NewImplicitSet(func(element T) bool {
		return set.Contains(element) || other.Contains(element)
	})
}

// Intersect returns an [ImplicitSet] set representing the set intersection of its arguments.
func (set ImplicitSet[T]) Intersect(other Container[T]) ImplicitSet[T] {
	return NewImplicitSet(func(element T) bool {
		return set.Contains(element) && other.Contains(element)
	})
}

// Complement returns an [ImplicitSet], that contains all elements where set.Contains() is false.
func (set ImplicitSet[T]) Complement() ImplicitSet[T] {
	return NewImplicitSet(func(element T) bool {
		return !set.Contains(element)
	})
}
