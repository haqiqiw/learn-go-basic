package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Windah Basudara", "Windah"))
	fmt.Println(strings.Split("Windah Basudara", " "))
	fmt.Println(strings.ToLower("Windah Basudara"))
	fmt.Println(strings.ToUpper("Windah Basudara"))
	fmt.Println(strings.ToTitle("windah basudara"))
	fmt.Println(strings.Trim("Windah Basudara  ", " "))
	fmt.Println(strings.ReplaceAll("Windah Basudara", "Windah", "Wawan"))
}
