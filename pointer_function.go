package main

import "fmt"

type Location struct {
	City, Province, Country string
}

func ChangeLocationToIndonesia(location Location) {
	location.Country = "Indonesia"
}

func ChangeLocationToSingapore(location *Location) {
	location.Country = "Singapore"
}

func main() {
	// pass by value
	location := Location{"Subang", "Jawa Barat", ""}

	ChangeLocationToIndonesia(location)

	fmt.Println(location) // address not changed

	// pass by reference
	location2 := &Location{"Subang", "Jawa Barat", ""}

	ChangeLocationToSingapore(location2)

	fmt.Println(location2) // address changed
}
