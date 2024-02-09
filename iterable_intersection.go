package cantor

type intersection[T comparable] struct {
	arg  IterableSet[T]
	args []Container[T]
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
	return union[T]{
		args: []IterableSet[T]{set, other},
	}
}

func (set intersection[T]) Intersect(other Container[T]) IterableSet[T] {
	return intersection[T]{
		arg:  set,
		args: append(set.args, other),
	}
}

func (set intersection[T]) Complement() ImplicitSet[T] {
	return complement[T]{
		inner: set,
	}
}

func (set intersection[T]) Iter() (rangefunc func(yield func(element T) (next bool))) {
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

func (set intersection[T]) Size() int {
	size := 0

	set.Iter()(func(element T) (next bool) {
		size++

		return true
	})

	return size
}

func (set intersection[T]) Evaluate() ExplicitSet[T] {
	return evaluate[T](set)
}
