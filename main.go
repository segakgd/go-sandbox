package main

import "fmt"

type Person struct {
	Name   string
	Age    uint8
	Height float32
}

func main() {
	var name string
	var age uint8
	var height float32

	fmt.Println("Привет, как тебя зовут?")
	fmt.Scan(&name)

	fmt.Println("Отлично ", name, ", а сколько тебе лет?")
	fmt.Scan(&age)

	fmt.Println("А рост?")
	fmt.Scan(&height)

	var person = Person{name, age, height}

	fmt.Printf("%s, спасибо, что ты мне написал. Тебе %d лет, а твой рост: %.2f см\n", person.Name, person.Age, person.Height)
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
