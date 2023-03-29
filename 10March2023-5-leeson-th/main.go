package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type City struct{
	FullName string
	FullAddress struct{
		Street string `json:"street"`
		City string `json:"city"`
	}
}

func main(){
NameCity()
}

func NameCity(){
	data, err := ioutil.ReadFile("city.json")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	// fmt.Println(data)

	var cityInfo []City
	err = json.Unmarshal(data, &cityInfo)
	if err != nil {
		fmt.Println("Error unmarshling JSON: ", err)
		return
	}

	fmt.Println(cityInfo)

	fmt.Println("Here are the city names: ")
	for _, item := range cityInfo{
		fmt.Println("Name: ", item.FullName)
		fmt.Println("Street: ", item.FullAddress)
		fmt.Println("City: ", item.FullAddress)
	}
}