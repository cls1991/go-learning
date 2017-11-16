package goroutines
//package main

//import "fmt"

//func main()  {
//	// channel
//	s := []int {1, 3, 5, 7, 9}
//	c := make(chan int)
//	go sum(s[:len(s)/2], c)
//	go sum(s[len(s)/2:], c)
//	x, y := <-c, <-c
//	fmt.Println(x, y, x + y)
//
//	// range and close
//	c2 := make(chan int, 10)
//	go fibonacci(cap(c2), c2)
//	for i := range c2 {
//		fmt.Println(i)
//	}
//
//	// select
//	c3 := make(chan int)
//	quit := make(chan int)
//	go func() {
//		for i := 0; i < 10; i++ {
//			fmt.Println(<-c3)
//		}
//		quit <- 0
//	}()
//	fibonacci2(c3, quit)
//}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	close(c)
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c<-x:
			x, y = y, x + y
		case <-quit:
			return
		}
	}
}
