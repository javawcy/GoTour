package iterator

import "datastructurecase/stack"

//迭代器模式，解决内部切换数据接口导致外部crash
//对外隐藏数据存储实现提供迭代器
//例子：浏览器历史

//Iterator 迭代器接口
type Iterator[T any] interface {
	hasNext() bool
	next() T
}

//BrowseHistoryStackIterator 栈迭代器实现
type BrowseHistoryStackIterator[T any] struct {
	history *BrowseHistory[T]
	i       int
}

func (s *BrowseHistoryStackIterator[T]) hasNext() bool {
	return s.history.s.Size() > 0
}

func (s *BrowseHistoryStackIterator[T]) next() T {
	s.i++
	return s.history.s.Pop()
}

//BrowseHistory 历史记录
type BrowseHistory[T any] struct {
	s stack.Stack[T]
}

func (s *BrowseHistory[T]) push(t T) {
	s.s.Push(t)
}

func (s *BrowseHistory[T]) pop() T {
	return s.s.Pop()
}

//创建私有迭代器
func (s *BrowseHistory[T]) iterator() Iterator[T] {
	return &BrowseHistoryStackIterator[T]{
		history: s,
		i:       0,
	}
}

//NewBrowseHistory init method
func NewBrowseHistory[T any](cap int) *BrowseHistory[T] {
	return &BrowseHistory[T]{
		s: stack.NewArrayStack[T](cap),
	}
}
