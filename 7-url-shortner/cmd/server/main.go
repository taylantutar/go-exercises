package main

import (
    "log"
    "net/http"
    "7-url-shortener/internal/handler"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/shorten", handler.ShortenHandler)
    mux.HandleFunc("/s/", handler.RedirectHandler)

    log.Println("Server started at :8080")
    http.ListenAndServe(":8080", mux)
}
