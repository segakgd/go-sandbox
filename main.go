package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", routeHandler)

	addr := ":8080"
	log.Printf("Сервер запущен на %s", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
