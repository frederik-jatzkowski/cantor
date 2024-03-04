package cantor

type intersection[T comparable] struct {
	arg  IterableSet[T]
	args []Container[T]
}

func newIntersection[T comparable](arg IterableSet[T], args ...Container[T]) IterableSet[T] {
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

func (set intersection[T]) Union(other IterableSet[T]) IterableSet[T] {
	return newUnion[T](set, other)
}

func (set intersection[T]) Intersect(other Container[T]) IterableSet[T] {
	return newIntersection[T](set.arg, append(set.args, other)...)
}

func (set intersection[T]) Complement() ImplicitSet[T] {
	return NewImplicitSet(func(element T) bool {
		return !set.Contains(element)
	})
}

func (set intersection[T]) Elements() Iterator[T] {
	return func(yield func(element T) (next bool)) {
		set.arg.Elements()(func(element T) (next bool) {
			for _, arg := range set.args {
				if !arg.Contains(element) {
					return true
				}
			}

			return yield(element)
		})
	}
}

func (set intersection[T]) String() string {
	return toString[T](set)
}

func (set intersection[T]) Size() int {
	return count(set.Elements())
}
