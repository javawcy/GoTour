package main

import (
	"fmt"
)

//go中没有类的概念，只有结构体，可以为结构体实现方法
//方法本质就是一个带有接收者的函数
type user struct {
	name string
	age  int
}

func (u user) getName() string {
	return u.name
}

func (u user) getAge() int {
	return u.age
}

//如果要通过方法修改接收者的值，应该用指针作为接收者, 否则不会被修改
func (u user) setAge(age int) {
	u.age = age
}
func (p *user) setName(name string) {
	p.name = name
}

//可以为非结构的其他类型声明方法,但不能是内建类型
type mystring string

func (s mystring) toArray() []mystring {
	arr := []mystring{s}
	return arr
}

func main() {
	u := user{"jone", 18}
	fmt.Println(u.getName())
	fmt.Println(u.getAge())
	u.setAge(20)
	fmt.Println(u)
	u.setName("Tom")
	fmt.Println(u)

	var s mystring
	s = "s"
	fmt.Println(s.toArray())

	//这里有一个语法糖，当调用者本身是值或指针，都可以调用任意接收者为指针和值的方法，因为go在编译时会自动为其重定向
	//但是指针和值作为参数传参必须严格指明

	//go中参数若为指针则是引用传递，为值则为值传递，如果需要对其进行同步修改需要为指针，否则会将该值复制一份，修改其内容不会对原值有影响
}
