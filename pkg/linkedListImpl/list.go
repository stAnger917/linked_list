package linkedListImpl

type Item[T any] struct {
	value T
	Prev  *Item[T]
	Next  *Item[T]
}

type LinkedList[T any] struct {
	len  int
	head *Item[T]
	tail *Item[T]
}

func InitLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) GetListLength() int {
	return l.len
}

func (l *LinkedList[T]) GetHead() *Item[T] {
	return l.head
}

func (l *LinkedList[T]) GetTail() *Item[T] {
	return l.tail
}

func (l *LinkedList[T]) ListToSlice() []T {
	result := make([]T, l.len)
	for h := l.head; h != nil; h = h.Next {
		result = append(result, h.value)
	}
	return result
}

func (l *LinkedList[T]) AddItemToFront(v T) {
	item := Item[T]{value: v}
	temp := l.head
	item.Next = temp
	l.head = &item
	if l.len == 0 {
		l.tail = l.head
	} else {
		temp.Prev = &item
	}
	l.len++
}

func (l *LinkedList[T]) AddItemToBack(v T) {
	item := Item[T]{value: v}
	temp := l.tail
	item.Prev = temp
	l.tail = &item
	if l.len == 0 {
		l.head = l.tail
	} else {
		temp.Next = &item
	}
	l.len++
}

func (l *LinkedList[T]) RemoveItem(item Item[T]) {
	prev := item.Prev
	next := item.Next
	if prev != nil {
		prev.Next = next
	} else {
		l.head = next
	}
	if next != nil {
		next.Prev = prev
	} else {
		l.tail = prev
	}
	l.len--
	item.Prev = nil
	item.Next = nil
}

func (l *LinkedList[T]) RemoveFrontItem() error {
	if l.head == nil {
		return ErrNoItem
	}
	l.head = l.head.Next
	l.len--
	return nil
}

func (l *LinkedList[T]) RemoveBackItem() error {
	if l.head == nil {
		return ErrNoItem
	}
	var prev *Item[T]
	current := l.head
	for current.Next != nil {
		prev = current
		current = current.Next
	}
	if prev != nil {
		prev.Next = nil
	} else {
		l.head = nil
	}
	l.len--
	return nil
}

func (i Item[T]) GetItemValue() T {
	return i.value
}

func (i Item[T]) GetNextItemPointer() *Item[T] {
	return i.Next
}

func (i Item[T]) GetPreviousItemPointer() *Item[T] {
	return i.Prev
}

func (l *LinkedList[T]) InsertAfterElem(element T, mark *Item[T]) error {
	if l.head == nil {
		return ErrNoItem
	}
	if l.head.Next == mark {
		z := l.head.Next
		x := z.Next
		x.Prev = z
		x.value = element
		l.len++
		return nil
	} else {
		currentItem := l.head.Next
		for currentItem != nil && currentItem.Next != nil && currentItem.Next != mark {
			currentItem = currentItem.Next
		}
		if currentItem == mark {
			z := currentItem.Next
			x := z.Next
			x.Prev = z
			x.value = element
			l.len++
			return nil
		}
	}
	return nil
}
