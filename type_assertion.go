package main

import (
	"fmt"
)

func random() interface{} {
	return "OK"
}

func main() {
	result := random()
	resultString := result.(string)
	fmt.Println(resultString)

	// error, panic
	// resultInt := result.(int)
	// fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("string:", value)
	case int:
		fmt.Println("int:", value)
	default:
		fmt.Println("unknown type")
	}
}
