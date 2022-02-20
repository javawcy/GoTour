package main

import (
	"fmt"
	"time"
)

// <- 箭头指向的位置为数据出口

func printChannelMsg(c chan string) {
	v := <-c
	fmt.Println(v)
}

func printChannelMsgByRange(c chan string) {
	for msg := range c {
		fmt.Println(msg)
	}
}

func main() {

	//使用内置的make创建channel
	c := make(chan string)

	go printChannelMsg(c)
	c <- "Hello World"
	close(c)

	time.Sleep(1 * time.Second)

	//使用带缓存的channel
	bc := make(chan string, 1)
	go printChannelMsg(bc)
	bc <- "Hello World"
	close(bc)

	time.Sleep(1 * time.Second)

	//使用range接收channel msg
	bc2 := make(chan string, 1)
	go printChannelMsgByRange(bc2)

	for {
		bc2 <- "Hi"
	}
}
