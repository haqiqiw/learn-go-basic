package main

import "fmt"

func main() {
	var resultA = 80
	var resultB = 75

	var passedA = resultA >= 80
	var passedB = resultB >= 80
	fmt.Println(passedA)
	fmt.Println(passedB)

	var passed = passedA && passedB
	fmt.Println(passed)
}
