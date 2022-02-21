package list

import "reflect"

type List[T any] interface {
	AddFirst(t T)
	AddLast(t T)
	RemoveFirst() error
	RemoveLast() error
	Contains(t T) bool
	IndexOf(i int) (T, error)
	Size() int
}

type LinkedListError struct {
	msg string
}

func (err *LinkedListError) Error() string {
	return err.msg
}

type node[T any] struct {
	data T
	prev *node[T]
	next *node[T]
}

type LinkedList[T any] struct {
	length int
	head   *node[T]
	tail   *node[T]
}

func NewLinkedList[T any]() List[T] {
	return &LinkedList[T]{
		length: 0,
	}
}

func (l *LinkedList[T]) AddFirst(t T) {
	n := &node[T]{data: t}
	if l.length == 0 {
		l.head = n
		l.tail = n
	} else {
		l.head.prev = n
		n.next = l.head
		l.head = n
	}
	l.length++
}

func (l *LinkedList[T]) AddLast(t T) {
	n := &node[T]{data: t}
	if l.length == 0 {
		l.head = n
		l.tail = n
	} else {
		l.tail.next = n
		n.prev = l.tail
		l.tail = n
	}
	l.length++
}

func (l *LinkedList[T]) RemoveFirst() error {
	var err error
	if l.length == 0 {
		err = &LinkedListError{"have no data"}
	}

	if temp := l.head.next; temp != nil {
		temp.prev = nil
		l.head = temp
	} else {
		l.head = nil
		l.tail = nil
	}
	l.length--
	return err
}

func (l *LinkedList[T]) RemoveLast() error {
	var err error
	if l.length == 0 {
		err = &LinkedListError{"have no data"}
	}

	if temp := l.tail.prev; temp != nil {
		temp.next = nil
		l.tail = temp
	} else {
		l.tail = nil
		l.head = nil
	}
	l.length--
	return err
}

func (l *LinkedList[T]) Contains(t T) bool {
	var result bool = false
	temp := l.head
	for temp != nil {
		if reflect.DeepEqual(temp.data, t) {
			result = true
			break
		}
		temp = temp.next
	}
	return result
}

func (l *LinkedList[T]) IndexOf(i int) (T, error) {
	var err error
	if i < 0 || i >= l.length {
		err = &LinkedListError{"out of index"}
	}
	var result T
	temp := l.head
	index := 0
	for temp != nil {
		if index == i {
			result = temp.data
			break
		}
		index++
		temp = temp.next
	}
	return result, err
}

func (l *LinkedList[T]) Size() int {
	return l.length
}
