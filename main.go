package main

import "fmt"

type Users struct {
	Users []User
}

type Books struct {
	Book []Book
}

type User struct {
	Name   string
	Age    uint8
	Height float32
}

type Book struct {
	Title  string
	Length uint16
}

func main() {
	var user User
	var users = Users{}
	var books = Books{}

	fmt.Println("Хочешь зарегистрироваться? y/n")

	var goReg = "n"
	fmt.Scan(&goReg)

	if goReg == "y" {
		user, users = makeUser(users)
	} else {
		user, users = makeUser(users)
	}

	for i, user := range users.Users {
		fmt.Printf("Пользователь %d: %s, %d лет\n", i+1, user.Name, user.Age)
	}

	fmt.Printf("%s, спасибо, что ты мне написал. Тебе %d лет, а твой рост: %.2f см\n", user.Name, user.Age, user.Height)

	books = addBook(books)

	for i, book := range books.Book {
		fmt.Printf("Книга %d: %s, %d стр.\n", i+1, book.Title, book.Length)
	}
}

func makeUser(users Users) (User, Users) {
	var name string
	var age uint8
	var height float32

	fmt.Println("Введи имя")
	fmt.Scanln(&name)

	fmt.Println("Введи возраст")
	fmt.Scan(&age)

	fmt.Println("Введи рост")
	fmt.Scan(&height)

	var user = User{name, age, height}

	users.Users = append(users.Users, user)

	return user, users
}

func addBook(books Books) Books {
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

	addBook(books)

	return books
}

//import (
//	"log"
//	"net/http"
//)
//
//func main() {
//	log.Print("Запуск сервера...")
//
//	http.HandleFunc("/", routeHandler)
//	http.HandleFunc("/test", routeHandlerTest)
//
//	addr := ":8080"
//	log.Printf("Сервер запущен на %s", addr)
//
//	if err := http.ListenAndServe(addr, nil); err != nil {
//		log.Fatalf("Ошибка запуска сервера: %v", err)
//	}
//}
