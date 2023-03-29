package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// This is the Struct
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	// Read the JSON file
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal the JSON data into a slice of Person structs
	var people []Person
	err = json.Unmarshal(data, &people)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Print the data
	for _, person := range people {
		fmt.Println(person.Name, person.Age)
	}

	// add a new person
	people = append(people, Person{Name: "Amirbek", Age: 23}, Person{Name: "John", Age: 30})

	// update a person
	for i, person := range people {
		if person.Name == "Diyorbek" {
			people[i].Age = 31
		}
	}

	data, err = json.Marshal(people)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Write the JSON data to a file
	err = ioutil.WriteFile("data.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

}
