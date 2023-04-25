package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)
	fmt.Println(args[1])
	fmt.Println(args[2])

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname:", hostname)
	} else {
		fmt.Println("Error:", err.Error())
	}

	username := os.Getenv("APP_USERNAME")
	fmt.Println(username)
}
