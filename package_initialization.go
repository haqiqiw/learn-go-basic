package main

import (
	"fmt"
	"learn-go-basic/database"
	_ "learn-go-basic/helper" // blank identifier, but this will call init
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
