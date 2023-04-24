package main

import "fmt"

func main() {
	const firstName string = "Ayam"
	const lastName = "Bakar"
	const age = 25

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)

	const (
		name     string = "Ayam"
		fullName        = "Ayam Bakar"
	)

	fmt.Println(name)
	fmt.Println(fullName)
}
