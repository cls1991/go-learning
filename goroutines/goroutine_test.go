package goroutines

import (
	"time"
	"fmt"
	"testing"
)

func TestGoroutine(t *testing.T) {
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
