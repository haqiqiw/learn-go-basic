package main

import "fmt"

func endApp() {
	fmt.Println("End app")
}

func startApp(error bool) {
	defer endApp()

	if error {
		panic("ERROR")
	}
	fmt.Println("Start app")
}

func main() {
	startApp(true)
}
