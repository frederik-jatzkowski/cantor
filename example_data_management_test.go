package cantor_test

import (
	"fmt"

	"github.com/frederik-jatzkowski/cantor"
)

// Set operations can be used to calculate the delta of two datasets.
// The advantage of cantor is, that the sets representing the delta are derived and
// do have to store the elements themselves. This results in a minimal memory footprint and high performance.
func Example_dataSets() {
	existing := cantor.NewHashSet("Alice", "Charles") // this might come from a database
	incoming := cantor.NewHashSet("Alice", "Bob")     // this might come from a rest api

	added := incoming.Difference(existing)
	removed := existing.Difference(incoming)
	unchanged := existing.Intersect(incoming)

	fmt.Println("added:", added)
	fmt.Println("removed:", removed)
	fmt.Println("unchanged:", unchanged)
	// Output:
	// added: {Bob}
	// removed: {Charles}
	// unchanged: {Alice}
}
