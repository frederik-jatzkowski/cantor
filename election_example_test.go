package cantor_test

import (
	"fmt"

	"github.com/frederik-jatzkowski/cantor"
	"github.com/frederik-jatzkowski/cantor/internal/testsuites"
)

// Set operations can be used to implement business logic.
// In this example, sets are used to control access during an election and keep track of people, who already voted.
func Example_election() {

	var (
		// Let's define some people.
		jeff    = testsuites.Person{Id: 1, Name: "Jeff", Age: 21}
		mary    = testsuites.Person{Id: 2, Name: "Mary", Age: 48}
		bob     = testsuites.Person{Id: 3, Name: "Bob", Age: 17}
		charles = testsuites.Person{Id: 4, Name: "Charles", Age: 19}

		// 3 of them are citizen.
		citizen = cantor.NewHashSet(jeff, mary, bob)

		// We keep track of people who already voted.
		hasVoted = cantor.NewHashSet[testsuites.Person]()

		// We define an ImplicitSet of all people, who are at least 18 years old.
		offAge = cantor.NewImplicitSet(testsuites.Person.IsOffAge)

		// We derive a set view with all people who can still vote.
		// Changes made to underlying sets are reflected in real time.
		canVote = citizen.Intersect(offAge).Intersect(hasVoted.Complement())

		// We still need a variable to collect all votes.
		votesInFavor uint
	)

	vote := func(person testsuites.Person, inFavor bool) error {
		if !canVote.Contains(person) {
			return fmt.Errorf("invalid vote by %v", person)
		}

		if inFavor {
			votesInFavor++
		}

		hasVoted.Add(person)

		return nil
	}

	// Mary and Jeff can without any issue.
	fmt.Println(vote(mary, true))
	fmt.Println(vote(jeff, false))

	// Charles is not a citizen and cannot vote
	fmt.Println(vote(charles, true))

	// Jeff cannot vote twice
	fmt.Println(vote(jeff, false))

	// Bob is too young and cannot vote.
	fmt.Println(vote(bob, false))

	fmt.Printf("%d out of %d valid votes were in favor\n", votesInFavor, hasVoted.Size())
	// Output:
	// <nil>
	// <nil>
	// invalid vote by {4 Charles 19}
	// invalid vote by {1 Jeff 21}
	// invalid vote by {3 Bob 17}
	// 1 out of 2 valid votes were in favor
}
