package goroutines

import (
	"fmt"
	"time"
)

func say(s string, t time.Duration) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * t * time.Microsecond)
		fmt.Println(s, ":", i)
	}
}
