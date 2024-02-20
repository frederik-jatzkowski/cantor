package cantor

type iterableIntersection[T comparable] struct {
	arg  Set[T]
	args []Container[T]
}

func newIterableIntersection[T comparable](arg Set[T], args ...Container[T]) Set[T] {
	return iterableIntersection[T]{
		arg:  arg,
		args: args,
	}
}

func (set iterableIntersection[T]) Contains(element T) bool {
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

func (set iterableIntersection[T]) Union(other Set[T]) DerivedSet[T] {
	return newIterableUnion[T](set, other)
}

func (set iterableIntersection[T]) Intersect(other Container[T]) DerivedSet[T] {
	return newIterableIntersection[T](set.arg, append(set.args, other)...)
}

func (set iterableIntersection[T]) Complement() DerivedImplicitSet[T] {
	return newComplement[T](set)
}

func (set iterableIntersection[T]) Iterator() FunctionIterator[T] {
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

func (set iterableIntersection[T]) Size() int {
	size := 0

	set.Iterator()(func(element T) (next bool) {
		size++

		return true
	})

	return size
}

func (set iterableIntersection[T]) Evaluate() HashSet[T] {
	return evaluate[T](set)
}

func (set iterableIntersection[T]) String() string {
	return toString[T](set)
}
