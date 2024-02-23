[<- README](../../README.md#release-notes)

# Ideas for future releases

The following are a couple of ideas for the future of this package.

## Shorthands for more complex operations

Currently, only the three basic set operations `Union`, `Intersection` and `Complement` are implemented.
These can be composed to achieve all other set operations. For some of the most usefuly operations, shorthand methods could be implemented.

```go
set1.Difference(set2)
set1.SymmetricDifference(set2)
set.FilterBy(predicate)
```

## Equalities

Implementation of set equality operations:

```go
set1.Equals(set2)
set1.IsSubsetOf(set2)
set1.IsSubsetEqOf(set2)
set1.IsSupersetOf(set2)
set1.IsSupersetEqOf(set2)
```

## Projection to a comparable key

Sets should not be limited to `comparable` types. Instead, an interface 
```go
type Keyer[T comparable] interface {
    Key() T
}
```
could be implemented to allow sets over any data structure.
```go
type KeyedHashSet[K comparable, T Keyer[K]] struct {}
```
Such sets should also implement a `set.Get(key K) E` method to retrieve all underlying data.

## Ordered sets using balanced search trees

Sometimes, it might be beneficial to keep order information about the elements of a set. Thus, types like
```go
type TreeSet[K Ordered] struct {}
type KeyedTreeSet[K Ordered, T Keyer[K]] struct {}
```
could be implemented.