package functions

import "fmt"

func ReversedNum(){
	fmt.Print("Write a number to reverse: ")
	var num int
	fmt.Scanln(&num)

	var rev_num int

	for num != 0{
		remainder := num % 10
		rev_num += remainder
		rev_num*=10
		num/=10
	}
	rev_num/=10
	fmt.Println(rev_num)

}