package main

import (
	"fmt"
)

type T struct {
	a int
	b float64
	c string
}

// 自定义结构体输出格式
func (t *T) String() string {
	return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)

}

type S string

// 循环调用, 错误示例
func (s S) String() string {
	return fmt.Sprintf("S=%s", s)

}

// 正解
func (s S) String() string {
	return fmt.Sprintf("S=%s", string(s))

}

func main() {
	// test print functions
	t := &T{0, 100.9, "hello,\tgo"}
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%#v\n", t)

}
