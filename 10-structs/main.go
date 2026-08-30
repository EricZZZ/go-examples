package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Titile string
	Author string
	Pages  int
}

type Library struct {
	Name  string
	Books []Book
}

func (b Book) summary() string {
	return fmt.Sprintf("%s by %s, %d pages", b.Titile, b.Author, b.Pages)
}

func (b *Book) UpdatePages(newPages int) {
	if newPages > 0 {
		b.Pages = newPages
	} else {
		fmt.Println("Invalid page count. Pages must be greater than zero.")
	}
}

func main() {
	// myBook := Book{
	// 	Titile: "The Go Programming Language",
	// 	Author: "Alan Donovan",
	// 	Pages:  0,
	// }

	var myBook Book
	myBook.Titile = "The Go Programming Language"
	myBook.Author = "Alan Donovan"
	myBook.Pages = 0

	fmt.Println(myBook)

	myLibrary := Library{
		Name:  "City Library",
		Books: []Book{myBook, {Titile: "Go in Action", Author: "Chris Williams", Pages: 599}},
	}

	fmt.Println(myLibrary)

	fmt.Println(myBook.summary())

	myBook.UpdatePages(2222)
	fmt.Println(myBook)

	book1 := Book{
		Titile: "The Go Programming Language",
		Author: "Alan Donovan",
		Pages:  0,
	}

	book2 := Book{
		Titile: "The Go Programming Language",
		Author: "Alan Donovan",
		Pages:  0,
	}

	if book1 == book2 {
		fmt.Println("book1 and book2 are equal")
	} else {
		fmt.Println("book1 and book2 are not equal")
	}

	jsonString := `{"Name":"Alice","Age":12}`

	person := struct {
		Name string
		Age  int
	}{}

	err := json.Unmarshal([]byte(jsonString), &person)

	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
	} else {
		fmt.Printf("Name:%s, Age: %d", person.Name, person.Age)
	}
}
