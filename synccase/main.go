package main

import (
	"fmt"
	"sync"
	"time"
)

type safeCounter struct {
	count int
	lock  sync.Mutex
}

func (c *safeCounter) incr() {
	c.lock.Lock()
	c.count++
	c.lock.Unlock()
}

func main() {

	c := safeCounter{count: 0}

	for i := 0; i < 1000; i++ {
		go c.incr()
	}

	time.Sleep(1 * time.Second)
	fmt.Printf("counter value is %d", c.count)

}
