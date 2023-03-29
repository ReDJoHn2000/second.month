package main

import "fmt"

func main(){

	// a := "(){}[]"

	// b := "))()}{"

	isValid()

	// fmt.Println(isValid("))()}{"))

	// a := "YesYesYessdfkurn"

	// b := "Yes"

	// var count int

	// for i := 0; i < len(a); i++ {
	// 	for k := 0; k < len(b); k++ {
	// 		if a[i] == b[k]{
	// 			fmt.Println("--", a[i])
	// 			fmt.Println(b[i])
	// 			count++
	// 		}
	// 	}
	// }

	// fmt.Println(count)

	// word := "01011010110"

	// var count int
	// target := '1'

	// for i := 0; i < len(word); i++ {

	// 	if word[i] == target{

	// 	}

		// fmt.Printf("%c\n", word[i])
	// }
	// fmt.Println(count)
// 	if len(s) > 1{
// 		if s[len(s)-2]=='[' && s[len(s)-1] == ']'{
// 			s = s[:len(s)-2]
// 		} else if s[len(s)-2]=='{' && s[len(s)-1] == '}'{
// 			s = s[:len(s)-2]
// 		}else if s[len(s)-2]=='(' && s[len(s)-1] == ')'{
// 			s = s[:len(s)-2]
// 		}        
// }
// if len(s) == 0{
// 	return true
// }
// return false
}

func isValid() {

	str := "(){}[]"

    var stack []byte

    for _, char := range str{

		if char == '(' ||char == '{' || char =='['{
			
			stack = append(stack, byte(char))

			
		}
	}
	fmt.Println(stack)
}