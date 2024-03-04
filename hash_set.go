package cantor

// [HashSet] implements [MutableSet] using an underlying hash map.
// This allows constant average time complexity for [HashSet.Add], [HashSet.Remove] and [HashSet.Contains].
type HashSet[T comparable] struct {
	elements map[T]struct{}
}

// [NewHashSet] returns an initialized [HashSet] containing all provided elements.
// The given elements are deduplicated.
func NewHashSet[T comparable](elements ...T) HashSet[T] {
	result := HashSet[T]{
		elements: make(map[T]struct{}, len(elements)),
	}

	for _, element := range elements {
		result.elements[element] = struct{}{}
	}

	return result
}

// [NewHashSetFromIterator] evaluates the entire iterator, adding all elements to the resulting [HashSet].
// The given elements are deduplicated.
func NewHashSetFromIterator[T comparable](iterator Iterator[T]) HashSet[T] {
	result := HashSet[T]{
		elements: make(map[T]struct{}),
	}

	iterator(func(element T) bool {
		result.elements[element] = struct{}{}

		return true
	})

	return result
}

// Add adds element and returns true if this operation actually changed the [HashSet].
// If the element was already contained, this leaves the set unchanged and returns false.
//
// This change will be reflected in sets, which are derived from this set.
func (set HashSet[T]) Add(element T) (setChanged bool) {
	before := len(set.elements)
	set.elements[element] = struct{}{}

	return before < len(set.elements)
}

// Remove removes element and returns true if this operation actually changed the [HashSet].
// If the element was not in the set, this leaves the set unchanged and returns false.
//
// This change will be reflected in sets, which are derived from this set.
func (set HashSet[T]) Remove(element T) (setChanged bool) {
	before := len(set.elements)
	delete(set.elements, element)

	return before > len(set.elements)
}

// Contains returns whether the element is contained in this [HashSet].
func (set HashSet[T]) Contains(element T) bool {
	_, contains := set.elements[element]

	return contains
}

// Union returns an [IterableSet] representing the set union of this set and the argument.
//
// Any future changes made to the underlying [HashSet] or the other [IterableSet] will be reflected in the result.
func (set HashSet[T]) Union(other IterableSet[T]) IterableSet[T] {
	return newUnion[T](set, other)
}

// Intersect returns an [IterableSet] representing the set intersection of this set and the argument.
//
// Any changes made to the underlying [HashSet] or the other [Container] will be reflected in the result.
func (set HashSet[T]) Intersect(other Container[T]) IterableSet[T] {
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
		for element := range set.elements {
			if !yield(element) {
				break
			}
		}
	}
}

// Size returns the number of unique elements contained in this [HashSet].
func (set HashSet[T]) Size() int {
	return len(set.elements)
}

func (set HashSet[T]) String() string {
	return toString[T](set)
}
