package main

import "fmt"

func main() {

	fmt.Println(add(1, 2))
	fmt.Println(addReturnNamed(1, 2))
	fmt.Println(swap("Hello", "World"))
}

func add(x, y int) int {
	return x + y
}

func addReturnNamed(x, y int) (z int) {
	z = x + y
	return
}

func swap(x, y string) (string, string) {
	return y, x
}
