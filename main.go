package main

import "fmt"

func main() {
	fmt.Println("Привет, Go!")
	fmt.Println("w1w31231")
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
