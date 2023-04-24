package main

import "fmt"

func main() {
	name := "Windah"

	switch name {
	case "Windah":
		fmt.Println("Hello Windah")
		fmt.Println("Hello Basudara")
	case "Brando":
		fmt.Println("Hello Brando")
	default:
		fmt.Println("Hello Bang")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Name too long")
	case false:
		fmt.Println("Valid name")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Name too long")
	case length > 5:
		fmt.Println("Name a bit long")
	default:
		fmt.Println("Valid name")
	}

}
