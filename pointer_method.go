package main

import "fmt"

type Man struct {
	Name string
}

func (man Man) Married() {
	man.Name = "Mr. " + man.Name
}

func (man *Man) Mamank() {
	man.Name = "Mamank. " + man.Name
}

func main() {
	// pass by value
	wawan := Man{"Wawan"}
	wawan.Married()
	fmt.Println(wawan)

	// pass by reference
	windah := Man{"Windah"}
	windah.Mamank()
	fmt.Println(windah)
}
