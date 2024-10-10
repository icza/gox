package slicesx_test

import (
	"fmt"

	. "github.com/icza/gox/slicesx"
)

type Person struct {
	ID   int
	Name string
	Age  int
}

func (p Person) GetID() int  { return p.ID }
func (p Person) GetAge() int { return p.Age }

func PersonName(p Person) string { return p.Name }

func Example() {
	persons := []Person{
		{1, "Bob", 12},
		{2, "Alice", 23},
		{3, "Joe", 23},
	}

	// Get all names (using a named function)
	names := Props(persons, PersonName)
	fmt.Println("Names:", names)

	// Get an ID -> Person map (using a method)
	idPersons := PropMap(persons, Person.GetID)
	fmt.Println("ID -> Person map:", idPersons)

	// Get an Age -> []Person map (using a method)
	agePersons := PropsMap(persons, Person.GetAge)
	fmt.Println("Age -> []Person map:", agePersons)

	// Filter out children (using an anoymous function)
	adults := Filter(persons, func(p Person) bool { return p.Age >= 18 })
	fmt.Println("Adults:", adults)

	// Safely index:
	indexResult := Index(persons, 100, Person{9, "Default", 99})
	fmt.Println("Index result:", indexResult)

	// Output:
	// Names: [Bob Alice Joe]
	// ID -> Person map: map[1:{1 Bob 12} 2:{2 Alice 23} 3:{3 Joe 23}]
	// Age -> []Person map: map[12:[{1 Bob 12}] 23:[{2 Alice 23} {3 Joe 23}]]
	// Adults: [{2 Alice 23} {3 Joe 23}]
	// Index result: {9 Default 99}
}
