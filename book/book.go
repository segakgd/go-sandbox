package book

import "fmt"

type Books struct {
	Book []Book
}

type Book struct {
	Title  string
	Length uint16
}

func AddBook(books Books) Books {
	fmt.Println("Хочешь добавить книгу? y/n")

	var iaAddBook = "n"
	fmt.Scan(&iaAddBook)

	if iaAddBook != "y" {
		return books
	}

	var title string
	var length uint16

	fmt.Println("Введи название книги")
	fmt.Scanln(&title)

	fmt.Println("Сколько количесво страниц")
	fmt.Scan(&length)

	var book = Book{title, length}

	books.Book = append(books.Book, book)

	fmt.Printf("Книга %s добавлена", book.Title)

	books = AddBook(books)

	return books
}

func ViewBooks(books Books) {
	for i, book := range books.Book {
		fmt.Printf("Книга %d: %s, %d стр.\n", i+1, book.Title, book.Length)
	}
}
