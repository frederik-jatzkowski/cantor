package cantor

// [HashSet] implements [Set] using an underlying hash map.
type HashSet[T comparable] map[T]struct{}

// [NewHashSet] returns an initialized [HashSet] containing all provided elements.
// The given elements are deduplicated.
func NewHashSet[T comparable](elements ...T) HashSet[T] {
	result := make(map[T]struct{}, len(elements))

	for _, element := range elements {
		result[element] = struct{}{}
	}

	return result
}

// [NewHashSetFromIterator] evaluates the iterator and adds all elements to the resulting [HashSet].
// The given elements are deduplicated.
func NewHashSetFromIterator[T comparable](iterator Iterator[T]) HashSet[T] {
	result := make(map[T]struct{})

	iterator(func(element T) bool {
		result[element] = struct{}{}

		return true
	})

	return result
}

// Add adds element and returns true if this operation actually changed the [HashSet].
// If the element was already contained, this leaves the set unchanged and returns false.
//
// Derived data views will reflect this change.
//
// The time complexity of this method is O(1).
func (set HashSet[T]) Add(element T) (setChanged bool) {
	before := len(set)
	set[element] = struct{}{}

	return before < len(set)
}

// Remove removes element and returns true if this operation actually changed the [HashSet].
// If the element was not in the set, this leaves the set unchanged and returns false.
//
// Derived data views will reflect this change.
//
// The time complexity of this method is O(1).
func (set HashSet[T]) Remove(element T) (setChanged bool) {
	before := len(set)
	delete(set, element)

	return before > len(set)
}

// Contains returns whether the element is contained in this [HashSet].
//
// The time complexity of this method is O(1).
func (set HashSet[T]) Contains(element T) bool {
	_, contains := set[element]

	return contains
}

// Union returns a [ReadableSet] representing the set union of this set and the argument.
//
// The result is a data view and will reflect future changes of the underlying structures.
func (set HashSet[T]) Union(other ReadableSet[T]) ReadableSet[T] {
	return newUnion[T](set, other)
}

// Intersect returns a [ReadableSet] representing the set intersection of this set and the argument.
//
// The result is a data view and will reflect future changes of the underlying structures.
func (set HashSet[T]) Intersect(other Container[T]) ReadableSet[T] {
	return newIntersection[T](set, other)
}

// Complement returns an [ImplicitSet], representing all element not contained in this set.
// This might represent infinitely many elements.
//
// The result is a data view and will reflect future changes of the underlying structures.
func (set HashSet[T]) Complement() ImplicitSet[T] {
	return NewImplicitSet(func(element T) bool {
		return !set.Contains(element)
	})
}

// Difference returns a [ReadableSet] with all elements of this [HashSet],
// which are not contained in the argument.
//
// The result is a data view and will reflect future changes of the underlying structures.
func (set HashSet[T]) Difference(other Container[T]) ReadableSet[T] {
	return set.Intersect(NewImplicitSet[T](func(element T) bool {
		return !other.Contains(element)
	}))
}

// SymmetricDifference returns a ReadableSet representing the set with all elements of this and the other set,
// which are contained in exactly one of the two.
//
// The result is a data view and will reflect future changes of the underlying structures.
func (set HashSet[T]) SymmetricDifference(other ReadableSet[T]) ReadableSet[T] {
	return set.Difference(other).Union(other.Difference(set))
}

// Elements returns an [Iterator] (https://go.dev/wiki/RangefuncExperiment).
// This [Iterator] can be used to yield the elements of a set one by one.
// Iteration is stopped, if the yield function returns false.
func (set HashSet[T]) Elements() Iterator[T] {
	return func(yield func(element T) (next bool)) {
		for element := range set {
			if !yield(element) {
				break
			}
		}
	}
}

// Size returns the number of unique elements contained in this [HashSet].
func (set HashSet[T]) Size() int {
	return len(set)
}

// String implements [fmt.Stringer] for this [HashSet].
func (set HashSet[T]) String() string {
	return toString[T](set)
}
