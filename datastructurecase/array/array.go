package array

// Array 实现动态数组
type Array[T any] interface {
	Add(element T)
	Get(i int) (T, error)
	Remove(i int) (int, error)
	Size() int
}

type List[T any] struct {
	elements []T
}

type ListError struct {
	msg string
}

func (e *ListError) Error() string {
	return e.msg
}

func checkIndex(index, length int, err error) error {
	if index < 0 || index >= length {
		err = &ListError{"invariable index"}
	}
	return err
}

func NewArrayList[T any](cap int) *List[T] {
	arr := List[T]{
		elements: make([]T, 0, cap),
	}
	return &arr
}

func (list *List[T]) Add(element T) {
	list.elements = append[T](list.elements, element)
}

func (list *List[T]) Get(i int) (T, error) {
	var err error
	err = checkIndex(i, list.Size(), err)
	return list.elements[i], err
}

func (list *List[T]) Remove(i int) (int, error) {
	var err error
	err = checkIndex(i, list.Size(), err)
	esL := list.elements[:i]
	esR := list.elements[i+1:]
	esL = append[T](esL, esR...)
	list.elements = esL
	return i, err
}

func (list *List[T]) Size() int {
	return len[T](list.elements)
}
