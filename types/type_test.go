package types

import (
	"fmt"
	"testing"
)

func TestType(t *testing.T)  {
	var i interface{} = "golang"
	s := i.(string)
	fmt.Println(s)
	s, ok := i.(string)
	fmt.Println(s, ok)
	f, ok := i.(float64)
	fmt.Println(f, ok)
	// panic
	// panic: interface conversion: interface {} is string, not float64
	//f = i.(float64)
	//fmt.Println(f)
	do(12)
	do("hello")
	do(false)
}
