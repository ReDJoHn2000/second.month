package main

import (
	"fmt"
)

type BookData interface {
	GetBooks() []Book
	AddBook(Book)
	RemoveBook(Book)
	UpdateBook(Book)
	GetBookById(Book)
	GetBookByCategory(Book)
}

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

type BookLibrary struct {
	Books []Book
}

// Geting the books
func (b *BookLibrary) GetBooks() []Book {
	return b.Books
}

// Inserting the Book to the list
func (b *BookLibrary) AddBook(book Book) {
	b.Books = append(b.Books, book)
}

func main() {

	library := &BookLibrary{Books: data.Books} // one of the main part

	// Example usage:
	fmt.Println(library.GetBooks())

	library.AddBook(Book{
		Title:  "1984",
		Author: "George Orwell",
		Year:   1949,
	})

	fmt.Println(library.GetBooks())

	library.RemoveBook(Book{
		Title:  "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Year:   1925,
	})

	library.UpdateBook(Book{
		Title:  "The Great Gatsby",
		Author: "Tohir Abdullaev",
		Year:   2020,
	})

	fmt.Println(library.GetBooks())
}
