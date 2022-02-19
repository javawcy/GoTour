package main

import "fmt"

//error 是一个内建函数，通过函数返回一个error类型返回值判断是否为nil判断函数是否出错

type MyError struct {
	msg string
}

func (e *MyError) Error() string {
	return e.msg
}

func errorTest() error {
	return &MyError{"some thing wrong"}
}

func main() {

	if err := errorTest(); err != nil {
		fmt.Printf("Error! msg = %s \n", err.Error())
	}

}
