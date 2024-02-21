package cantor_test

import (
	"math/rand"
	"testing"

	"github.com/frederik-jatzkowski/cantor"
)

func BenchmarkSet_Contains(b *testing.B) {
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

func BenchmarkSet_Iter(b *testing.B) {
	set := buildUnionOfIntersectionsOfDifferences(2, 2, 100000)

	b.ResetTimer()

	// this benchmark should not exceed on a modern CPU:
	// 30 ms/op
	// 1000 B/op
	// 20 allocs/op
	for i := 0; i < b.N; i++ {
		set.Iterator()(func(element int) (next bool) {
			return true
		})
	}
}

func BenchmarkSet_IntoHashSet(b *testing.B) {
	set := buildUnionOfIntersectionsOfDifferences(2, 2, 100000)

	b.ResetTimer()

	// this benchmark should not exceed on a modern CPU:
	// 30 ms/op
	// 800000 B/op
	// 750 allocs/op
	for i := 0; i < b.N; i++ {
		set.IntoHashSet()
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
) cantor.DerivedSet[int] {
	intersections := make([]cantor.DerivedSet[int], 0, numberOfIntersections)

	for iIntersection := 0; iIntersection < numberOfIntersections; iIntersection++ {
		differences := make([]cantor.DerivedSet[int], 0, numberOfDifferences)

		for iDifference := 0; iDifference < numberOfDifferences; iDifference++ {
			set1 := cantor.NewHashSet[int]()
			set2 := cantor.NewHashSet[int]()

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
