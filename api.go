// Package cantor provides a comprehensive implementation of set operations,
// only constrained by computational practicality.
// In addition to basic set data structures, it  features performant lazy-evaluation,
// infinite sets exposed through a type-safe and generic API.
package cantor

// [Container] represents any structure, which can implicitly or explicitly contain elements.
// The Contains method must be a deterministic predicate and must not create side effects.
//
// [Container] is directly implemented by [ImplicitSet] and extended by [ReadableSet].
type Container[T any] interface {
	Contains(element T) bool
}
