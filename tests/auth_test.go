package auth_test

import (
    "testing"

    "sistema-gestion-ebooks/internal/auth"
)

func TestHashAndCheck(t *testing.T) {
    hash, err := auth.HashPassword("secret")
    if err != nil {
        t.Fatalf("err %v", err)
    }
    if !auth.CheckPasswordHash("secret", hash) {
        t.Fatal("password should match")
    }
}
