package main

import (
	"fmt"
	l "linked_list/pkg/linkedListImpl"
)

func main() {
	fmt.Println("Let`s make our linked list!")
	list := l.InitLinkedList[string]()
	list.AddItemToBack("hello")
	list.AddItemToBack("world")
	e := list.AddItemToBack("kek")
	list.RemoveItem(e)

	list.AddItemToBack("kek")

	fmt.Println(list.ListToSlice())
}
