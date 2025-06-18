package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sistema-gestion-ebooks/internal/book"
)

type Storage struct {
	FilePath string
}

func NewStorage(path string) *Storage {
	return &Storage{FilePath: path}
}

func (s *Storage) Save(libros []book.Book) error {
	data, err := json.MarshalIndent(libros, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(s.FilePath, data, 0644)
}

func (s *Storage) Load() ([]book.Book, error) {
	file, err := os.ReadFile(s.FilePath)
	if err != nil {
		return nil, err
	}

	var libros []book.Book
	err = json.Unmarshal(file, &libros)
	return libros, err
}
