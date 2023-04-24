package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello Bang"
	} else {
		return "Hello " + name
	}
}

func main() {
	hello := getHello("Windah")
	fmt.Println(hello)
	fmt.Println(getHello(""))
}
