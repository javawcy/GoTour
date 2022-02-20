package array

//实现动态数组
type Array[T any] interface {
	Add(element T)
	Get(i int) (T, error)
	Remove(i int) (int, error)
	Size() int
}

type ArrayList[T any] struct {
	elements []T
}

type ArrayListError struct {
	msg string
}

func (e *ArrayListError) Error() string {
	return e.msg
}

func checkIndex(index, lenght int, err error) error {
	if index < 0 || index >= lenght {
		err = &ArrayListError{"invaliable index"}
	}
	return err
}

func NewArrayList[T any](cap int) *ArrayList[T] {
	arr := ArrayList[T]{
		elements: make([]T, 0, cap),
	}
	return &arr
}

func (arr *ArrayList[T]) Add(element T) {
	arr.elements = append(arr.elements, element)
}

func (arr *ArrayList[T]) Get(i int) (T, error) {
	var err error
	checkIndex(i, arr.Size(), err)
	return arr.elements[i], err
}

func (arr *ArrayList[T]) Remove(i int) (int, error) {
	var err error
	checkIndex(i, arr.Size(), err)
	esL := arr.elements[:i]
	esR := arr.elements[i+1:]
	esL = append(esL, esR...)
	arr.elements = esL
	return i, err
}

func (arr *ArrayList[T]) Size() int {
	return len(arr.elements)
}
