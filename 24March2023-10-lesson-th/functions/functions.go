package functions

import (
	"fmt"
	"math"
)

type Shape interface{
	Area() float64
	Perimeter() float64
}

type Circle struct{
	radius float64
}

type Triangle struct{
	height float64
	side float64
}

type Rectangle struct{
	height float64
	width float64
}

func (c Circle) Perimeter() float64 {
	fmt.Print("Perimeter of Circle : ")
	return 2 * math.Pi * c.radius
}

func (c Circle) Area() float64 {
	fmt.Print("Area of Circle : ")
	return math.Pi * c.radius * c.radius
}

func (t Triangle) Perimeter() float64 {
	fmt.Print("Perimeter of Triangle : ")
	return 3 * t.side
}

func (t Triangle) Area() float64 {
	fmt.Print("Area of Triangle : ")
	return (t.side * t.height)/2
}

func (r Rectangle) Perimeter() float64 {
	fmt.Print("Perimeter of Rectangle : ")
	return 2 * (r.height + r.width)
}

func (r Rectangle) Area() float64 {
	fmt.Print("Perimeter of Circle : ")
	return r.height * r.width;
}

func resultInfo(s Shape) {
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}

func Result() {
	c := Circle{radius: 2.5}
	t := Triangle{ height: 4.0, side: 3.0}
	r := Rectangle{width: 3.0, height: 4.0}

	resultInfo(c)
	resultInfo(t)
	resultInfo(r)
}

