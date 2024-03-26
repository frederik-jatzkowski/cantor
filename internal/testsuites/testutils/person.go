package testutils

type Person struct {
	Id   uint
	Name string
	Age  uint
}

func (person Person) IsOffAge() bool {
	return person.Age >= 18
}
