package queue

import "fmt"

//Queue IDL
type Queue[T any] interface {
	EnQueue(t T) error
	DeQueue() T
	Peek() T
	Full() bool
	Size() int
	Print()
}

//ArrayQueueError for ArrayQueue
type ArrayQueueError struct {
	msg string
}

func (q *ArrayQueueError) Error() string {
	return q.msg
}

//ArrayQueue impl by array
type ArrayQueue[T any] struct {
	elements []T
	cap      int
	size     int
	front    int
	rear     int
}

//NewArrayQueue init method
func NewArrayQueue[T any](cap int) Queue[T] {
	return &ArrayQueue[T]{
		elements: make([]T, cap),
		cap:      cap,
		size:     0,
		front:    0,
		rear:     0,
	}
}

func (q *ArrayQueue[T]) Size() int {
	return q.size
}

func (q *ArrayQueue[T]) Full() bool {
	return q.size == q.cap
}

func (q *ArrayQueue[T]) EnQueue(t T) error {
	var err error
	if q.size == q.cap {
		err = &ArrayQueueError{"Queue is full"}
	} else {
		q.elements[q.rear] = t
		q.rear = (q.rear + 1) % q.cap
		q.size++
	}
	return err
}

func (q *ArrayQueue[T]) DeQueue() T {
	var t T
	if q.size > 0 {
		t = q.elements[q.front]
		q.front = (q.front + 1) % q.cap
		q.size--
	}
	return t
}

func (q *ArrayQueue[T]) Peek() T {
	var t T
	if q.size > 0 {
		t = q.elements[q.front]
	}
	return t
}

func (q *ArrayQueue[T]) Print() {
	fmt.Println(q.elements)
}
