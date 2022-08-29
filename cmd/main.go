package main

import (
	"fmt"
	l "github.com/stAnger917/linked_list/v2/pkg/linked_list"
	"log"
)

func main() {
	fmt.Println("Let`s make our linked list!")
	list := l.InitLinkedList[string]()
	list.AddItemToFront("hello")
	list.AddItemToBack("kek")
	e := list.AddItemToBack("world")
	err := list.InsertBeforeElem("wow", e)
	if err != nil {
		log.Fatal("blyat", err)
	}
	fmt.Println(list.ListToSlice())
}
