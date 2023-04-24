package main

import "fmt"

func main() {
	name := "Hehe"

	if name == "Windah" {
		fmt.Println("Hello Windah")
	} else if name == "Brando" {
		fmt.Println("Hello Brando")
	} else {
		fmt.Println("Hello Bang")
	}

	if length := len(name); length > 5 {
		fmt.Println("Name too long")
	} else {
		fmt.Println("Valid name")
	}

}
