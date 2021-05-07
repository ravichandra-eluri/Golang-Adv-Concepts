package main

import "fmt"

// Base struct
type Animal struct {
	Name string
}

func (a Animal) Breathe() string {
	return fmt.Sprintf("%s breathes", a.Name)
}

func (a Animal) Describe() string {
	return fmt.Sprintf("animal: %s", a.Name)
}

// Dog embeds Animal — promotes Breathe() and Describe()
type Dog struct {
	Animal
	Breed string
}

// Dog overrides Describe(), shadowing Animal.Describe()
func (d Dog) Describe() string {
	return fmt.Sprintf("dog: %s (%s)", d.Name, d.Breed)
}

func (d Dog) Fetch() string {
	return fmt.Sprintf("%s fetches!", d.Name)
}

// Interface embedding composes smaller interfaces
type Breather interface {
	Breathe() string
}

type Describer interface {
	Describe() string
}

type Living interface {
	Breather
	Describer
}

func introduce(l Living) {
	fmt.Println(l.Describe())
	fmt.Println(l.Breathe())
}

// Embedding enables field promotion too
type Address struct {
	City    string
	Country string
}

type Employee struct {
	Address        // City and Country promoted to Employee
	Name    string
	Role    string
}

func main() {
	d := Dog{Animal: Animal{Name: "Rex"}, Breed: "Labrador"}

	fmt.Println(d.Breathe())         // promoted from Animal
	fmt.Println(d.Fetch())           // Dog's own method
	fmt.Println(d.Describe())        // Dog's override
	fmt.Println(d.Animal.Describe()) // explicit parent call

	fmt.Println()
	introduce(d) // Dog satisfies the Living interface

	fmt.Println()
	e := Employee{
		Address: Address{City: "London", Country: "UK"},
		Name:    "Alice",
		Role:    "Engineer",
	}
	fmt.Println(e.Name, e.Role, e.City, e.Country) // City/Country promoted
}
