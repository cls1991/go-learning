package slices

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	// test array, slice
	array := [...]float64{10.9, 21.5, 12.4, 10}
	slice := []float64{10.9, 21.5, 12.4, 10}
	fmt.Printf("array: %v\n", array)
	fmt.Printf("sum of array is %f\n", Sum1(array))
	fmt.Println("sum of array is", Sum1(array))
	fmt.Printf("array: %v\n", array)
	fmt.Println("*********************************")
	fmt.Printf("array: %v\n", array)
	fmt.Printf("sum of array is %f\n", Sum2(&array))
	fmt.Println("sum of array is", Sum2(&array))
	fmt.Printf("array: %v\n", array)
	fmt.Println("*********************************")
	fmt.Printf("slice: %v\n", slice)
	fmt.Printf("sum of slice is %f\n", Sum3(slice))
	fmt.Println("sum of slice is", Sum3(slice))
	fmt.Printf("slice: %v\n", slice)
	fmt.Println("*********************************")

	// test slice append
	x := []int{1, 10, 2}
	x = append(x, 9, 7)
	fmt.Printf("%#v\n", x)
	fmt.Println(x)
	fmt.Println("********************")
	x2 := []string{"hello", "go"}
	y2 := []string{"go is nice", "hello, go"}
	x2 = append(x2, y2...)
	fmt.Printf("%#v\n", x2)
	fmt.Println(x2)
}
