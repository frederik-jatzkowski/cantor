package cantor

type union[T comparable] struct {
	args []ReadableSet[T]
}

func newUnion[T comparable](args ...ReadableSet[T]) ReadableSet[T] {
	return union[T]{
		args: args,
	}
}

func (set union[T]) Contains(element T) bool {
	for _, arg := range set.args {
		if arg.Contains(element) {
			return true
		}
	}

	return false
}

func (set union[T]) Union(other ReadableSet[T]) ReadableSet[T] {
	return newUnion[T](append(set.args, other)...)
}

func (set union[T]) Intersect(other Container[T]) ReadableSet[T] {
	return newIntersection[T](set, other)
}

func (set union[T]) Complement() ImplicitSet[T] {
	return NewImplicitSet(func(element T) bool {
		return !set.Contains(element)
	})
}

func (set union[T]) Elements() Iterator[T] {
	return func(yield func(element T) (next bool)) {
		for i := 0; i < len(set.args); i++ {
			arg := set.args[i]

			arg.Elements()(func(element T) (next bool) {
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

func (set union[T]) String() string {
	return toString[T](set)
}

func (set union[T]) Size() int {
	return count(set.Elements())
}
