package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

type Circle struct {
	Radius int
}

type Shape interface {
	Area() int
	Perimeter() int
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() int {
	return 2*r.Width + 2*r.Height
}
func (c Circle) Area() int {
	return c.Radius * c.Radius
}
func (c Circle) Perimeter() int {
	return 2*c.Radius + 2*c.Radius
}

func main() {
	fmt.Println("面向对象 test")

	r := Rectangle{10, 20}
	c := Circle{10}
	s1 := Shape(r)
	s2 := Shape(c)

	fmt.Println(s1.Area())
	fmt.Println(s1.Perimeter())
	fmt.Println(s2.Area())
	fmt.Println(s2.Perimeter())

}
