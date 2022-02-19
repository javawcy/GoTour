package main

import (
	"fmt"
	"strings"
)

func main() {

	//切片类型定义为 []T

	var s []string
	var arr = [3]string{"a", "b", "c"}
	s = arr[0:1]

	fmt.Println(s)

	//切片就像数组的引用，修改切片中的内容会修改其底层数组
	s[0] = "d"
	fmt.Println(s)
	fmt.Println(arr)

	//切片上下界可以忽略。默认 0：lenght 前开后闭
	s2 := arr[:]
	fmt.Println(s2)

	//切片可以像数组一样初始化，其原理也是先构建数组再构建一个引用了数组的切片
	var sl = []string{"a", "b", "c"}
	fmt.Println(sl)

	//切片长度=元素个数，容量=从第一个元素到底层数组最后一个元素的个数
	var arr3 = [4]int{1, 2, 3, 4}
	s3 := arr3[0:1]
	fmt.Printf("len=%d,cap=%d \n", len(s3), cap(s3))

	s4 := arr3[1:2]
	fmt.Printf("len=%d,cap=%d \n", len(s4), cap(s4))

	//使用内置的make函数创建切片, 切片中每个元素会被赋予默认值
	s5 := make([]int, 5)
	fmt.Printf("slice=%v,len=%d,cap=%d \n", s5, len(s5), cap(s5))

	//切片可以包含任意类型，包含切片
	board := [][]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	//底层提供了一个可以向切片中添加元素的函数append，当底层函数不满足时会自动重新创建数组并将切片指向新数组
	board = append(board, []string{"-", "-", "-"})
	fmt.Println(board)

	//使用range可以遍历集合
	for i, v := range board {
		fmt.Printf("index=%d, value=%v", i, v)
	}

}
