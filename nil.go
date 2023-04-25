package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	newMap := NewMap("Wawan")
	fmt.Println(newMap)

	fmt.Println(NewMap(""))

	data := NewMap("")
	if data == nil {
		fmt.Println("Data is empty")
	} else {
		fmt.Println("Data is exist")
	}
}
