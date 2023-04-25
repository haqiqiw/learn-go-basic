package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var customer Customer
	customer.Name = "Windah Basudara"
	customer.Address = "Bekasi"
	customer.Age = 30

	fmt.Println(customer)
	fmt.Println(customer.Name)
	fmt.Println(customer.Address)
	fmt.Println(customer.Age)

	wawan := Customer{
		Name:    "Wawan",
		Address: "Pejaten",
		Age:     25,
	}
	fmt.Println(wawan)

	// not recommended
	agus := Customer{"Agus", "Cilandak", 20}
	fmt.Println(agus)
}
