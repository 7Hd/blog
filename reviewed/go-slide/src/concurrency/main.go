package main

import (
	"fmt"
)

func goroutinefn() {
	go func() {
		// do something
	}()

	go func(i int) {
		// do something with param
	}(0)
}

func channel() {
	c := make(chan int)
	go func() {
		// channel 會 block 所以需要使用 goroutine
		c <- 1
	}()
	fmt.Println(<-c)
}

func main() {}
