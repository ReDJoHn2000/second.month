package functions

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var Bookings = make([]Buyer, 0)

type Buyer struct{
	firstName string
	lastName string
	email     string
	ticketAmmount int
	phone int
}

func UserInfo(){

	var firstName, lastName, email string
	var ticketAmmount, phone int

	var ticketNumber int = 50

	var ticketAvailable int = 50

	fmt.Println("Welcome to a ComicCon!")
	fmt.Print("\n")

	// userInfo := make([]string, 0)
	
	for ticketAvailable >= 0 {

	fmt.Print("Your Name: ")
	fmt.Scanln(&firstName)
	fmt.Print("Your Surname: ")
	fmt.Scanln(&lastName)
	fmt.Print("Your email: ")
	fmt.Scanln(&email)
	fmt.Print("Ticket Number: ")
	fmt.Scanln(&ticketAmmount)
	fmt.Print("Phone Number: ")
	fmt.Scanln(&phone)
	fmt.Print("\n")
	
	isValidFullName, isValidEmail, isValidTicketNumber, isValidPhone := ValidateUserInput(firstName, lastName, email, phone, ticketAmmount)
	
	// fmt.Println(isValidFullName)
	// fmt.Println(isValidEmail)
	// fmt.Println(isValidTicketNumber)
	// fmt.Println(isValidPhone)

	if isValidFullName && isValidEmail && isValidTicketNumber && isValidPhone{

	if ticketAmmount > ticketAvailable{
		fmt.Println("Available ticket number is: ", ticketAvailable)
		fmt.Print("Input another Ticket Number: ")
		fmt.Scanln(&ticketAmmount)
	}
	
	ticketNumber = ticketNumber - ticketAmmount
	
	// userInfo = append(userInfo, firstName)
	
	Bookings = append(Bookings, Buyer{})
	
	user := Buyer{
		firstName: firstName,
		lastName: lastName,
		email: email, 
		ticketAmmount: ticketAmmount,
		phone: phone,
	}
	
	vals := reflect.ValueOf(user)
	keys := []string{"Firstname: ", "Surname: ", "Email: ", "Ticket Number: ", "Phone Number: "}
	
	for i := 0; i < vals.NumField(); i++ {
		fmt.Print(keys[i])
		fmt.Println(vals.Field(i))
	}
	fmt.Print("\n")
	
	// userInfoM := make(map[int]string)
	
	// for i := 0; i < len(userInfo); i++{
		// 	   Bookings[i + 1] = userInfo[i]
		// }
			
			Bookings := GetFirstNames()
			
			fmt.Println(append(Bookings, firstName))
			
			if ticketNumber == 0{
				fmt.Println("Tickets are run out.")
				break
			}
			
			ticketAvailable = ticketNumber
			fmt.Print("\n")
		} else {
			if !isValidFullName{
				fmt.Println("Name should be more than 3 letters!")
			}
			if !isValidEmail{
				fmt.Println("Email should contain the sign '@'.")
			}
			if !isValidPhone{
				fmt.Println("The number should be from Uzbekistan Companies!")
			}
			if !isValidTicketNumber{
				fmt.Println("Input positive number!")
			}
		}
}

}

func GetFirstNames()[]string{
	firstName := []string{}
	for _, booking := range Bookings{
		firstName = append(firstName, booking.firstName)
	}
	return firstName
}

func ValidateUserInput(firstName string, lastName string, email string, phone int, ticketAmmount int) (bool, bool, bool, bool){
	isValidFullName := len(firstName) >= 3 && len(lastName) >= 3;
	isValidEmail := len(email) >= 5 && strings.Contains(email, "@")
	isValidTicketNumber := ticketAmmount > 0
	isValidPhone := len(strconv.Itoa(phone)) > 10

	return isValidFullName, isValidEmail, isValidTicketNumber, isValidPhone
}