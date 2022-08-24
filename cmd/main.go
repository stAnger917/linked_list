package main

import (
	"fmt"
	l "linked_list/pkg/linkedListImpl"
)

func main() {
	fmt.Println("Let`s make our linked list!")
	list := l.InitLinkedList()
	list.AddItemToFront("Test value-1")
	list.AddItemToBack("Test value-2")
	fmt.Println(list.Head)
	fmt.Println(list.Tail)
	fmt.Println(list.Len)
}
