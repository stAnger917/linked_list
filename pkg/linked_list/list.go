package linked_list

import "fmt"

type Item struct {
	Value Value
	Prev  *Item
	Next  *Item
}

type Value struct {
	Data any
}

type LinkedList struct {
	Len  int
	Head *Item
	Tail *Item
}

func InitLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) GetListLength() int {
	return l.Len
}

func (l *LinkedList) GetFirstElemPointer() *Item {
	return l.Head
}

func (l *LinkedList) GetLastElemPointer() *Item {
	return l.Tail
}

func (l *LinkedList) AddItemToFront(v any) {
	item := Item{Value: Value{Data: v}}
	temp := l.GetFirstElemPointer()
	item.Next = temp
	l.Head = &item
	if l.Len == 0 {
		l.Tail = l.Head
	} else {
		temp.Prev = &item
	}
	l.Len++
}

func (l *LinkedList) AddItemToBack(v any) {
	item := Item{Value: Value{
		Data: v,
	}}
	temp := l.Tail
	item.Prev = temp
	l.Tail = &item
	if l.Len == 0 {
		l.Head = l.Tail
	} else {
		temp.Next = &item
	}
	l.Len++
}

func (l *LinkedList) RemoveItem(item Item) {
	prev := item.Prev
	next := item.Next
	if prev != nil {
		prev.Next = next
	} else {
		l.Head = next
	}
	if next != nil {
		next.Prev = prev
	} else {
		l.Tail = prev
	}
	l.Len--
	item.Prev = nil
	item.Next = nil
}

func (l *LinkedList) RemoveFrontItem() error {
	if l.Head == nil {
		return fmt.Errorf("removeFrontItem err: list is empty")
	}
	l.Head = l.Head.Next
	l.Len--
	return nil
}

func (l *LinkedList) RemoveBackItem() error {
	if l.Head == nil {
		return fmt.Errorf("removeBackItem err: list is empty")
	}
	var prev *Item
	current := l.Head
	for current.Next != nil {
		prev = current
		current = current.Next
	}
	if prev != nil {
		prev.Next = nil
	} else {
		l.Head = nil
	}
	l.Len--
	return nil
}

func (i Item) GetItemValue() any {
	return i.Value
}

func (i Item) GetNextItemPointer() *Item {
	return i.Next
}

func (i Item) GetPreviousItemPointer() *Item {
	return i.Prev
}
