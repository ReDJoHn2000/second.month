package maps

import (
	"fmt"
)

func CreateMap(){
	mapsAuthor := make(map[string]string)

	var author string

	mapsAuthor["William"] = "Romeo"
	mapsAuthor["Konan"] = "Sherlock"
	mapsAuthor["Rowling"] = "Hurry Potter"

	fmt.Println("Available authors, below: ")

	for key := range mapsAuthor{
		fmt.Println("Authors: ", key)
	}

	fmt.Print("What Author u want to delete?: ")
	fmt.Scan(&author)

	for key := range mapsAuthor{
		if key == author {
			delete(mapsAuthor, key)
		}
	}

	fmt.Println("After deleted: ")

	for key, val := range mapsAuthor{
		fmt.Println("Authors: ", key, ", books: ", val)
	}
}