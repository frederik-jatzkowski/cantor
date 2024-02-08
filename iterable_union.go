package cantor

type union[T comparable] struct {
	args []IterableSet[T]
}

func (set union[T]) Contains(element T) bool {
	for _, arg := range set.args {
		if arg.Contains(element) {
			return true
		}
	}

	return false
}

func (set union[T]) Union(other IterableSet[T]) IterableSet[T] {
	return union[T]{
		args: append(set.args, other),
	}
}

func (set union[T]) Intersect(other Container[T]) IterableSet[T] {
	return intersection[T]{
		arg:  set,
		args: []Container[T]{other},
	}
}

func (set union[T]) Complement() ImplicitSet[T] {
	return complement[T]{
		inner: set,
	}
}

func (set union[T]) Iter() (rangefunc func(yield func(element T) (next bool))) {
	return func(yield func(element T) (next bool)) {
		for i := 0; i < len(set.args); i++ {
			arg := set.args[i]

			arg.Iter()(func(element T) (stop bool) {
				for i2 := 0; i2 < i; i2++ {
					arg2 := set.args[i2]

					if arg2.Contains(element) {
						return true
					}
				}

				return yield(element)
			})
		}
	}
}

func (set union[T]) Size() int {
	size := 0

	set.Iter()(func(element T) (stop bool) {
		size++

		return true
	})

	return size
}

func (set union[T]) Evaluate() ExplicitSet[T] {
	return evaluate[T](set)
}
