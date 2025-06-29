package auth

import (
    "time"

    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("change‑me‑in‑prod")

// HashPassword hashes a plain‑text password.
func HashPassword(pw string) (string, error) {
    b, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
    return string(b), err
}

// CheckPasswordHash compares a plain‑text password with its hash.
func CheckPasswordHash(pw, hash string) bool {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw)) == nil
}

// GenerateToken returns a signed JWT valid for 24h.
func GenerateToken(username string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sub": username,
        "exp": time.Now().Add(24 * time.Hour).Unix(),
    })
    return token.SignedString(jwtKey)
}
