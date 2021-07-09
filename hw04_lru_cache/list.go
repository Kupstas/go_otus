package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	size int
	head *ListItem
	tail *ListItem
}

func (l *list) Len() int {
	return l.size
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	newElem := &ListItem{Value: v}
	head := l.Front()

	if head != nil {
		newElem.Next = head
		head.Prev = newElem
	} else {
		l.tail = newElem
	}

	l.head = newElem
	l.size++

	return newElem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newElem := &ListItem{Value: v}
	tail := l.Back()

	if tail != nil {
		newElem.Prev = tail
		tail.Next = newElem
	} else {
		l.head = newElem
	}

	l.tail = newElem
	l.size++

	return newElem
}

func (l *list) Remove(i *ListItem) {
	prev := i.Prev
	next := i.Next

	if next != nil {
		next.Prev = prev
	}

	if prev != nil {
		prev.Next = next
	}

	l.size--
}

func (l *list) MoveToFront(i *ListItem) {
	if i != l.Front() {
		l.Remove(i)
		l.PushFront(i.Value)
	}
}

func NewList() List {
	return new(list)
}
