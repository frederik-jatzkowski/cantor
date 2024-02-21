package cantor

type intersection[T comparable] struct {
	arg  IterableContainer[T]
	args []Container[T]
}

func newIntersection[T comparable](arg IterableContainer[T], args ...Container[T]) DerivedSet[T] {
	return intersection[T]{
		arg:  arg,
		args: args,
	}
}

func (set intersection[T]) Contains(element T) bool {
	if !set.arg.Contains(element) {
		return false
	}

	for _, arg := range set.args {
		if !arg.Contains(element) {
			return false
		}
	}

	return true
}

func (set intersection[T]) Union(other IterableContainer[T]) DerivedSet[T] {
	return newUnion[T](set, other)
}

func (set intersection[T]) Intersect(other Container[T]) DerivedSet[T] {
	return newIntersection[T](set.arg, append(set.args, other)...)
}

func (set intersection[T]) Complement() ImplicitSet[T] {
	return newComplement[T](set)
}

func (set intersection[T]) Iterator() FunctionIterator[T] {
	return func(yield func(element T) (next bool)) {
		set.arg.Iterator()(func(element T) (next bool) {
			for _, arg := range set.args {
				if !arg.Contains(element) {
					return true
				}
			}

			return yield(element)
		})
	}
}

func (set intersection[T]) IntoHashSet() HashSet[T] {
	return evaluate[T](set)
}

func (set intersection[T]) String() string {
	return toString[T](set)
}
