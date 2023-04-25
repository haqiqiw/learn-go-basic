package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	ups := Ups(1)
	fmt.Println(ups)

	fmt.Println(Ups(2))
	fmt.Println(Ups(3))
}
