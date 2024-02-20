package cantor

type iterableUnion[T comparable] struct {
	args []ReadableSet[T]
}

func newIterableUnion[T comparable](args ...ReadableSet[T]) ReadableSet[T] {
	return iterableUnion[T]{
		args: args,
	}
}

func (set iterableUnion[T]) Contains(element T) bool {
	for _, arg := range set.args {
		if arg.Contains(element) {
			return true
		}
	}

	return false
}

func (set iterableUnion[T]) Union(other ReadableSet[T]) ReadableSet[T] {
	return newIterableUnion[T](append(set.args, other)...)
}

func (set iterableUnion[T]) Intersect(other Container[T]) ReadableSet[T] {
	return newIterableIntersection[T](set, other)
}

func (set iterableUnion[T]) Complement() ImplicitSet[T] {
	return newComplement[T](set)
}

func (set iterableUnion[T]) IterateDistinct() (rangefunc func(yield func(element T) (next bool))) {
	return func(yield func(element T) (next bool)) {
		for i := 0; i < len(set.args); i++ {
			arg := set.args[i]

			arg.IterateDistinct()(func(element T) (next bool) {
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

func (set iterableUnion[T]) Size() int {
	size := 0

	set.IterateDistinct()(func(element T) (next bool) {
		size++

		return true
	})

	return size
}

func (set iterableUnion[T]) Evaluate() Set[T] {
	return evaluate[T](set)
}

func (set iterableUnion[T]) String() string {
	return toString[T](set)
}
