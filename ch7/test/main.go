package main

import "fmt"

// Shape 定义一个接口
type Shape interface {
	Area() float64
}

type Type interface {
	Type() string
}

// Rectangle 定义一个结构体类型
type Rectangle struct {
	Width  float64
	Height float64
}

// Area 实现接口中的方法
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Type() string {
	return "长方形"
	panic("implement me")
}

func main() {
	// 创建一个 Rectangle 类型的对象
	rect := &Rectangle{Width: 10, Height: 5}

	// 将对象赋值给接口类型的变量
	var shape Shape
	shape = rect

	// 调用接口中的方法
	fmt.Println(shape.Area()) // 输出：50
	fmt.Println(rect.Area())  // 输出：50

	fmt.Println(rect.Type()) // 输出：50

	var s string = "hello, world"
	fmt.Println(len(s))     // "12"
	fmt.Println(s[0], s[7]) // "104 119" ('h' and 'w')
}
