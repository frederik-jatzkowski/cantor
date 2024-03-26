package cantor

// [ImplicitSet] represents a set, which is only defined by an arbitrary [Predicate].
// Thus, an [ImplicitSet] can represent an infinite amount of elements without performance or memory overhead.
// Due to its unconstrained nature, this type of set can only be used for lookups.
//
// [ImplicitSet] implements [Container].
type ImplicitSet[T comparable] Predicate[T]

// [NewImplicitSet] returns an [ImplicitSet] representing all elements e where the given [Predicate] returns true.
func NewImplicitSet[T comparable](predicate Predicate[T]) ImplicitSet[T] {
	return ImplicitSet[T](predicate)
}

func (predicate ImplicitSet[T]) Contains(element T) bool {
	return predicate(element)
}

// Union returns an [ImplicitSet] set representing the set union of its arguments.
//
// The result is a data view and will reflect future changes of the underlying structures.
func (predicate ImplicitSet[T]) Union(other Container[T]) ImplicitSet[T] {
	return func(element T) bool {
		return predicate(element) || other.Contains(element)
	}
}

// Intersect returns an [ImplicitSet] set representing the set intersection of its arguments.
//
// The result is a data view and will reflect future changes of the underlying structures.
func (predicate ImplicitSet[T]) Intersect(other Container[T]) ImplicitSet[T] {
	return func(element T) bool {
		return predicate(element) && other.Contains(element)
	}
}

// Complement returns an [ImplicitSet], that contains all elements not contained in this [ImplicitSet].
//
// The result is a data view and will reflect future changes of the underlying structures.
func (predicate ImplicitSet[T]) Complement() ImplicitSet[T] {
	return func(element T) bool {
		return !predicate.Contains(element)
	}
}

// Difference returns an [ImplicitSet] set with all elements of this [ImplicitSet],
// which are not contained in the argument.
//
// The result is a data view and will reflect future changes of the underlying structures.
func (predicate ImplicitSet[T]) Difference(other Container[T]) ImplicitSet[T] {
	return func(element T) bool {
		return predicate(element) && !other.Contains(element)
	}
}

// SymmetricDifference returns an [ImplicitSet] set with all elements of this [ImplicitSet] and the other [Container],
// which are contained in exactly one of the two.
//
// The result is a data view and will reflect future changes of the underlying structures.
func (predicate ImplicitSet[T]) SymmetricDifference(other Container[T]) ImplicitSet[T] {
	return func(element T) bool {
		return predicate(element) != other.Contains(element)
	}
}
