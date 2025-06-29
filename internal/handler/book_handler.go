package handler

import (
    "encoding/json"
    "net/http"

    "github.com/go-chi/chi/v5"
    "sistema-gestion-ebooks/internal/book"
    "sistema-gestion-ebooks/pkg/storage"
)

var bookStore = storage.NewJSONStore[book.Book]("data/libros.json")

// CreateBook POST /books
func CreateBook(w http.ResponseWriter, r *http.Request) {
    var req book.Book
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    // force new ID
    b := book.NewBook(req.Title, req.Author, req.Year, req.Category)

    var books []book.Book
    _ = bookStore.Load(&books)
    books = append(books, b)
    if err := bookStore.Save(books); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    _ = json.NewEncoder(w).Encode(b)
}

// ListBooks GET /books
func ListBooks(w http.ResponseWriter, r *http.Request) {
    var books []book.Book
    if err := bookStore.Load(&books); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    _ = json.NewEncoder(w).Encode(books)
}

// GetBook GET /books/{id}
func GetBook(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    var books []book.Book
    _ = bookStore.Load(&books)
    for _, b := range books {
        if b.ID == id {
            _ = json.NewEncoder(w).Encode(b)
            return
        }
    }
    http.NotFound(w, r)
}

// DeleteBook DELETE /books/{id}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    var books []book.Book
    _ = bookStore.Load(&books)

    filtered := make([]book.Book, 0, len(books))
    for _, b := range books {
        if b.ID != id {
            filtered = append(filtered, b)
        }
    }
    if len(filtered) == len(books) {
        http.NotFound(w, r)
        return
    }
    if err := bookStore.Save(filtered); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

// SearchBooks GET /books/search?title=foo
func SearchBooks(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("title")
    var books []book.Book
    _ = bookStore.Load(&books)

    var res []book.Book
    for _, b := range books {
        if q == "" || containsIgnoreCase(b.Title, q) {
            res = append(res, b)
        }
    }
    _ = json.NewEncoder(w).Encode(res)
}

// helper
func containsIgnoreCase(s, substr string) bool {
    return len(substr) == 0 || (len(s) >= len(substr) &&
        (s == substr || containsIgnoreCase(s[1:], substr)))
}
