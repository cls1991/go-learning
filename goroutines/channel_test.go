package goroutines

import (
	"fmt"
	"testing"
)

func TestChannel(t *testing.T)  {
	// channel
	s := []int {1, 3, 5, 7, 9}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x + y)

	// range and close
	c2 := make(chan int, 10)
	go fibonacci(cap(c2), c2)
	for i := range c2 {
		fmt.Println(i)
	}

	// select
	c3 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c3)
		}
		quit <- 0
	}()
	fibonacci2(c3, quit)
}
