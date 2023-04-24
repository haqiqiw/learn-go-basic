package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Ayam"
	names[1] = "Bakar"
	names[2] = "Enak"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	values := [3]int{
		90,
		80,
		70,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(names))
	fmt.Println(len(values))
}
