package main

import (
	"fmt"
	"go-sanbox/book"
	"go-sanbox/user"
)

func main() {
	var userD user.User
	var users = user.Users{}
	var books = book.Books{}

	fmt.Println("Хочешь зарегистрироваться? y/n")

	var goReg = "n"
	fmt.Scan(&goReg)

	if goReg == "y" {
		userD, users = user.MakeUser(users)
	} else {
		userD, users = user.MakeUser(users)
	}

	for i, userY := range users.Users {
		fmt.Printf("Пользователь %d: %s, %d лет\n", i+1, userY.Name, userY.Age)
	}

	fmt.Printf("%s, спасибо, что ты мне написал. Тебе %d лет, а твой рост: %.2f см\n", userD.Name, userD.Age, userD.Height)

	books = book.AddBook(books)

	for i, bookB := range books.Book {
		fmt.Printf("Книга %d: %s, %d стр.\n", i+1, bookB.Title, bookB.Length)
	}
}
