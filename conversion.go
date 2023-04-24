package main

import "fmt"

func main() {
	var amount32 int32 = 100000
	var amount64 int64 = int64(amount32)
	var amount8 int8 = int8(amount32) // integer overflow

	fmt.Println(amount32)
	fmt.Println(amount64)
	fmt.Println(amount8)

	name := "Ayam Bakar"
	firstIndex := name[0]
	firstString := string(firstIndex)

	fmt.Println(firstString)
}
