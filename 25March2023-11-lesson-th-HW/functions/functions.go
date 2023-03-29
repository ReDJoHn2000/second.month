package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type BookData interface{
	GetBooks() []BookInfo
	AddBook(BookInfo)
	RemoveBook(BookInfo)
	UpdateBook(BookInfo)
	GetBookById(BookInfo)
	GetBookByCategory(BookInfo)
}

type BookInfo struct{
	Id 			int 	`json:"id"`
	Title 		string 	`json:"title"`
	Author 		string 	`json:"author"`
	Year 		int 	`json:"year"`
	Status 		string 	`json:"status"`
	Price 		int 	`json:"price"`
	Period 		int 	`json:"period"`
	Category 	string 	`json:"category"`
	Page 		int 	`json:"page"`
}

type BookLibrary struct{
	Books []BookInfo
}

func (b *BookLibrary) GetBooks() []BookInfo {
	return b.Books
}

func (b *BookLibrary) AddBook(book BookInfo) {
	b.Books = append(b.Books, book)
}

func (b *BookLibrary) RemoveBook(id BookInfo){
	for _, v := range b.Books {
		if v.Id == id.Id {
			elem := b.Books[id.Id - 1]

			fmt.Println("The removed book is : ", elem)

			b.Books = append(b.Books[:v.Id - 1], b.Books[v.Id:]...)

			fmt.Println("The updated book-list is : ", b.Books)
		}
	}
}

func (b *BookLibrary) UpdateBook(BookInfo BookInfo) {
	for _, book := range b.Books {
		if book.Id == BookInfo.Id{
			book.Title = BookInfo.Title
			book.Author = BookInfo.Author
			book.Year = BookInfo.Year
			book.Status = BookInfo.Status
			book.Price = BookInfo.Price
			book.Period = BookInfo.Period
			book.Category = BookInfo.Category
			book.Page = BookInfo.Page
		}
	}
	fmt.Println("The list after update ", BookInfo)
}

func (b *BookLibrary) GetBookById(id BookInfo) {
	for _, book := range b.Books {
		if book.Id == id.Id{
			fmt.Println(book)
		}
	}
}

func (b *BookLibrary) GetBookByCategory(category BookInfo) {
	for _, book := range b.Books {
		if book.Category == category.Category{
			fmt.Println(book)
		}
	}
}

func Functions(){

		fmt.Println("Welcome to our Library!")
		fmt.Println("What would you like to do!")

		fmt.Println("These are our books")
		data, err := ioutil.ReadFile("books.json")
		if err != nil {
			fmt.Println("Error reading File: ", err)
			return
		}

		var Books []BookInfo
		err = json.Unmarshal(data, &Books)
		if err != nil {
			fmt.Println("Error Unmarshling JSON: ", err)
			return
		}

		data, err = json.Marshal(Books)
		if err != nil {
			fmt.Println("Error marshling JSON: ", err)
			return
		}

		library := &BookLibrary{Books: Books}

		fmt.Println("- 1 -")
		fmt.Println(library.GetBooks())

		
		fmt.Println("- 2 -")
		library.AddBook(BookInfo{
		Id: 4,
		Title: "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Year: 1925,	
		Status: "given",		
		Price: 120,	
		Period: 12,	
		Category: "Sci-fic",
		Page: 20,
	})
	fmt.Println(library.GetBooks())

	err = ioutil.WriteFile("books.json", data, 0644)
		if err != nil {
			fmt.Println("Error writing File: ", err)
			return
		}
	
	fmt.Println("- 3 -")
	library.RemoveBook(BookInfo{Id: 2})

	err = ioutil.WriteFile("books.json", data, 0644)
		if err != nil {
			fmt.Println("Error writing File: ", err)
			return
		}
	
	
	fmt.Println("- 4 -")
		library.UpdateBook(BookInfo{
		Id: 4,
		Title: "The Fallen",
		Author: "F. Scott Fitzgerald",
		Year: 1825,	
		Status: "given",		
		Price: 220,	
		Period: 22,	
		Category: "Detective",
		Page: 42,
	})

	data, err = json.Marshal(Books)
		if err != nil {
			fmt.Println("Error marshling JSON: ", err)
			return
		}

	err = ioutil.WriteFile("books.json", data, 0644)
		if err != nil {
			fmt.Println("Error writing File: ", err)
			return
		}


	fmt.Println("- 5 -")
	library.GetBookById(BookInfo{Id: 1})

	fmt.Println("- 6 -")
	library.GetBookByCategory(BookInfo{Category: "Sci-fic"})
}

func ReadAndWrite(){
	data, err := ioutil.ReadFile("books.json")
		if err != nil {
			fmt.Println("Error reading File: ", err)
			return
		}

		var Books []BookInfo
		err = json.Unmarshal(data, &Books)
		if err != nil {
			fmt.Println("Error Unmarshling JSON: ", err)
			return
		}

		data, err = json.Marshal(Books)
		if err != nil {
			fmt.Println("Error marshling JSON: ", err)
			return
		}

		err = ioutil.WriteFile("books.json", data, 0644)
		if err != nil {
			fmt.Println("Error writing File: ", err)
			return
		}
}