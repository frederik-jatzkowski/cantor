package cantor

// [HashSet] implements [Set] using an underlying hash map.
// This allows constant average time complexity for [HashSet.Add], [HashSet.Remove] and [HashSet.Contains].
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

// [NewHashSetFromIterator] evaluates the entire iterator, adding all elements to the resulting [HashSet].
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
// This change will be reflected in sets, which are derived from this set.
func (set HashSet[T]) Add(element T) (setChanged bool) {
	before := len(set)
	set[element] = struct{}{}

	return before < len(set)
}

// Remove removes element and returns true if this operation actually changed the [HashSet].
// If the element was not in the set, this leaves the set unchanged and returns false.
//
// This change will be reflected in sets, which are derived from this set.
func (set HashSet[T]) Remove(element T) (setChanged bool) {
	before := len(set)
	delete(set, element)

	return before > len(set)
}

// Contains returns whether the element is contained in this [HashSet].
func (set HashSet[T]) Contains(element T) bool {
	_, contains := set[element]

	return contains
}

// Union returns a [ReadableSet] representing the set union of this set and the argument.
//
// Any future changes made to the underlying [HashSet] or the other [ReadableSet] will be reflected in the result.
func (set HashSet[T]) Union(other ReadableSet[T]) ReadableSet[T] {
	return newUnion[T](set, other)
}

// Intersect returns a [ReadableSet] representing the set intersection of this set and the argument.
//
// Any changes made to the underlying [HashSet] or the other [Container] will be reflected in the result.
func (set HashSet[T]) Intersect(other Container[T]) ReadableSet[T] {
	return newIntersection[T](set, other)
}

// Complement returns an [ImplicitSet], representing all element not contained in this set.
// This might represent infinitely many elements.
//
// Any changes made to the underlying [HashSet] will be reflected in the result.
func (set HashSet[T]) Complement() ImplicitSet[T] {
	return NewImplicitSet(func(element T) bool {
		return !set.Contains(element)
	})
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

func (set HashSet[T]) String() string {
	return toString[T](set)
}
