package cantor

type union[T comparable] struct {
	args []IterableContainer[T]
}

func newUnion[T comparable](args ...IterableContainer[T]) DerivedSet[T] {
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

func (set union[T]) Union(other IterableContainer[T]) DerivedSet[T] {
	return newUnion[T](append(set.args, other)...)
}

func (set union[T]) Intersect(other Container[T]) DerivedSet[T] {
	return newIntersection[T](set, other)
}

func (set union[T]) Complement() ImplicitSet[T] {
	return newComplement[T](set)
}

func (set union[T]) Iterator() FunctionIterator[T] {
	return func(yield func(element T) (next bool)) {
		for i := 0; i < len(set.args); i++ {
			arg := set.args[i]

			arg.Iterator()(func(element T) (next bool) {
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

func (set union[T]) IntoHashSet() HashSet[T] {
	return evaluate[T](set)
}

func (set union[T]) String() string {
	return toString[T](set)
}
