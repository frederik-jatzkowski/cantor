package cantor

import "fmt"

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

// [ReadableSet] represents a collection of unique and enumerable elements, which has a limited, known size.
// The elements can be iterated using an [Iterator].
//
// [ReadableSet] is extended by [Set].
type ReadableSet[T comparable] interface {
	Container[T]
	fmt.Stringer

	// Elements returns an Iterator over the elements of this ReadableSet.
	Elements() Iterator[T]

	// Size returns the number of unique elements in this ReadableSet.
	Size() int

	// Union returns a ReadableSet representing the set union of its arguments.
	//
	// The result is a data view and will reflect future changes of the underlying structures.
	Union(other ReadableSet[T]) ReadableSet[T]

	// Intersect returns a ReadableSet representing the set intersection of its arguments.
	//
	// The result is a data view and will reflect future changes of the underlying structures.
	Intersect(other Container[T]) ReadableSet[T]

	// Complement provides an ImplicitSet, which represents all elements that are not contained in this Set.
	//
	// The result is a data view and will reflect future changes of the underlying structures.
	Complement() ImplicitSet[T]

	// Difference returns a ReadableSet representing the set with all elements of this set,
	// which are not contained in the argument.
	//
	// The result is a data view and will reflect future changes of the underlying structures.
	Difference(other Container[T]) ReadableSet[T]
}

// [Set] represents a [ReadableSet], where elements can freely be added or removed.
//
// [Set] is directly implemented by [HashSet].
type Set[T comparable] interface {
	ReadableSet[T]

	// Add adds element and returns true if this operation actually changed the Set.
	// If the element was already contained, this leaves the set unchanged and returns false.
	//
	// Derived data views will reflect this change.
	Add(element T) (setChanged bool)

	// Remove removes element and returns true if this operation actually changed the Set.
	// If the element was not in the set, this leaves the set unchanged and returns false.
	//
	// Derived data views will reflect this change.
	Remove(element T) (setChanged bool)
}
