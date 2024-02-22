// Package cantor provides a comprehensive implementation of set operations,
// only constrained by computational practicality.
// In addition to basic set data structures, it  features performant lazy evaluation,
// infinite sets and set comprehension, exposed through a type safe, generic and extensible api.
package cantor

import "fmt"

// [Container] represents any structure, which can implicitly or explicitly contain elements.
//
// [Container] is extended by [ImplicitSet] and [IterableContainer].
type Container[T comparable] interface {
	// Contains must be a deterministic predicate and must not create side effects.
	Contains(element T) bool
}

// [Predicate] is a type of function that receives an element and returns a boolean value, indicating set membership.
type Predicate[T comparable] func(element T) bool

// [IterableContainer] is a [Container] whose elements can be iterated.
// For use in cantor set expressions, all iterated elements must be deduplicated.
//
// [IterableContainer] is extended by [DerivedSet].
type IterableContainer[T comparable] interface {
	Container[T]

	// Iterator returns a function iterator (https://go.dev/wiki/RangefuncExperiment), which can be used for iteration.
	// This function iterator can be used to yield the elements of a set one by one.
	// Iteration is stopped, if the yield function returns false.
	Iterator() FunctionIterator[T]
}

// [FunctionIterator] is a function that can be used to iterate over elements.
// Iteration will be started by calling the iterator with a yield callback.
// This callback will be run for each element.
// Once the yield function returns false, iteration is stopped, just like a break-statement in a loop.
//
// This interface is inspired by the rangefunc experiment: https://go.dev/wiki/RangefuncExperiment.
type FunctionIterator[T comparable] func(yield func(element T) (next bool))

// [DerivedSet] represents a set derived from other sets via set expressions.
// Method calls on a [DerivedSet] are computed just in time and
// subsequent calls will reflect changes made to the underlying sets.
// This requires underlying evaluation, which might be more computationally expensive than
// equivalent calls on a [HashSet].
// Since a [DerivedSet] is not fully evaluated, it supports less operations than a [Set].
//
// To obtain a [Set] from a [DerivedSet], use the IntoHashSet-method.
//
// [DerivedSet] is extended by [Set].
type DerivedSet[T comparable] interface {
	IterableContainer[T]
	fmt.Stringer

	// Union returns a DerivedSet representing the set union of its arguments.
	Union(other IterableContainer[T]) DerivedSet[T]

	// Intersect returns a DerivedSet representing the set intersection its arguments.
	Intersect(other Container[T]) DerivedSet[T]

	// Complement provides an ImplicitSet, which represents all elements that are not contained in this Set.
	// This result might be infinite, thus its elements cannot be iterated and it can only be defined implicitly.
	Complement() ImplicitSet[T]

	// IntoHashSet will create an independent HashSet from the elements of this DerivedSet.
	IntoHashSet() HashSet[T]
}

// [Set] represents a collection of unique and enumerable elements, which has a limited, known size.
// As a [Set] is fully evaluated, it supports more operations than a [DerivedSet].
//
// [Set] is implemented by [HashSet].
type Set[T comparable] interface {
	DerivedSet[T]

	// Size returns the number of unique elements in this Set.
	Size() int
}

// [ImplicitSet] represents a set, which is only defined by an arbitrary predicate, its Contains-method.
// Thus, an [ImplicitSet] can represent an infinite amount of elements without performance or memory overhead.
// Due to its unconstrained nature, this type of set can only be used in places,
// where no iteration of elements is required.
type ImplicitSet[T comparable] interface {
	Container[T]

	// Union returns an Implicit set representing the set union of its arguments.
	Union(other Container[T]) ImplicitSet[T]

	// Intersect returns an Implicit set representing the set intersection of its arguments.
	Intersect(other Container[T]) ImplicitSet[T]

	// Complement returns an ImplicitSet, that contains all elements where set.Contains() is false.
	Complement() ImplicitSet[T]
}
