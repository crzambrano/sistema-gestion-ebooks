package handler

import (
    "encoding/json"
    "net/http"

    "sistema-gestion-ebooks/internal/auth"
)

type userReq struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

var userDB = map[string]string{} // username -> hash (demo)

// RegisterUser POST /users/register
func RegisterUser(w http.ResponseWriter, r *http.Request) {
    var req userReq
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if _, exists := userDB[req.Username]; exists {
        http.Error(w, "user exists", http.StatusConflict)
        return
    }
    hash, err := auth.HashPassword(req.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    userDB[req.Username] = hash
    w.WriteHeader(http.StatusCreated)
}

// LoginUser POST /users/login
func LoginUser(w http.ResponseWriter, r *http.Request) {
    var req userReq
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    hash, ok := userDB[req.Username]
    if !ok || !auth.CheckPasswordHash(req.Password, hash) {
        http.Error(w, "invalid credentials", http.StatusUnauthorized)
        return
    }
    token, _ := auth.GenerateToken(req.Username)
    _ = json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// AuthMiddleware validates JWT in Authorization header.
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenStr := r.Header.Get("Authorization")
        if len(tokenStr) < 7 || tokenStr[:7] != "Bearer " {
            http.Error(w, "missing token", http.StatusUnauthorized)
            return
        }
        tokenStr = tokenStr[7:]
        if _, err := auth.Parse(tokenStr); err != nil {
            http.Error(w, "invalid token", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    })
}
