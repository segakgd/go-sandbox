package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type Response struct {
    Message string `json:"message"`
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    response := Response{
        Message: "Какой-то ответ",
    }

    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func main() {
    http.HandleFunc("/", routeHandler)

    addr := ":8080"
    log.Printf("Сервер запущен на %s", addr)

    if err := http.ListenAndServe(addr, nil); err != nil {
        log.Fatalf("Ошибка запуска сервера: %v", err)
    }
}
