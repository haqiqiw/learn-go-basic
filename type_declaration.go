package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpUser NoKTP = "12345678980"
	var marriedStatus Married = true
	fmt.Println(noKtpUser)
	fmt.Println(marriedStatus)
}
