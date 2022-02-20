package main

import (
	"fmt"
	"time"
)

//通过go关键字启动一个轻量化的线程

func sayHi() {
	fmt.Println("Hi")
}

func main() {
	go sayHi()

	time.Sleep(1 * time.Second)
	sayHi()
}
