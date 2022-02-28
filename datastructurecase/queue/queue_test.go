package queue

import (
	"reflect"
	"testing"
)

func TestNewArrayQueue(t *testing.T) {
	type args struct {
		cap int
	}
	tests := []struct {
		name string
		args args
		want Queue[int]
	}{
		{"queue1", args{3}, &ArrayQueue[int]{make([]int, 3), 3, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArrayQueue[int](tt.args.cap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArrayQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueueIDL(t *testing.T) {
	var q Queue[int]
	q = NewArrayQueue[int](5)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	if q.Size() != 4 {
		t.Error("Queue size error")
	}
	q.EnQueue(5)
	err := q.EnQueue(6)
	if err == nil {
		t.Error("Queue cap control error")
	}
	if q.DeQueue() != 1 {
		t.Error("Queue first error")
	}
	if q.DeQueue() != 2 {
		t.Error("Queue first error")
	}
	q.EnQueue(6)
	q.EnQueue(7)
	if !q.Full() {
		t.Error("Queue full check error")
	}
	if q.Peek() != 3 {
		t.Error("Queue peed error")
	}
	q.Print()
}
