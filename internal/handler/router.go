package handler

import (
    "encoding/json"
    "net/http"

    "github.com/go-chi/chi/v5"
)

// Routes returns configured router with all endpoints.
func Routes() http.Handler {
    r := chi.NewRouter()

    r.Post("/users/register", RegisterUser)
    r.Post("/users/login", LoginUser)

    r.Route("/books", func(br chi.Router) {
        br.Use(AuthMiddleware)
        br.Get("/", ListBooks)
        br.Post("/", CreateBook)
        br.Get("/{id}", GetBook)
        br.Delete("/{id}", DeleteBook)
        br.Get("/search", SearchBooks)
    })

    r.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
        _ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
    })

    return r
}
