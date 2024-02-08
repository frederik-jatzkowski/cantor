package cantor

type intersection[T comparable] struct {
	arg  EnumerableSet[T]
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

func (set intersection[T]) Union(other EnumerableSet[T]) EnumerableSet[T] {
	return union[T]{
		args: []EnumerableSet[T]{set, other},
	}
}

func (set intersection[T]) Intersect(other Container[T]) EnumerableSet[T] {
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

func (set intersection[T]) Enumerate(callback func(element T) (stop bool)) {
	set.arg.Enumerate(func(element T) (stop bool) {
		for _, arg := range set.args {
			if !arg.Contains(element) {
				return false
			}
		}

		return callback(element)
	})
}

func (set intersection[T]) Size() int {
	size := 0

	set.Enumerate(func(element T) (stop bool) {
		size++

		return false
	})

	return size
}

func (set intersection[T]) Evaluate() ExplicitSet[T] {
	return evaluate[T](set)
}
