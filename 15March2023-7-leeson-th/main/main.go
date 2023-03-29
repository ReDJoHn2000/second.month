package main

import (
	// "Amirbek/bootcamp/04152023/functions"
	"fmt"
)

type rectangle struct{
	height, width int
}

func main(){

	side := rectangle{height: 2, width: 3}

	a := 3
	b := &a

	fmt.Println(*b)
	
	fmt.Println("The Area of rectangle with sides: ", side, " is ", side.Area())
	fmt.Println("The Perimeter of rectangle with sides: ", side, " is ", side.Perimetr())
}

func (r *rectangle) Area() int {
	return r.height *r.width;
}

func (r rectangle) Perimetr() int {
	return 2*(r.height + r.width);
}