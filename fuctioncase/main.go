package main

import "fmt"

func main() {

	fmt.Println(add(1, 2))
	fmt.Println(addReturnNamed(1, 2))
	fmt.Println(swap("Hello", "World"))
	fmt.Println(sqrt(mulit, 2))

	adder1 := adder()
	adder2 := adder()
	fmt.Println(adder1(2))
	fmt.Println(adder2(2))
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

//函数可以当成参数传递和返回
func mulit(x, y int) int {
	return x * y
}

func sqrt(fn func(int, int) int, p int) int {
	return fn(p, p)
}

//由于函数本身可以当参数传递，此时函数中定义一个变量，这个变量相当于绑定在了函数中，形成了闭包 closure
func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}
