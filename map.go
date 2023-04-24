package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Windah",
		"address": "Bekasi",
	}

	person["country"] = "Indonesia"

	fmt.Println(person)
	fmt.Println(len(person))
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["country"])

	delete(person, "country")
	fmt.Println(person)
	fmt.Println(len(person))
}
