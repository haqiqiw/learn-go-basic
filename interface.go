package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayName(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

// using interface but not explicit
func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

// using interface but not explicit
func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Wawan"}
	sayName(person)

	cat := Animal{Name: "Cat"}
	sayName(cat)
}
