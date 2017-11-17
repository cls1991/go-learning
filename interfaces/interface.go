package interfaces

import (
	"fmt"
	"math"
)

type AbSer interface {
	Abs() (float64, error)
}

type MyFloat float64

func (f MyFloat) Abs() (float64, error) {
	if f < 0 {
		return float64(-f), nil
	}
	return float64(f), nil
}

type Vertex struct {
	X, Y float64
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot sqrt negative number: %v\n", float64(e))
}

func (v *Vertex) Abs() (float64, error) {
	if v == nil {
		return 0, nil
	} else if v.X < 0 {
		return 0, ErrNegativeSqrt(v.X)
	} else if v.Y < 0 {
		return 0, ErrNegativeSqrt(v.Y)
	}
	return math.Sqrt(v.X*v.X + v.Y*v.Y), nil
}

func describe(a AbSer) {
	fmt.Printf("(%v, %T)\n", a, a)
}
