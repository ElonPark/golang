package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter()	float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// Rect 타입에 대한 Shape 인터페이스 구현
func (receiver Rect) area() float64 {
	return receiver.width * receiver.height
}

func (receiver Rect) perimeter() float64 {
	return 2 * (receiver.width + receiver.height)
}

// Circle 타입에 대한 Shape 인터페이스 구현
func (receiver Circle) area() float64 {
	return math.Pi * receiver.radius * receiver.radius
}

func (receiver Circle) perimeter() float64 {
	return 2 * math.Pi * receiver.radius
}

func main() {
	r := Rect{10.0, 20.0}
	c := Circle{10}

	showArea(r, c)

}

func showArea(shapes ...Shape)  {
	for _, s := range shapes {
		a := s.area() // 인터페이스에서 메소드 호출
		fmt.Println(a)
	}
}
