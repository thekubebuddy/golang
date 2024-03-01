package main

import (
	"fmt"
)

func main() {
	p1 := Person{
		FirstName:  "John",
		LastName:   "Doe",
		CreditCard: "super secret",
	}

	// Marshiling a Personn struct to JSON
	p1.ShowPersonInfo()

	p1Json := p1.MarshalPerson()
	fmt.Println(p1Json)

	// Unmarshaling JSON to a Person struct from a JSON file
	p2 := UnmarshalPersonFromFile("./person.json")
	fmt.Printf("%+v", p2)

}
