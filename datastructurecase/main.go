package main

import (
	"datastructurecase/array"
	"fmt"
)

func main() {

	//数组
	arr := array.NewArrayList[int](3)
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

	//栈

	//队列
}
