package main

import (
	"fmt"
	"math"
)

func main(){
	quadEqu()
}

func quadEqu(){
	var a, b, c float64

	for {
		defer PanicHandle()

		fmt.Println("Write the a, b, c of the equation: ")
		fmt.Print("a = ")
		fmt.Scan(&a)
		fmt.Print("b = ")
		fmt.Scan(&b)
		fmt.Print("c = ")
		fmt.Scan(&c)
		
		if a == 0{
			panic("a equals to 0, it's not a quadratic equation!")
		}
		
		x1 := ( -b + math.Sqrt(b*b - 4*a*c))/ (2*a)
		
		x2 := ( -b - math.Sqrt(b*b - 4*a*c))/ (2*a)
		
		fmt.Println("The roots of the equation are x1 = ", x1, " x2 = ", x2)
	}
}

func PanicHandle(){
	r := recover()

	if r != nil {
		fmt.Println("The function is recovered! ", r)
	}
}