package main

import "fmt"

type Person struct {
	Name   string
	Age    uint8
	Height float32
}

type Book struct {
	Title  string
	Author Person
	Length uint16
}

func main() {
	var name string
	var age uint8
	var height float32

	fmt.Println("Привет, как тебя зовут?")
	fmt.Scanln(&name)

	fmt.Println("Отлично ", name, ", а сколько тебе лет?")
	fmt.Scan(&age)

	fmt.Println("А рост?")
	fmt.Scan(&height)

	var person = Person{name, age, height}

	fmt.Printf("%s, спасибо, что ты мне написал. Тебе %d лет, а твой рост: %.2f см\n", person.Name, person.Age, person.Height)

	fmt.Println("Хочешь добавить книгу? y/n")

	var iaAddBook = "n"

	fmt.Scan(&iaAddBook)

	if iaAddBook == "y" {
		var title string
		var length uint16

		fmt.Println("Как называется книга?")
		fmt.Scanln(&title)

		fmt.Println("Сколько в ней страниц")
		fmt.Scan(&length)

		var book = Book{title, person, length}

		fmt.Printf("%s, спасибо что ты добавил книгу %s, количество страниц: %d \n", book.Title, book.Author.Name, book.Length)

	} else {
		fmt.Println("Как хочешь")
	}
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
