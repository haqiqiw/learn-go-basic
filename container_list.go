package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Wawan")
	data.PushBack("Knalpot")
	data.PushBack("Gantenk")
	data.PushFront("Mamank")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)
	fmt.Println(data.Front().Next().Next().Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
