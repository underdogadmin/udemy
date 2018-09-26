package main

import "fmt"

// Person - adding a comment so VS Code doesnt bark at me
type Person struct {
	Name string
	Age  int
}

// DoubleZero - adding a comment so VS Code doesnt bark at me
type DoubleZero struct {
	Person
	LicenseToKill bool
}

// Greeting Method - adding a comment so VS Code doesnt bark at me
func (p Person) Greeting() {
	fmt.Println("I'm just a regular person.")
}

// Greeting Method - adding a comment so VS Code doesnt bark at me
func (dz DoubleZero) Greeting() {
	fmt.Println("Miss Moneypenny, so good to see you.")
}

func main() {
	p1 := Person{
		Name: "Ian Fleming",
		Age:  44,
	}

	p2 := DoubleZero{
		Person: Person{
			Name: "James Bond",
			Age:  23,
		},
		LicenseToKill: true,
	}

	p1.Greeting()
	p2.Greeting()
	p2.Person.Greeting()
}
