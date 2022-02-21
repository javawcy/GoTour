package main

import (
	"datastructurecase/array"
	"datastructurecase/list"
	"fmt"
)

func main() {

	//数组
	arr := array.NewArray[int](3)
	arr.Add(1)
	arr.Add(2)
	arr.Add(3)
	arr.Add(4)
	arr.Add(5)
	arr.Add(6)
	if e, err := arr.Get(4); err == nil {
		fmt.Println(e)
	}
	if i, err := arr.Remove(4); err == nil {
		fmt.Println(i)
	}
	fmt.Println(arr)
	//链表
	list := list.NewLinkedList[int]()
	list.AddFirst(3)
	list.AddFirst(2)
	list.AddFirst(1)
	list.AddLast(4)
	list.AddLast(5)
	list.AddLast(6)
	if n, err := list.IndexOf(2); err == nil {
		fmt.Println(n)
	}
	if err := list.RemoveFirst(); err == nil {
		fmt.Println("delete one")
	}
	if err := list.RemoveLast(); err == nil {
		fmt.Println("delete one")
	}
	fmt.Println(list.Contains(2))
	fmt.Println(list.Size())
	//栈

	//队列
}
