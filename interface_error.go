package main

import (
	"errors"
	"fmt"
)

func Divided(value int, divide int) (int, error) {
	if divide == 0 {
		return 0, errors.New("Can't divide by 0")
	} else {
		return value / divide, nil
	}
}

func main() {
	fmt.Println(Divided(10, 2))
	fmt.Println(Divided(10, 0))

	result, err := Divided(10, 0)
	if err == nil {
		fmt.Println("Result:", result)
	} else {
		fmt.Println("Error:", err.Error())
	}
}
