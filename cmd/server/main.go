package main

import (
    "log"
    "net/http"

    "sistema-gestion-ebooks/internal/handler"
)

func main() {
    log.Println("Servidor escuchando en :8080")
    if err := http.ListenAndServe(":8080", handler.Routes()); err != nil {
        log.Fatal(err)
    }
}
