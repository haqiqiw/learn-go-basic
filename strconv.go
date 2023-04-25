package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println("Result:", boolean)
	} else {
		fmt.Println("Error:", err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println("Result:", number)
	} else {
		fmt.Println("Error:", err.Error())
	}

	value := strconv.FormatInt(1000000, 10)
	fmt.Println(value)

	valueInt, err := strconv.Atoi("2000000")
	if err == nil {
		fmt.Println("Result:", valueInt)
	} else {
		fmt.Println("Error:", err.Error())
	}
}
