package main

import (
	"fmt"
	"time"
)

//使用selector处理多个channel消息
//定义一个消费者，打印消息，接收信号关闭

func receiver(c, quit chan string) {
	for {
		select {
		case msg := <-c:
			fmt.Println(msg)
		case <-quit:
			fmt.Println("quit")
			break
		default:
			fmt.Println("no msg")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	c := make(chan string)
	quit := make(chan string)

	go receiver(c, quit)

	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		c <- "Hi"
	}
	quit <- ""
}
