package linkedListImpl

import "fmt"

type Item[T any] struct {
	Value T
	Prev  *Item[T]
	Next  *Item[T]
}

type LinkedList[T any] struct {
	Len  int
	Head *Item[T]
	Tail *Item[T]
}

func InitLinkedList() *LinkedList[any] {
	return &LinkedList[any]{}
}

func (l *LinkedList[any]) GetListLength() int {
	return l.Len
}

func (l *LinkedList[any]) GetFirstElemPointer() *Item[any] {
	return l.Head
}

func (l *LinkedList[any]) GetLastElemPointer() *Item[any] {
	return l.Tail
}

func (l *LinkedList[any]) AddItemToFront(v any) {
	item := Item[any]{Value: v}
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

func (l *LinkedList[any]) AddItemToBack(v any) {
	item := Item[any]{Value: v}
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

func (l *LinkedList[any]) RemoveItem(item Item[any]) {
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

func (l *LinkedList[any]) RemoveFrontItem() error {
	if l.Head == nil {
		return fmt.Errorf("removeFrontItem err: list is empty")
	}
	l.Head = l.Head.Next
	l.Len--
	return nil
}

func (l *LinkedList[any]) RemoveBackItem() error {
	if l.Head == nil {
		return fmt.Errorf("removeBackItem err: list is empty")
	}
	var prev *Item[any]
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

func (i Item[any]) GetItemValue() any {
	return i.Value
}

func (i Item[any]) GetNextItemPointer() *Item[any] {
	return i.Next
}

func (i Item[any]) GetPreviousItemPointer() *Item[any] {
	return i.Prev
}

func (l *LinkedList[any]) InsertAfterElem(element any, mark *Item[any]) error {
	if l.Head == nil {
		return fmt.Errorf("insertAfterElem err: list is empty")
	}
	if l.Head.Next == mark {
		z := l.Head.Next
		x := z.Next
		x.Prev = z
		x.Value = element
		l.Len++
		return nil
	} else {
		currentItem := l.Head.Next
		for currentItem != nil && currentItem.Next != nil && currentItem.Next != mark {
			currentItem = currentItem.Next
		}
		if currentItem == mark {
			z := currentItem.Next
			x := z.Next
			x.Prev = z
			x.Value = element
			l.Len++
			return nil
		}
	}
	return nil
}
