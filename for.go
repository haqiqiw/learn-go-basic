package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("counter:", counter)
		counter++
	}

	for i := 1; i <= 10; i++ {
		fmt.Println("index:", i)
	}

	slice := []string{"Windah", "Basudara", "Brando"}

	for j := 0; j < len(slice); j++ {
		fmt.Println(slice[j])
	}

	for index, value := range slice {
		fmt.Println("index", index, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "Windah"
	person["title"] = "Youtuber"

	for key, value := range person {
		fmt.Println("key", key, "=", value)
	}
}
