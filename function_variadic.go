package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(total)

	numbers := []int{10, 10, 10, 10, 10}
	total = sumAll(numbers...)
	fmt.Println(total)
}
