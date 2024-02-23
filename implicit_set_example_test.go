package cantor_test

import (
	"fmt"

	"github.com/frederik-jatzkowski/cantor"
)

// First, we have a set of all our citizen.
var citizen = cantor.NewHashSet(Person{"Jeff", 21}, Person{"Mary", 48}, Person{"Bob", 17})

// Then, we use a set to keep track of all people who already voted, which is initially empty.
var hasVoted = cantor.NewHashSet[Person]()

// We define an ImplicitSet of all people, who are at least 18 years old.
var offAge = cantor.NewImplicitSet(func(person Person) bool {
	return person.Age >= 18
})

// Finally, we can derive a set with all people who can vote.
// This DerivedSet will act as a real time view and reflect changes made to underlying sets.
var canVote = citizen.Intersect(offAge).Intersect(hasVoted.Complement())

func vote(person Person, inFavor bool) error {
	if !canVote.Contains(person) {
		return fmt.Errorf("invalid vote by %v", person)
	}

	if inFavor {
		votesInFavor++
	}

	hasVoted.Add(person)

	return nil
}

// Set operations can be used to implement business logic.
// In this example, sets are used to control access during an election and keep track of people, who already voted.
func Example_accessControl() {
	// Mary and Jeff can without any issue.
	fmt.Println(vote(Person{"Mary", 48}, true))
	fmt.Println(vote(Person{"Jeff", 21}, false))

	// Charles is not a citizen and cannot vote
	fmt.Println(vote(Person{"Charles", 19}, true))

	// Jeff cannot vote twice
	fmt.Println(vote(Person{"Jeff", 21}, false))

	// Bob is too young and cannot vote.
	fmt.Println(vote(Person{"Bob", 17}, false))

	fmt.Printf("%d out of %d valid votes were in favor\n", votesInFavor, hasVoted.Size())
	// Output:
	// <nil>
	// <nil>
	// invalid vote by {Charles 19}
	// invalid vote by {Jeff 21}
	// invalid vote by {Bob 17}
	// 1 out of 2 valid votes were in favor
}

type Person struct {
	Name string
	Age  uint8
}

// We still need a variable to collect all votes.
var votesInFavor uint
