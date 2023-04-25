package main

import (
	"fmt"
	"sort"
)

type Account struct {
	Name string
	Age  int
}

type AccountSlice []Account

func (value AccountSlice) Len() int {
	return len(value)
}

func (value AccountSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value AccountSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	accounts := []Account{
		{"Wawan", 30},
		{"Windah", 35},
		{"Brando", 25},
		{"Mamank", 23},
	}

	fmt.Println(accounts)
	sort.Sort(AccountSlice(accounts))
	fmt.Println(accounts)
}
