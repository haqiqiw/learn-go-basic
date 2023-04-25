package main

import "fmt"

func endSystem() {
	message := recover()
	if message != nil {
		fmt.Println("Error:", message)
	}

	fmt.Println("End system")
}

func startSystem(error bool) {
	defer endSystem()

	if error {
		panic("ERROR")
	}

	// this wrong, put recover in defer func
	// message := recover()
	// fmt.Println("Error:", message)

	fmt.Println("Start system")
}

func main() {
	startSystem(true)
	fmt.Println("Main is running")
}
