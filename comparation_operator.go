package main

import "fmt"

func main() {
	var name1 = "Ayam"
	var name2 = "Ayam"

	var result bool = name1 == name2
	fmt.Println(result)

	amount1 := 100
	amount2 := 200
	fmt.Println(amount1 > amount2)
	fmt.Println(amount1 < amount2)
	fmt.Println(amount1 == amount2)
	fmt.Println(amount1 != amount2)
}
