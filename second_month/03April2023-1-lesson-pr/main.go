package main

import (
	"fmt"
	"strconv"
)

func main(){
	// Task1()
	// fmt.Println(Task2Leetcode1832())
	// fmt.Println(Task3Leetcode2520())
	// fmt.Println(Task4Leetcode1773())
	// fmt.Println(Task5Leetcode2000())
	fmt.Println(Task5Leetcode1995())
}

func Task1(){
	arr := make([]int, 4)

	arr[0] = 4
	arr[1] = 6
	arr[2] = 3
	arr[3] = 8

	var result []int

	for i := 0; i < len(arr); i++ {
	
		arr2 := arr[i] + i + 1

		// result = append(result, arr2)

		// fmt.Println(arr2)

		if arr2 > 9{
			result = append(result, arr2 % 10)
		} else {
			result = append(result, arr2)
		}
	}
	fmt.Print(result)
}

func Task2Leetcode1832() bool {
	sentence := "thequickbrownfoxjumpsoverthelazydog"
	
	alph := make(map[rune]bool)

	for _, l := range sentence{
		alph[l] = true
	}
	return len(alph) == 26
}

func Task3Leetcode2520() int {
	num := 8
	digits := strconv.Itoa(num)
	var count int

	for i := 0; i < len(digits); i++ {

		val, _ := strconv.Atoi(string(digits[i]))

		// fmt.Println(val)

		if num % val == 0{
			count++
		}
	}
	return count
}

func Task4Leetcode1773() int {

	// items [][]string, ruleKey string, ruleValue string

	items := [][]string {{"phone","blue","pixel"},{"computer","silver","phone"},{"phone","gold","iphone"}}
    var ruleKey string = "type"
	var ruleValue string = "phone"
    var count int
    
    for _, i := range items {
        if ruleKey == "type" {
            if ruleValue == i[0]{
                count++
                }
        } else if ruleKey == "color"{
            if ruleValue == i[1] {
                count++
                }
        } else if ruleKey == "name"{
            if ruleValue == i[2]{
                count++
                }
        }
    }
return count
}

func Task5Leetcode2000() string {
	
	word := "abczde"
	ch := "j"
	var First string
	var Second string
	var new string = ""

	for i, v := range word {
		if string(v) == string(ch) {
			First = word[:i+1]
			Second = word[i+1:]
			// fmt.Println(new, "reverse")

			for k := len(First) - 1; k >= 0; k--{
				// fmt.Println(string(First[k]), "first")
				new += string(First[k])
				// fmt.Println(new, "reverse")
			}
			
			for j := 0; j < len(Second); j++ {
				// fmt.Println(string(Second[j]), "second")
				new += string(Second[j])
				// fmt.Println(new, "reverse word")
			}
			return new
		} else {
			new = ""
			// fmt.Println(new)
		}
	}
	// fmt.Println(new, "reverse")
	return word
}

func Task5Leetcode1995() int {
	var nums = make([]int, 4)

	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums[3] = 6

	var count, length = 0, len(nums)

	for i := 0; i < length - 3; i++ {
		for k := i + 1; k < length - 2; k++ {
			for j := k + 1; j < length - 1; j++ {
				for d := j + 1; d < length; d++ {
					if nums[i] + nums[k] + nums[j] == nums[d] {
						count++
					}
				}				
			}
		}
	}
	return count
}