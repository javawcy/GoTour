package main

import "fmt"

func main() {

	//使用make函数创建一个map
	var m = make(map[string]int)
	m["java"] = 1
	fmt.Println(m)
	fmt.Println(m["java"])

	//初始化map值
	var m2 = map[string]int{
		"java":   1,
		"golang": 2,
	}
	fmt.Println(m2)

	//CRUD
	if elem, ok := m2["java"]; ok {
		fmt.Println(elem)
	}

	m2["golang"] = 3
	fmt.Println(m2)

	delete(m2, "java")
	fmt.Println(m2)

}
