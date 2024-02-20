// Package cantor provides a comprehensive implementation of set theory, only constrained by computational practicality.
// In addition to basic set data structures, it  features performant lazy evaluation,
// infinite sets and set comprehension, exposed through a type safe, generic and extensible api.
package cantor

import "fmt"

// [Container] represents any structure, which can implicitly or explicity contain elements.
type Container[T comparable] interface {
	// Contains must be a deterministic predicate and must not create side effects.
	Contains(element T) bool
}

// [ImplicitSet] represents a set, which is only defined by an arbitrary predicate, its Contains-method.
// Thus, an [ImplicitSet] can represent an infinite amount of elements without performance or memory overhead.
// Due to the unconstrained nature, this type of set can only be used in places,
// where no enumeration of elements is required.
type ImplicitSet[T comparable] interface {
	Container[T]

	// Union provides the set union of this ImplicitSet and any other value with a Contains-method as a DerivedImplicitSet.
	Union(other Container[T]) DerivedImplicitSet[T]

	// Intersect provides the set intersection of this ImplicitSet and any other value with a
	// Contains-method as a DerivedImplicitSet.
	Intersect(other Container[T]) DerivedImplicitSet[T]

	// Complement provides a DerivedImplicitSet, which contains all elements that are not contained in this ImplicitSet.
	Complement() DerivedImplicitSet[T]
}

// [Set] represents a collection of unique and enumerable elements, which has a limited, known size.
//
// [Set] is implemented by [HashSet].
type Set[T comparable] interface {
	Container[T]
	fmt.Stringer

	// Size returns the number of unique elements in this set.
	Size() int

	// Iterator returns a function iterator (https://go.dev/wiki/RangefuncExperiment), which can be used for iteration.
	// This function iterator can be used to yield the elements of a set one by one.
	// Iteration is stopped, if the yield function returns false.
	Iterator() FunctionIterator[T]

	// Union returns a DerivedSet containing the set union of two Sets.
	Union(other Set[T]) DerivedSet[T]

	// Intersect returns a DerivedSet containing the set intersection of this Set and a Container.
	Intersect(other Container[T]) DerivedSet[T]

	// Complement provides a DerivedImplicitSet, which contains all elements that are not contained in this Set.
	// Since the result might be infinite, its elements cannot be enumerated anymore and is only implicit.
	Complement() DerivedImplicitSet[T]

	// Evaluate will evaluate the underlying Set into a new and independent HashSet.
	// It is guaranteed, that the result is not influenced by changes to the Evaluator or vice-versa.
	Evaluate() HashSet[T]
}

// [DerivedSet] is an alias to [Set].
// It indicates that this set is derived from other sets and will reflect changes made to those underlying sets.
// Additionally, methods on a [DerivedSet] might require underlying evaluation,
// that is more expensive than calling the same methods on an evaluated [HashSet] directly.
type DerivedSet[T comparable] Set[T]

// [DerivedImplicitSet] is an alias to [ImplicitSet].
// It indicates that this set is derived from other sets and will reflect changes made to those underlying sets.
type DerivedImplicitSet[T comparable] ImplicitSet[T]

// [FunctionIterator] is a function that can be used to iterate over elements.
// Iteration will be started by calling the iterator with a yield callback.
// This callback will be run for each element.
// Once the yield function returns false, iteration is stopped, just like a break-statement in a loop.
//
// This interface is inspired by the rangefunc experiment: https://go.dev/wiki/RangefuncExperiment.
type FunctionIterator[T comparable] func(yield func(element T) (next bool))
