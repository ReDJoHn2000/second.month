package functions

import (
	"fmt"
	"strings"
)

func SecMinHour(){

	var seconds int

	fmt.Println("Enter time in sencods: ")
	fmt.Scan(&seconds)

	if seconds > 0 {
		days := seconds/86400
		seconds -= days*86400
		hours := seconds/3600
		seconds -= hours*3600
		minutes := seconds/60
		seconds -= minutes*60
		fmt.Println("The time in order is: ", days, " days ", hours, " hours ", minutes, " minutes ", seconds, " seconds")
	}
}

func Morse(){
	var count, word string
	
	fmt.Print("Write word to make it morse: ")
	fmt.Scan(&word)

	input := strings.ToUpper(word)

	var morseAlfa = make(map[string]string)

  morseAlfa["A"]=".-"
  morseAlfa["B"]="-..."
  morseAlfa["C"]="-.-."
  morseAlfa["D"]="-.."
  morseAlfa["E"]="."
  morseAlfa["F"]="..-."
  morseAlfa["G"]="--."
  morseAlfa["H"]="...."
  morseAlfa["I"]=".."
  morseAlfa["J"]=".---"
  morseAlfa["K"]="-.-"
  morseAlfa["L"]=".-.."
  morseAlfa["M"]="--"
  morseAlfa["N"]="-."
  morseAlfa["O"]="---"
  morseAlfa["P"]=".--."
  morseAlfa["Q"]="--.-"
  morseAlfa["R"]=".-."
  morseAlfa["S"]="..."
  morseAlfa["T"]="-"
  morseAlfa["U"]="..-"
  morseAlfa["V"]="...-"
  morseAlfa["W"]=".--"
  morseAlfa["X"]="-..-"
  morseAlfa["Y"]="-.--"
  morseAlfa["Z"]="--.."
  morseAlfa["1"]=".----"
  morseAlfa["2"]="..---"
  morseAlfa["3"]="...--"
  morseAlfa["4"]="....-"
  morseAlfa["5"]="....."
  morseAlfa["6"]="-...."
  morseAlfa["7"]="--..."
  morseAlfa["8"]="---.."
  morseAlfa["9"]="---."
  morseAlfa["0"]="-----"

//   for _, slice := range strings.Split(input, "") {
// 		fmt.Println(morseAlfa["A"])
// 	}
	// fmt.Printf("%s\n", slice)
// }

	// for i := 0; i < len(input); i++ {
	// 	fmt.Printf("%c\n", input[i])
	// }

	for i := 0; i < len(input); i++ {
		// fmt.Println(string(input[i]))
		// fmt.Println(morseAlfa[input])
		if morseAlfa[string(input[i])] != "" {
			count += morseAlfa[string(input[i])]
		}
	}
	fmt.Println(count)
}

func RomeNum(){
	var word string

	for {

		
		fmt.Print("Write Rome Number: ")
		fmt.Scan(&word)
		
		input := strings.ToUpper(word)
		
		var romanNum = make(map[string]int)
		
		romanNum["I"] = 1
		romanNum["V"] = 5
		romanNum["X"] = 10
		romanNum["L"] = 50
		romanNum["C"] = 100
		romanNum["D"] = 500
		romanNum["M"] = 1000

		max := 0
		sum := 0
		
		for i := len(input) - 1; i >= 0; i-- {
			
			if max > romanNum[string(input[i])]{
				max = romanNum[string(input[i])]
				sum -= max
			} else {
				max = romanNum[string(input[i])]
				sum += romanNum[string(input[i])]
			}
			// fmt.Println(romanNum[string(input[i])])
		}
		fmt.Println(sum)
	}
}