package methods
//package main

import (
	//"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

/*
	Methods
 */

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(rate float64) {
	v.X *= rate
	v.Y *= rate
}

/*
	Functions equivalent
 */

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, rate float64) {
	v.X *= rate
	v.Y *= rate
}

// test case
//func main() {
//	// value receiver
//	v := Vertex{X: 10}
//	v.Scale(10)
//	fmt.Println(v.Abs(), Abs(v), v)
//
//	// pointer receiver
//	p := &Vertex{3, 4}
//	p.Scale(10)
//	fmt.Println(p.Abs(), Abs(*p), p)
//
//	// func
//	Scale(&v, 10)
//	fmt.Println(Abs(v), v)
//}
