package types
//package main

import (
	"fmt"
)

//func main()  {
//	var i interface{} = "golang"
//	s := i.(string)
//	fmt.Println(s)
//	s, ok := i.(string)
//	fmt.Println(s, ok)
//	f, ok := i.(float64)
//	fmt.Println(f, ok)
//	// panic
//	// panic: interface conversion: interface {} is string, not float64
//	//f = i.(float64)
//	//fmt.Println(f)
//	do(12)
//	do("hello")
//	do(false)
//}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
