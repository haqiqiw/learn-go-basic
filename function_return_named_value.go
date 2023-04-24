package main

import "fmt"

func getCompletename() (firstName string, middleName string, lastName string) {
	firstName = "Windah"
	middleName = "Basudara"
	lastName = "Gantenk"

	// return -> or use this more simple
	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getCompletename()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)

	firstValue, _, _ := getCompletename()
	fmt.Println(firstValue)
}
