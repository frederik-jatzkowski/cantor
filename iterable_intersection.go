package cantor

type iterableIntersection[T comparable] struct {
	arg  IterableSet[T]
	args []Container[T]
}

func newIterableIntersection[T comparable](arg IterableSet[T], args ...Container[T]) IterableSet[T] {
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

func (set iterableIntersection[T]) Union(other IterableSet[T]) IterableSet[T] {
	return newIterableUnion[T](set, other)
}

func (set iterableIntersection[T]) Intersect(other Container[T]) IterableSet[T] {
	return newIterableIntersection[T](set.arg, append(set.args, other)...)
}

func (set iterableIntersection[T]) Complement() ImplicitSet[T] {
	return newComplement[T](set)
}

func (set iterableIntersection[T]) Iter() (rangefunc func(yield func(element T) (next bool))) {
	return func(yield func(element T) (next bool)) {
		set.arg.Iter()(func(element T) (next bool) {
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

	set.Iter()(func(element T) (next bool) {
		size++

		return true
	})

	return size
}

func (set iterableIntersection[T]) Evaluate() Set[T] {
	return evaluate[T](set)
}

func (set iterableIntersection[T]) String() string {
	return toString[T](set)
}
