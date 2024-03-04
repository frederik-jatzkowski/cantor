// Package cantor provides a comprehensive implementation of set operations,
// only constrained by computational practicality.
// In addition to basic set data structures, it  features performant lazy-evaluation,
// infinite sets and set comprehension, exposed through a type-safe, generic, and extensible API.
package cantor

import "fmt"

// [Container] represents any structure, which can implicitly or explicitly contain elements.
//
// [Container] is implemented by [ImplicitSet] and extended by [IterableSet].
type Container[T any] interface {
	// Contains must be a deterministic predicate and must not create side effects.
	Contains(element T) bool
}

// [Predicate] is a type of function that receives an element and returns a boolean value.
// It can be used to define an [ImplicitSet].
type Predicate[T any] func(element T) bool

// [Iterator] is a function that can be used to iterate over elements.
// Iteration starts when the iterator is called with a yield callback.
// This callback will be run for each element.
// Once the yield callback returns false, iteration is stopped, just like a break-statement in a loop.
//
// This interface is inspired by the rangefunc experiment: https://go.dev/wiki/RangefuncExperiment.
type Iterator[T any] func(yield func(element T) (next bool))

// [IterableSet] represents a collection of unique and enumerable elements, which has a limited, known size.
// The elements can be iterated using an [Iterator].
//
// [IterableSet] is extended by [MutableSet].
type IterableSet[T comparable] interface {
	Container[T]
	fmt.Stringer

	// Elements returns an Iterator over the elements of this IterableSet.
	Elements() Iterator[T]

	// Size returns the number of unique elements in this IterableSet.
	Size() int

	// Union returns an IterableSet representing the set union of its arguments.
	Union(other IterableSet[T]) IterableSet[T]

	// Intersect returns a IterableSet representing the set intersection of its arguments.
	Intersect(other Container[T]) IterableSet[T]

	// Complement provides an ImplicitSet, which represents all elements that are not contained in this Set.
	Complement() ImplicitSet[T]
}

// [MutableSet] represents an [IterableSet], where elements can freely be added or removed.
//
// [MutableSet] is implemented by [HashSet].
type MutableSet[T comparable] interface {
	IterableSet[T]

	// Add adds element and returns true if this operation actually changed the MutableSet.
	// If the element was already contained, this leaves the set unchanged and returns false.
	//
	// This change will be reflected in sets, which are derived from this set.
	Add(element T) (setChanged bool)

	// Remove removes element and returns true if this operation actually changed the MutableSet.
	// If the element was not in the set, this leaves the set unchanged and returns false.
	//
	// This change will be reflected in sets, which are derived from this set.
	Remove(element T) (setChanged bool)
}
