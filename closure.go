package main

import "fmt"

func main() {
	counter := 0
	name := "Wawan"

	increment := func() {
		name := "Windah"
		fmt.Println("increment", counter)
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
