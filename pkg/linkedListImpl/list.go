package linkedListImpl

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

func InitLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) GetListLength() int {
	return l.Len
}

func (l *LinkedList[T]) GetHead() (*Item[T], error) {
	if l.Head != nil {
		return l.Head, nil
	}
	return nil, ErrNilHead
}

func (l *LinkedList[T]) GetTail() (*Item[T], error) {
	if l.Tail != nil {
		return l.Tail, nil
	}
	return nil, ErrNilTail
}

func (l *LinkedList[T]) ListToSlice() []any {
	var result []any
	if l.Head == nil {
		return result
	}
	current := l.Head
	for current.Next != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	result = append(result, current)
	return result
}

func (l *LinkedList[T]) AddItemToFront(v T) {
	item := Item[T]{Value: v}
	temp, _ := l.GetHead()
	item.Next = temp
	l.Head = &item
	if l.Len == 0 {
		l.Tail = l.Head
	} else {
		temp.Prev = &item
	}
	l.Len++
}

func (l *LinkedList[T]) AddItemToBack(v T) {
	item := Item[T]{Value: v}
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

func (l *LinkedList[T]) RemoveItem(item Item[T]) {
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

func (l *LinkedList[T]) RemoveFrontItem() error {
	if l.Head == nil {
		return ErrNoItem
	}
	l.Head = l.Head.Next
	l.Len--
	return nil
}

func (l *LinkedList[T]) RemoveBackItem() error {
	if l.Head == nil {
		return ErrNoItem
	}
	var prev *Item[T]
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

func (i Item[T]) GetItemValue() T {
	return i.Value
}

func (i Item[T]) GetNextItemPointer() *Item[T] {
	return i.Next
}

func (i Item[T]) GetPreviousItemPointer() *Item[T] {
	return i.Prev
}

func (l *LinkedList[T]) InsertAfterElem(element T, mark *Item[T]) error {
	if l.Head == nil {
		return ErrNoItem
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
