package main

// import (
// "fmt"
// "os"
// )

// func main(){

// 	panic("a problem")

// 	_, err := os.Create("/tmp/file")
// 	if err != nil{
// 		panic(err)
// 	}
// }

// func f() {

// 	defer func() {
// 		if r := recover();
// 		r != nil {
// 			fmt.Println("Recovered in f", r)
// 		}
// 	}()

// 	fmt.Println("Calling g.")

// 	g(0)

// 	fmt.Println("Returned normally from g.")
// }

// func g(i int) {

// 	if i > 3 {
// 		fmt.Println("Panicking!")
// 		panic(fmt.Sprintf("%v", i))
// 	}

// 	defer fmt.Println("Defer in g", i)

// 	fmt.Println("Printing in g", i)

// 	g(i + 1)
// }
