package main

import "fmt"

func getFullname() (string, string, string) {
	return "Windah", "Basudara", "Gantenk"
}

func main() {
	firstName, middleName, lastName := getFullname()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)

	firstValue, _, _ := getFullname()
	fmt.Println(firstValue)
}
