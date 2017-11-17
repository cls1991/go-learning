package interfaces

import (
	"testing"
	"math"
	"fmt"
)

func TestInterface(t *testing.T) {
	var a AbSer
	// nil interface value
	// panic: runtime error: invalid memory address or nil pointer dereference
	//describe(a)
	//fmt.Println(a.Abs())
	f := MyFloat(-math.Sqrt2)
	var v *Vertex
	a = f
	describe(a)
	fmt.Println(a.Abs())
	a = v
	describe(a)
	fmt.Println(a.Abs())
	a = &Vertex{3, 4}
	describe(a)
	fmt.Println(a.Abs())
	// error
	// Vertex does not implement AbSer (Abs method has pointer receiver)
	//a = v
	//fmt.Println(a.Abs())
}
