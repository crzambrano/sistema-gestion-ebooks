package book

import "github.com/google/uuid"

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
	Genre  string `json:"genre"`
}

func NewBook(title, author string, year int, genre string) Book {
	return Book{
		ID:     uuid.New().String(),
		Title:  title,
		Author: author,
		Year:   year,
		Genre:  genre,
	}
}
