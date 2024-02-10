package cantor_test

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/frederik-jatzkowski/cantor"
)

// In cantor, a set can be derived from one or more other sets, eg. the union of multiple sets.
// Without explicit evaluation, all method results are computed on demand and reflect changes made to their arguments.
// This allows you to define views (somewhat similar to an SQL VIEW) on changing data,
// which are composable and usually very performant.
// Sometimes, it might be good to evaluate such a view into a new ExplicitSet,
// eg. for performance improvements or to decouple it from its arguments.
func Example_lazyEvaluation() {
	birds := cantor.NewExplicitSet("eagle", "pigeon", "duck", "swan")
	mammals := cantor.NewExplicitSet("lion", "pig", "tiger", "giraffe")
	fishes := cantor.NewExplicitSet("shark", "barracuda", "goldfish", "guppy")

	animals := birds.Union(mammals).Union(fishes)

	// In this example, we want to search for animals using a search term.
	search := ""

	// We define an ImplicitSet which contains all strings that start with the search term.
	startingWithSearchTerm := cantor.NewImplicitSet(func(element string) bool {
		return strings.HasPrefix(element, search)
	})

	// We derive a new set, which filters the elements by the search term.
	result := animals.Intersect(startingWithSearchTerm)

	// Lets search for "pig".
	search = "pig"

	fmt.Println(result) // {pigeon, pig}

	// Since result.String() is evaluated on demand, changes in the search term are reflected in our derived set.
	search = "g"

	fmt.Println(result) // {giraffe, goldfish, guppy}

	// We can evaluate the result into a new ExplicitSet.
	evaluated := result.Evaluate()

	// Changes in the search term do not influence this new set.
	search = "s"

	fmt.Println(result)    // {swan, shark}
	fmt.Println(evaluated) // {giraffe, goldfish, guppy}
}

func BenchmarkIterableSet_Contains(b *testing.B) {
	set := buildUnionOfIntersectionsOfDifferences(2, 2, 100000)
	random := rand.Int()

	b.ResetTimer()

	// this benchmark should not exceed on a modern CPU:
	// 50 ns/op
	// 0 B/op
	// 0 allocs/op
	for i := 0; i < b.N; i++ {
		set.Contains(random)
	}
}

func BenchmarkIterableSet_Size(b *testing.B) {
	set := buildUnionOfIntersectionsOfDifferences(2, 2, 100000)

	b.ResetTimer()

	// this benchmark should not exceed on a modern CPU:
	// 25 ms/op
	// 1000 B/op
	// 20 allocs/op
	for i := 0; i < b.N; i++ {
		set.Size()
	}
}

func BenchmarkIterableSet_Iter(b *testing.B) {
	set := buildUnionOfIntersectionsOfDifferences(2, 2, 100000)

	b.ResetTimer()

	// this benchmark should not exceed on a modern CPU:
	// 30 ms/op
	// 1000 B/op
	// 20 allocs/op
	for i := 0; i < b.N; i++ {
		set.Iter()(func(element int) (next bool) {
			return true
		})
	}
}

func BenchmarkIterableSet_Evaluate(b *testing.B) {
	set := buildUnionOfIntersectionsOfDifferences(2, 2, 100000)

	b.ResetTimer()

	// this benchmark should not exceed on a modern CPU:
	// 30 ms/op
	// 800000 B/op
	// 750 allocs/op
	for i := 0; i < b.N; i++ {
		set.Evaluate()
	}
}

// This function builds a fairly complicated expression of set operations which is build the following way:
// The union of numberOfIntersections many intersections of numberOfIntersections many differences of two sets each,
// which were constructed using numberOfRandomSamplesPerInput integers.
// This results in an expression with numberOfIntersections*numberOfDifferences many inputs of
// size <= numberOfRandomSamplesPerInput.
//
// This expression can be used for benchmarking of different operations on a lazy evaluated set.
func buildUnionOfIntersectionsOfDifferences(
	numberOfIntersections int,
	numberOfDifferences int,
	numberOfRandomSamplesPerInput int,
) cantor.IterableSet[int] {
	intersections := make([]cantor.IterableSet[int], 0, numberOfIntersections)

	for iIntersection := 0; iIntersection < numberOfIntersections; iIntersection++ {
		differences := make([]cantor.IterableSet[int], 0, numberOfDifferences)

		for iDifference := 0; iDifference < numberOfDifferences; iDifference++ {
			set1 := cantor.NewExplicitSet[int]()
			set2 := cantor.NewExplicitSet[int]()

			for iSample := 0; iSample < numberOfRandomSamplesPerInput; iSample++ {
				set1.Add(rand.Intn(5 * numberOfRandomSamplesPerInput))
				set2.Add(rand.Intn(5 * numberOfRandomSamplesPerInput))
			}

			differences = append(differences, set1.Intersect(set2.Complement()))
		}

		intersection := differences[0]
		differences = differences[1:]

		for _, difference := range differences {
			intersection = intersection.Intersect(difference)
		}

		intersections = append(intersections, intersection)
	}

	result := intersections[0]
	intersections = intersections[1:]

	for _, intersection := range intersections {
		result = result.Union(intersection)
	}

	return result
}
