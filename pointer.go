package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// pass by value
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)

	// pass by reference
	var address3 Address = Address{"Jakarta Selatan", "DKI Jakarat", "Indonesia"}
	var address4 *Address = &address3 // pointer
	var address5 *Address = &address3 // pointer

	address4.City = "Jakarta Pusat"

	// address3 will not changed by this
	// address4 = &Address{"Malang", "Jawa Timur", "Indonesia"}

	// address3 & address5 will changed by this
	// use * operator
	*address4 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address3)
	fmt.Println(address4)
	fmt.Println(address5)

	var address6 *Address = new(Address) // empty pointer
	address6.City = "Jakarta Barat"
	fmt.Println(address6)

	// new pointer
	var address7 *Address = &Address{"Surabaya", "Jawa Timur", "Indonesia"}
	fmt.Println(address7)
}
