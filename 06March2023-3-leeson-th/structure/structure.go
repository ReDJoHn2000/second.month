package structure

import (
	"fmt"
	"reflect"
)

type car struct{
	name string
	model string
	color string
	price int
}

func NewCar(){

	var name, model, color string
	var price int

	fmt.Print("Write name: ")
	fmt.Scan(&name)
	fmt.Print("Write model: ")
	fmt.Scan(&model)
	fmt.Print("Write color: ")
	fmt.Scan(&color)
	fmt.Print("Write price: ")
	fmt.Scan(&price)
	fmt.Print("\n")

	c := car{
		name: name,
		model: model, 
		color: color, 
		price: price,
	}
	
	a := reflect.ValueOf(c)
	b := []string{"Name: ", "Model: ", "Color: ", "Price: "}
	
	for i := 0; i < a.NumField(); i++{
		fmt.Print(b[i])
		fmt.Println(a.Field(i))
	}
	fmt.Print("\n")
}
