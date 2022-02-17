package main

import "fmt"

func main() {

	//defer在函数最后执行，类似栈，defer中的函数会立即求值

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	var x, y int
	x = 1
	y = 2

	defer fmt.Println(add(x, y))

	x = 2
	y = 3
}

func add(x, y int) int {
	return x + y
}
