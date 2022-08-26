package linkedListImpl

type Item[T any] struct {
	value T
	prev  *Item[T]
	next  *Item[T]
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
	result := make([]T, 0, l.len)
	for h := l.head; h != nil; h = h.next {
		result = append(result, h.value)
	}
	return result
}

func (l *LinkedList[T]) AddItemToFront(v T) Item[T] {
	item := Item[T]{value: v}
	temp := l.head
	item.next = temp
	l.head = &item
	if l.len == 0 {
		l.tail = l.head
	} else {
		temp.prev = &item
	}
	l.len++
	return item
}

func (l *LinkedList[T]) AddItemToBack(v T) Item[T] {
	item := Item[T]{value: v}
	temp := l.tail
	item.prev = temp
	l.tail = &item
	if l.len == 0 {
		l.head = l.tail
	} else {
		temp.next = &item
	}
	l.len++
	return item
}

func (l *LinkedList[T]) RemoveItem(item Item[T]) {
	prev := item.prev
	next := item.next
	if prev != nil {
		prev.next = next
	} else {
		l.head = next
	}
	if next != nil {
		next.prev = prev
	} else {
		l.tail = next
	}
	l.len--
	item.prev = nil
	item.next = nil
}

func (l *LinkedList[T]) RemoveFrontItem() error {
	if l.head == nil {
		return ErrNoItem
	}
	l.head = l.head.next
	l.len--
	return nil
}

func (l *LinkedList[T]) RemoveBackItem() error {
	if l.head == nil {
		return ErrNoItem
	}
	var prev *Item[T]
	current := l.head
	for current.next != nil {
		prev = current
		current = current.next
	}
	if prev != nil {
		prev.next = nil
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
	return i.next
}

func (i Item[T]) GetPreviousItemPointer() *Item[T] {
	return i.prev
}

func (l *LinkedList[T]) InsertAfterElem(element T, mark *Item[T]) error {
	if l.head == nil {
		return ErrNoItem
	}
	if l.head.next == mark {
		z := l.head.next
		x := z.next
		x.prev = z
		x.value = element
		l.len++
		return nil
	} else {
		currentItem := l.head.next
		for currentItem != nil && currentItem.next != nil && currentItem.next != mark {
			currentItem = currentItem.next
		}
		if currentItem == mark {
			z := currentItem.next
			x := z.next
			x.prev = z
			x.value = element
			l.len++
			return nil
		}
	}
	return nil
}
