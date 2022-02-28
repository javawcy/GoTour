package stack

import (
	"reflect"
	"testing"
)

func TestNewArrayStack(t *testing.T) {
	type args struct {
		cap int
	}
	tests := []struct {
		name string
		args args
		want *ArrayStack[int]
	}{
		{"stack1", args{1}, &ArrayStack[int]{elements: make([]int, 1), cap: 1, size: 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArrayStack[int](tt.args.cap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArrayStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStackIDL(t *testing.T) {
	var s *ArrayStack[int]
	s = NewArrayStack[int](5)
	var size int = 2
	s.Push(1)
	s.Push(2)
	if size != s.Size() {
		t.Error("Push test fail, size error")
	}
	s.Push(3)
	s.Push(4)
	s.Push(5)
	if err := s.Push(6); err == nil {
		t.Error("Stack size control error")
	}
	if s.Pop() != 5 {
		t.Error("Stack Pop error")
	}
	if s.Peek() != 4 {
		t.Error("Stack Peek error")
	}
	if s.Full() {
		t.Error("Stack Full check error")
	}
}
