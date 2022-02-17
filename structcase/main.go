package main

import "fmt"

//定义结构体
type customer struct {
	name string
	age  int
}

func main() {
	c := customer{"jone", 32}
	fmt.Println(c)
	p := &c
	p.age = 12
	fmt.Println(c)
	//访问结构体通常需要使用指针访问，但是这么写太啰嗦，所以语言允许我们隐式访问
	c.name = "java"
	fmt.Println(c)
}
