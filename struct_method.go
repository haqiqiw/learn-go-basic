package main

import "fmt"

type User struct {
	Name, Address string
	Age           int
}

func (user User) sayHello() {
	fmt.Println("Hello, my name is", user.Name)
}

func (user User) sayHi() {
	fmt.Println("Hi, my name is", user.Name)
}

func main() {
	user := User{
		Name:    "Wawan",
		Address: "Pejaten",
		Age:     25,
	}
	user.sayHello()
	user.sayHi()
}
