package stack

//Stack IDL
type Stack[T any] interface {
	Push(t T) error
	Pop() T
	Peek() T
	Size() int
	Full() bool
}

//ArrayStackError Error for stack
type ArrayStackError struct {
	msg string
}

//Impl Error()
func (err *ArrayStackError) Error() string {
	return err.msg
}

//ArrayStack Use array impl stack
type ArrayStack[T any] struct {
	elements []T
	cap      int
	size     int
}

//NewArrayStack init method
func NewArrayStack[T any](cap int) *ArrayStack[T] {
	return &ArrayStack[T]{
		elements: make([]T, cap),
		cap:      cap,
		size:     0,
	}
}

//Impl Push
func (s *ArrayStack[T]) Push(t T) error {
	var err *ArrayStackError
	if s.size == s.cap {
		err = &ArrayStackError{"stack is full"}
	} else {
		s.elements[s.size] = t
		s.size++
	}
	return err
}

//Impl Pop
func (s *ArrayStack[T]) Pop() T {
	var t T
	if s.size > 0 {
		t = s.elements[s.size-1]
		s.size--
	}
	return t
}

//Impl Peek
func (s *ArrayStack[T]) Peek() T {
	var t T
	if s.size > 0 {
		t = s.elements[s.size-1]
	}
	return t
}

//Impl Size
func (s *ArrayStack[T]) Size() int {
	return s.size
}

//Impl Full
func (s *ArrayStack[T]) Full() bool {
	return s.size == s.cap
}
