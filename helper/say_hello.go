package helper

import "fmt"

var version = "1.0.0" // not accesible from outside
var Application = "learn-go-basic"

func init() {
	fmt.Println("Init helper")
}

func SayHello(name string) {
	fmt.Println("Hello", name)
}

// not accesible from outside
func sayGoodBye(name string) {
	fmt.Println("Good bye", name)
}
