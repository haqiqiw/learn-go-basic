package main

import "fmt"

func main() {
	var name string

	name = "Ayam Goreng"
	fmt.Println(name)

	name = "Ayam Bakar"
	fmt.Println(name)

	var aliasName = "Bebek Goreng"
	fmt.Println(aliasName)

	var age = 30
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)

	country = "Singapore"
	fmt.Println(country)

	// Multiple variable
	var (
		firstName = "Ayam"
		lastName  = "Bakar"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
