// Package cantor provides a comprehensive implementation of set theory, only constrained by computational practicality.
// In addition to basic set data structures, it  features performant lazy evaluation,
// infinite sets and set comprehension, exposed through a type safe, generic and extensible api.
package cantor

import "fmt"

// Container represents any structure, which can implicitly or explicity contain elements.
// Container is implemented by [ImplicitSet], [IterableSet] and [ExplicitSet].
type Container[T comparable] interface {
	// Contains must be a deterministic predicate and must not create side effects.
	Contains(element T) bool
}

// Iterable represents any structure, which can iterate over a predetermined number of elements.
// Iterable is implemented by [IterableSet] and [ExplicitSet].
type Iterable[T comparable] interface {
	// Iter returns a rangefunc (https://go.dev/wiki/RangefuncExperiment), which can be used to iterate over all elements.
	Iter() (rangefunc func(yield func(element T) (next bool)))

	// Size returns the number of elements to iterate over.
	Size() int
}

// Evaluator represents any structure, that can be evaluated into an [ExplicitSet].
// Evaluator is implemented by [IterableSet] and [ExplicitSet].
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

// IterableSet can be understood as an intermediate between [ImplicitSet] and [ExplicitSet].
// While the exact size of the set is known and its elements can be enumerated, it does not allow modification.
// One real world example of this are intervals over integer numbers.
// IterableSet is also implemented by [ExplicitSet].
type IterableSet[T comparable] interface {
	Container[T]
	Iterable[T]
	Evaluator[T]
	fmt.Stringer

	// Union provides the set union of this IterableSet with another IterableSet.
	Union(other IterableSet[T]) IterableSet[T]

	// Intersect provides the set intersection of this IterableSet and any other value with a Contains()-method as a IterableSet.
	Intersect(other Container[T]) IterableSet[T]

	// Complement provides an ImplicitSet, which contains all elements that are not contained in this IterableSet.
	// Since the result might be infinite, its elements cannot be enumerated and it has to be defined implicitly.
	Complement() ImplicitSet[T]
}

// ExplicitSet allows for explicit declaration of the contained elements.
// This generally comes at the cost, that all elements must be held in memory,
// but at the same time ensures, that Size(), Enumerate() and Contains() are extremely fast.
type ExplicitSet[T comparable] interface {
	IterableSet[T]

	// After set.Add(e) was called, set.Contains(e) must return true.
	// The return value indicates, if the value was not present before and actually was added.
	Add(element T) (wasAdded bool)

	// After set.Remove(e) was called, set.Contains(e) must return false.
	// The return value indicates, if the value was present before and actually was removed.
	Remove(element T) (wasRemoved bool)
}
