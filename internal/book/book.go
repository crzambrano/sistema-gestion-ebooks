package book

import "github.com/google/uuid"

// Book represents an e‑book in the system.
type Book struct {
    ID       string `json:"id"`
    Title    string `json:"title"`
    Author   string `json:"author"`
    Year     int    `json:"year"`
    Category string `json:"category"`
}

// NewBook creates a new Book with a unique ID.
func NewBook(title, author string, year int, category string) Book {
    return Book{
        ID:       uuid.NewString(),
        Title:    title,
        Author:   author,
        Year:     year,
        Category: category,
    }
}
