package storage

import (
    "encoding/json"
    "os"
)

// JSONStore provides generic file persistence.
type JSONStore[T any] struct {
    Path string
}

// NewJSONStore instantiates a JSONStore.
func NewJSONStore[T any](path string) *JSONStore[T] { return &JSONStore[T]{Path: path} }

// Load reads JSON file into slice.
func (s *JSONStore[T]) Load(out *[]T) error {
    file, err := os.ReadFile(s.Path)
    if err != nil {
        if os.IsNotExist(err) {
            *out = []T{}
            return nil
        }
        return err
    }
    return json.Unmarshal(file, out)
}

// Save writes slice to JSON file.
func (s *JSONStore[T]) Save(in []T) error {
    data, err := json.MarshalIndent(in, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(s.Path, data, 0644)
}
