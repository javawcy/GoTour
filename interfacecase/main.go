package main

import (
	"fmt"
)

//go中的结构都为隐式实现，即duck type,只要方法定义相同则认同为实现，无需显式的implements
type human interface {
	getName() string
	getAge() int
}

//鸭子类型使用
func getHumanInfo(h human) {
	if h == nil {
		fmt.Println("nil")
		return
	}
	fmt.Printf("human name is %s, age is %d \n", h.getName(), h.getAge())
}

//因为golang接口的隐式实现 空接口可以当作任何类型，类似Object
func printMsg(v interface{}) {
	fmt.Println(v)
}

type student struct {
	name string
	age  int
}

func (s *student) getName() string {
	return s.name
}

func (s *student) getAge() int {
	return s.age
}

func main() {
	s := student{"Jone", 12}
	getHumanInfo(&s)

	//接口传参为空时
	getHumanInfo(nil)

	printMsg(s)

	//使用接口类型断言判断鸭子类型
	var h interface{}
	h = &s
	if _, ok := h.(*student); ok {
		fmt.Println("h is student")
	}

	//使用switch进行类型判断
	switch h.(type) {
	case *student:
		fmt.Println("student")
	default:
		fmt.Println("unkonw")
	}

}
