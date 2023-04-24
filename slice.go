package main

import "fmt"

func main() {
	months := [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	slice1 := months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// change data
	// months[5] = "Hehe"
	// fmt.Println(slice1)

	slice2 := months[10:]
	fmt.Println(slice2)

	slice3 := append(slice2, "Hehe")
	fmt.Println(slice3)
	slice3[1] = "Not December"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Ayam"
	newSlice[1] = "Bakar"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	thisArray := [5]int{1, 2, 3, 4, 5}
	thisSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(thisArray)
	fmt.Println(thisSlice)
}
