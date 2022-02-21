package array

// Array 实现动态数组
type Array[T any] interface {
	Add(element T)
	Get(i int) (T, error)
	Remove(i int) (int, error)
	Size() int
}

type DynamicArray[T any] struct {
	elements []T
}

type DynamicArrayError struct {
	msg string
}

func (e *DynamicArrayError) Error() string {
	return e.msg
}

func checkIndex(index, length int, err error) error {
	if index < 0 || index >= length {
		err = &DynamicArrayError{"invariable index"}
	}
	return err
}

func NewArray[T any](cap int) Array[T] {
	arr := DynamicArray[T]{
		elements: make([]T, 0, cap),
	}
	return &arr
}

func (list *DynamicArray[T]) Add(element T) {
	list.elements = append(list.elements, element)
}

func (list *DynamicArray[T]) Get(i int) (T, error) {
	var err error
	err = checkIndex(i, list.Size(), err)
	return list.elements[i], err
}

func (list *DynamicArray[T]) Remove(i int) (int, error) {
	var err error
	err = checkIndex(i, list.Size(), err)
	esL := list.elements[:i]
	esR := list.elements[i+1:]
	esL = append(esL, esR...)
	list.elements = esL
	return i, err
}

func (list *DynamicArray[T]) Size() int {
	return len(list.elements)
}
