package cantor_test

import (
	"fmt"

	"github.com/frederik-jatzkowski/cantor"
)

// In cantor, a set can be derived from one or more other sets.
// All methods on a derived set are computed just in time and reflect changes made to the underlying sets.
// This allows you to define real time views on changing data, which are composable and usually very performant.
func ExampleDerivedSet() {
	var (
		birds                             = cantor.NewHashSet("eagle", "pigeon")
		mammals                           = cantor.NewHashSet("lion", "giraffe")
		fishes                            = cantor.NewHashSet("shark", "goldfish")
		animals cantor.DerivedSet[string] = birds.Union(mammals).Union(fishes)
	)

	fmt.Println(animals.Contains("dog")) // false

	mammals.Add("dog")

	// The derived set animals reflects the change of the underlying set mammals.
	fmt.Println(animals.Contains("dog")) // true
	// Output:
	// false
	// true
}

// Sometimes, it might be beneficial to evaluate such a DerivedSet into an independent Set.
// During such evaluation, no intermediate sets must be stored, making the evaluation highly performant
// and avoiding pressure on the garbage collector.
// This is possible due to evaluation of boolean expressions under the hood.
func ExampleDerivedSet_intoHashSet() {
	var (
		birds   = cantor.NewHashSet("eagle", "pigeon")
		mammals = cantor.NewHashSet("lion", "giraffe")
		fishes  = cantor.NewHashSet("shark", "goldfish")
		animals = birds.Union(mammals).Union(fishes).IntoHashSet()
	)

	fmt.Println(animals.Contains("dog")) // false

	mammals.Add("dog")

	// Since we evaluated all animals into a HashSet, changes to the set mammals are not reflected anymore.
	// This effectively took a snapshot at the time of the evaluation.
	fmt.Println(animals.Contains("dog")) // false
	// Output:
	// false
	// false
}

// Before the implementation of go rangefuncs, you can use a FunctionIterator like this.
// Afterwards, it can be used in native range loops.
func ExampleFunctionIterator() {
	set := cantor.NewHashSet(1, 2, 2, 3)
	iterate := set.Iterator()
	sum := 0

	iterate(func(element int) (next bool) {
		sum += element

		return true
	})

	fmt.Println(sum) // 6
	// Output:
	// 6
}

// Set operations can be used to calculate the delta of two datasets.
// The advantage of cantor is, that the sets representing the delta are derived and
// do have to store the elements themselves. This results in a minimal memory footprint and high performance.
func Example_deltaOfDatasets() {
	existing := cantor.NewHashSet("Alice", "Charles") // this might come from a database
	incoming := cantor.NewHashSet("Alice", "Bob")     // this might come from a rest api

	added := incoming.Intersect(existing.Complement())
	removed := existing.Intersect(incoming.Complement())
	unchanged := existing.Intersect(incoming)

	fmt.Println("added:", added)
	fmt.Println("removed:", removed)
	fmt.Println("unchanged:", unchanged)

	// Output:
	// added: {Bob}
	// removed: {Charles}
	// unchanged: {Alice}
}
