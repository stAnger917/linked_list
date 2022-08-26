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
		l.tail = prev
	}
	l.len--
	item.prev = nil
	item.next = nil
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

func (l *LinkedList[T]) InsertBeforeElem(element T, mark *Item[T]) error {
	if l.head == nil {
		return ErrNoItem
	}
	if mark.next == nil {
		if mark.prev == l.tail.prev {
			l.addItemBefore(l.tail, element)
			return nil
		}
	}
	if l.head.next == mark.next {
		l.AddItemToFront(element)
	} else {
		currentItem := l.head.next
		for currentItem != nil && currentItem.prev != nil && currentItem.prev != mark {
			if currentItem.prev == mark.prev && currentItem.next == mark.next {
				l.addItemBefore(currentItem, element)
				return nil
			}
			currentItem = currentItem.next
		}
	}
	return nil
}

func (l *LinkedList[T]) addItemBefore(currentItem *Item[T], element T) {
	z := currentItem.prev
	newItem := &Item[T]{value: element}
	newItem.next = currentItem
	newItem.prev = z
	z.next = newItem
	currentItem.prev = newItem
	if currentItem.prev.prev != nil {
		z.prev = currentItem.prev.prev
	}
	l.len++
}
