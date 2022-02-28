package queue

import (
	"reflect"
	"testing"
)

func TestNewPriorityIntQueue(t *testing.T) {
	type args struct {
		cap int
	}
	tests := []struct {
		name string
		args args
		want Queue[int]
	}{
		{"priority1", args{5}, &PriorityIntQueue[int]{make([]int, 5), 5, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPriorityIntQueue(tt.args.cap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPriorityIntQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPriorityIDL(t *testing.T) {
	var p Queue[int]
	p = NewPriorityIntQueue(5)
	p.EnQueue(6)
	p.EnQueue(3)
	p.EnQueue(4)
	p.EnQueue(2)
	p.EnQueue(1)
	if p.Size() != 5 {
		t.Error("Piority size control error")
	}
	if !p.Full() {
		t.Error("Piority full control error")
	}
	if p.DeQueue() != 1 {
		t.Error("Piority dequeue error")
	}
	if p.Peek() != 2 {
		t.Error("Priority peek error")
	}
	p.Print()
}
