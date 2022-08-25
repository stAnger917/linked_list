package main

import (
	"fmt"
	l "linked_list/pkg/linkedListImpl"
)

func main() {
	fmt.Println("Let`s make our linked list!")
	list := l.InitLinkedList[any]()
	list.AddItemToFront("Test value-1")
	list.AddItemToBack("Test value-2")
}
