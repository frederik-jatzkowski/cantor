package cantor

type iterableUnion[T comparable] struct {
	args []IterableContainer[T]
}

func newIterableUnion[T comparable](args ...IterableContainer[T]) DerivedSet[T] {
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

func (set iterableUnion[T]) Union(other IterableContainer[T]) DerivedSet[T] {
	return newIterableUnion[T](append(set.args, other)...)
}

func (set iterableUnion[T]) Intersect(other Container[T]) DerivedSet[T] {
	return newIterableIntersection[T](set, other)
}

func (set iterableUnion[T]) Complement() ImplicitSet[T] {
	return newComplement[T](set)
}

func (set iterableUnion[T]) Iterator() FunctionIterator[T] {
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

func (set iterableUnion[T]) IntoHashSet() HashSet[T] {
	return evaluate[T](set)
}

func (set iterableUnion[T]) String() string {
	return toString[T](set)
}
