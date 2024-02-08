// Package cantor provides a comprehensive implementation of set theory, only constrained by computational practicality.
// In addition to basic set data structures, it  features performant lazy evaluation,
// infinite sets and set comprehension, exposed through a type safe, generic and extensible api.
package cantor

// Container represents any structure, which can implicitly or explicity contain elements.
// Container is implemented by [ImplicitSet], [EnumerableSet] and [ExplicitSet].
type Container[T comparable] interface {
	// Contains must be a deterministic predicate and must not create side effects.
	Contains(element T) bool
}

// Enumerator represents any structure, which can enumerate a predetermined amount of elements.
// Enumerator is implemented by [EnumerableSet] and [ExplicitSet].
type Enumerator[T comparable] interface {
	// Enumerate calls callback for each element once and stops if callback returns true.
	Enumerate(callback func(element T) (stop bool))

	// Size returns the number of elements, that Enumerate would provide.
	Size() int
}

// Evaluator represents any structure, that can be evaluated into an [ExplicitSet].
// Evaluator is implemented by [EnumerableSet] and [ExplicitSet].
type Evaluator[T comparable] interface {
	// Evaluate will evaluate the underlying structure into a new and independent ExplicitSet.
	Evaluate() ExplicitSet[T]
}

// ImplicitSet represents a set, which is only defined by an arbitrary predicate, its Contains()-method.
// Thus, an ImplicitSet can represent an infinite amount of elements without performance or memory overhead.
// Due to the unconstrained nature, this type of set can only be used in places, where no enumeration of elements is required.
type ImplicitSet[T comparable] interface {
	Container[T]

	// Union provides the set union of this ImplicitSet and any other value with a Contains()-method as an ImplicitSet.
	Union(other Container[T]) ImplicitSet[T]

	// Intersect provides the set intersection of this ImplicitSet and any other value with a Contains()-method as an ImplicitSet.
	Intersect(other Container[T]) ImplicitSet[T]

	// Complement provides an ImplicitSet, which contains all elements that are not contained in this ImplicitSet.
	Complement() ImplicitSet[T]
}

// EnumerableSet can be understood as an intermediate between [ImplicitSet] and [ExplicitSet].
// While the exact size of the set are known and its elements can be enumerated, this type of set is read only.
// One real world example of this are intervals over integer numbers.
// EnumerableSet is also implemented by [ExplicitSet].
type EnumerableSet[T comparable] interface {
	Container[T]
	Enumerator[T]
	Evaluator[T]

	// Union provides the set union of this EnumerableSet with another EnumerableSet.
	Union(other EnumerableSet[T]) EnumerableSet[T]

	// Intersect provides the set intersection of this EnumerableSet and any other value with a Contains()-method as a EnumerableSet.
	Intersect(other Container[T]) EnumerableSet[T]

	// Complement provides an ImplicitSet, which contains all elements that are not contained in this EnumerableSet.
	// Since the result might be infinite, its elements cannot be enumerated and it has to be defined implicitly.
	Complement() ImplicitSet[T]
}

// ExplicitSet allows for explicit declaration of the contained elements.
// This generally comes at the cost, that all elements must be held in memory,
// but at the same time ensures, that Size(), Enumerate() and Contains() are extremely fast.
type ExplicitSet[T comparable] interface {
	EnumerableSet[T]

	// After set.Add(e) was called, set.Contains(e) must return true.
	// The return value indicates, if the value was not present before and actually was added.
	Add(element T) (wasAdded bool)

	// After set.Remove(e) was called, set.Contains(e) must return false.
	// The return value indicates, if the value was present before and actually was removed.
	Remove(element T) (wasRemoved bool)
}
