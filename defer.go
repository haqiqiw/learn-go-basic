package main

import "fmt"

func logging() {
	fmt.Println("Logging")
}

func runApp(value int) {
	// this method will still executed
	// even though there is an error
	defer logging()

	fmt.Println("Run app")

	// trigger error
	result := 10 / value
	fmt.Println("Result", result)

	// this will not executed, because there is an error
	logging()
}

func main() {
	runApp(0)
}
