// Package cantor provides a comprehensive implementation of set theory, only constrained by computational practicality.
// In addition to basic set data structures, it  features performant lazy evaluation,
// infinite sets and set comprehension, exposed through a type safe, generic and extensible api.
package cantor

import "fmt"

// [Container] represents any structure, which can implicitly or explicity contain elements.
//
// [Container] is extended by [ImplicitSet], [ReadableSet] and [Set].
type Container[T comparable] interface {
	// Contains must be a deterministic predicate and must not create side effects.
	Contains(element T) bool
}

// [DistinctIterator] represents any structure, which can iterate over its elements,
// ensuring that no two elements given are equal.
//
// [DistinctIterator] is extended by [ReadableSet] and [Set].
type DistinctIterator[T comparable] interface {
	// IterateDistinct returns a rangefunc (https://go.dev/wiki/RangefuncExperiment), which can be used for iteration.
	// This rangefunc can be used to yield the elements of a set one by one.
	// Iteration is stopped, if the yield function returns false.
	IterateDistinct() (rangefunc func(yield func(element T) (next bool)))
}

// [Evaluator] represents any structure, that can be evaluated into an [Set].
//
// [Evaluator] is extended by [ReadableSet] and [Set].
type Evaluator[T comparable] interface {
	// Evaluate will evaluate the underlying structure into a new and independent Set.
	// It is guaranteed, that the result is not influenced by changes to the Evaluator or vice-versa.
	Evaluate() Set[T]
}

// [ImplicitSet] represents a set, which is only defined by an arbitrary predicate, its Contains()-method.
// Thus, an [ImplicitSet] can represent an infinite amount of elements without performance or memory overhead.
// Due to the unconstrained nature, this type of set can only be used in places,
// where no enumeration of elements is required.
type ImplicitSet[T comparable] interface {
	Container[T]

	// Union provides the set union of this ImplicitSet and any other value with a Contains()-method as an ImplicitSet.
	Union(other Container[T]) ImplicitSet[T]

	// Intersect provides the set intersection of this ImplicitSet and any other value with a
	// Contains()-method as an ImplicitSet.
	Intersect(other Container[T]) ImplicitSet[T]

	// Complement provides an ImplicitSet, which contains all elements that are not contained in this ImplicitSet.
	Complement() ImplicitSet[T]
}

// [ReadableSet] can be understood as an intermediate between [ImplicitSet] and [Set].
// While the set has a determined size and its elements can be iterated, it does not allow modification.
// One real world example of this are intervals over integer numbers.
//
// [ReadableSet] is extended by [Set].
type ReadableSet[T comparable] interface {
	Container[T]
	DistinctIterator[T]
	Evaluator[T]
	fmt.Stringer

	// Size returns the number of unique elements in this set.
	Size() int

	// Union provides the set union of this ReadableSet with another ReadableSet.
	Union(other ReadableSet[T]) ReadableSet[T]

	// Intersect provides the set intersection of this ReadableSet and any other value with a
	// Contains()-method as a ReadableSet.
	Intersect(other Container[T]) ReadableSet[T]

	// Complement provides an ImplicitSet, which contains all elements that are not contained in this ReadableSet.
	// Since the result might be infinite, its elements cannot be enumerated and it has to be defined implicitly.
	Complement() ImplicitSet[T]
}

// [Set] represents sets as collections of arbitrary elements of type T.
// This freedom usually requires, that all elements are stored in memory.
// Operations on this type of set are usually very quick.
//
// [Set] is implemented by [HashSet].
type Set[T comparable] interface {
	ReadableSet[T]

	// After set.Add(e) was called, set.Contains(e) must return true.
	// The return value indicates, if the value was not present before and actually was added.
	Add(element T) (wasAdded bool)

	// After set.Remove(e) was called, set.Contains(e) must return false.
	// The return value indicates, if the value was present before and actually was removed.
	Remove(element T) (wasRemoved bool)
}
