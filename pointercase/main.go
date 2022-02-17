package main

import "fmt"

func main() {

	// *T 是指向 T类型 的指针
	// &T 用于生成 T类型 的指针
	// *p 用于取指针指向T的value

	var p *int

	intValue := 1
	p = &intValue

	fmt.Println(p)
	fmt.Println(*p)

}
