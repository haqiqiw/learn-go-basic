package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`a([a-z])i`)

	fmt.Println(regex.MatchString("aqi"))
	fmt.Println(regex.MatchString("aki"))
	fmt.Println(regex.MatchString("aCi"))

	fmt.Println(regex.FindAllString("aci aqi aqu a1i", 10))
}
