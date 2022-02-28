package queue

import "fmt"

//PriorityIntQueueError for PriorityIntQueue
type PriorityIntQueueError struct {
	msg string
}

func (err *PriorityIntQueueError) Error() string {
	return err.msg
}

//PriorityIntQueue int 升序优先级队列
type PriorityIntQueue[T int] struct {
	elements []T
	cap      int
	size     int
}

//NewPriorityIntQueue init method
func NewPriorityIntQueue[T int](cap int) Queue[T] {
	return &PriorityIntQueue[T]{
		elements: make([]T, cap),
		size:     0,
		cap:      cap,
	}
}

func (q *PriorityIntQueue[T]) Full() bool {
	return q.size == q.cap
}

func (q *PriorityIntQueue[T]) Size() int {
	return q.size
}

func (q *PriorityIntQueue[T]) EnQueue(t T) error {
	var err error
	if q.size == q.cap {
		err = &PriorityIntQueueError{"queue is full"}
	}
	if q.size == 0 {
		q.elements[0] = t
	} else {
		for i := q.size - 1; i >= 0; i-- {
			if t >= q.elements[i] {
				q.elements[i+1] = t
				break
			} else {
				q.elements[i+1] = q.elements[i]
				if i == 0 {
					q.elements[i] = t
				}
			}
		}
	}
	q.size++
	return err
}

func (q *PriorityIntQueue[T]) DeQueue() T {
	var t T
	if q.size > 0 {
		t = q.elements[0]
		for i := 0; i < q.size-1; i++ {
			q.elements[i] = q.elements[i+1]
		}
		q.size--
	}
	return t
}

func (q *PriorityIntQueue[T]) Peek() T {
	var t T
	if q.size > 0 {
		t = q.elements[0]
	}
	return t
}

func (q *PriorityIntQueue[T]) Print() {
	fmt.Println(q.elements)
}
