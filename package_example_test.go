package cantor_test

import (
	"fmt"
	"strings"

	"github.com/frederik-jatzkowski/cantor"
)

// In cantor, a set can be derived from one or more other sets.
// All methods on a derived set are computed just in time and reflect changes made to the underlying sets.
// This allows you to define views on changing data, which are composable and usually very performant.
// Sometimes, it might be good to evaluate such a derived set into a new Set.
// During such evaluation, no intermediate sets must be stored, making the evaluation highly performant
// and avoiding pressure on the garbage collector.
func Example_derivedSets() {
	birds := cantor.NewHashSet("eagle", "pigeon", "duck", "swan")
	mammals := cantor.NewHashSet("lion", "pig", "tiger", "giraffe")
	fishes := cantor.NewHashSet("shark", "barracuda", "goldfish", "guppy")

	// We derive a set animals which contains all given birds, mammals and fishes.
	animals := birds.Union(mammals).Union(fishes)

	// In this example, we want to search for animals using a search term.
	search := ""

	// We define an ImplicitSet which contains all strings that start with the search term.
	startingWithSearchTerm := cantor.NewImplicitSet(func(element string) bool {
		return strings.HasPrefix(element, search)
	})

	// We derive a new set, which filters the animals by the search term.
	searchResult := animals.Intersect(startingWithSearchTerm)

	// Lets search for "pig".
	search = "pig"

	fmt.Println(searchResult) // {pigeon, pig}

	// Since result.String() is evaluated just in time, changes in the search term are reflected in our derived set.
	search = "g"

	fmt.Println(searchResult) // {giraffe, goldfish, guppy}

	// We can evaluate the result into a new and independent Set.
	evaluated := searchResult.Evaluate()

	// Changes in the underlying sets of searchResult do not influence this new set.
	// Additionally, operations on the evaluated set are more performant.
	mammals.Add("golden retriever")

	fmt.Println(searchResult) // {giraffe, golden retriever, goldfish, guppy}
	fmt.Println(evaluated)    // {giraffe, goldfish, guppy}
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
