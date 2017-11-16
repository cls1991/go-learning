package goroutines
//package main

import (
	"fmt"
	"time"
)

func main() {
	go say("golang", 1)
	say("hello", 3)

	tick := time.Tick(100 * time.Microsecond)
	boom := time.After(500 * time.Microsecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("     .")
			time.Sleep(50 * time.Microsecond)
		}
	}
}

func say(s string, t time.Duration) {
	for i := 0; i < 3 ; i++ {
		time.Sleep(100 * t * time.Microsecond)
		fmt.Println(s, ":", i)
	}
}
