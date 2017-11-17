package methods

import (
	"fmt"
	"testing"
)

func TestMethod(t *testing.T) {
	// value receiver
	v := Vertex{X: 10}
	v.Scale(10)
	fmt.Println(v.Abs(), Abs(v), v)

	// pointer receiver
	p := &Vertex{3, 4}
	p.Scale(10)
	fmt.Println(p.Abs(), Abs(*p), p)

	// func
	Scale(&v, 10)
	fmt.Println(Abs(v), v)
}
