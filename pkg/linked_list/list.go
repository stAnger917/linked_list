package linked_list

type Item[T any] struct {
	value     T
	prev      *Item[T]
	next      *Item[T]
	isDeleted bool
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
		if !h.isDeleted {
			result = append(result, h.value)
		}
	}
	return result
}

func (l *LinkedList[T]) SingleElementList(v T) *Item[T] {
	item := &Item[T]{
		value: v,
	}
	l.head = item
	l.tail = item
	l.len++
	return item
}

func (l *LinkedList[T]) AddItemToFront(v T) *Item[T] {
	if l.len == 0 {
		item := l.SingleElementList(v)
		return item
	}
	return l.InsertBeforeElem(v, l.head)
}

func (l *LinkedList[T]) AddItemToBack(v T) *Item[T] {
	if l.len == 0 {
		item := l.SingleElementList(v)
		return item
	}
	return l.InsertAfterElem(v, l.tail)
}

func (l *LinkedList[T]) RemoveItem(item *Item[T]) {
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
	item.isDeleted = true
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

func (l *LinkedList[T]) InsertAfterElem(element T, mark *Item[T]) *Item[T] {
	if mark.isDeleted {
		return nil
	}
	next := mark.next
	newItem := &Item[T]{
		value:     element,
		prev:      mark,
		next:      next,
		isDeleted: false,
	}
	mark.next = newItem
	if next != nil {
		next.prev = newItem
	} else {
		l.tail = newItem
	}
	l.len++
	return newItem
}

func (l *LinkedList[T]) InsertBeforeElem(element T, mark *Item[T]) *Item[T] {
	if mark.isDeleted {
		return nil
	}
	prev := mark.prev
	newItem := &Item[T]{
		value:     element,
		prev:      prev,
		next:      mark,
		isDeleted: false,
	}
	mark.prev = newItem
	if prev != nil {
		prev.next = newItem
	} else {
		l.head = newItem
	}
	l.len++
	return newItem
}

func (l *LinkedList[T]) RemoveFrontItem() error {
	if l.head == nil {
		return ErrNoItem
	}
	l.RemoveItem(l.head)
	return nil
}

func (l *LinkedList[T]) RemoveBackItem() error {
	if l.head == nil {
		return ErrNoItem
	}
	l.RemoveItem(l.tail)
	return nil
}
