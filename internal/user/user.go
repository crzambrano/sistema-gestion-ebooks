package user

// User represents a system user.
type User struct {
    Username string `json:"username"`
    Password string `json:"password"` // stored as bcrypt hash
}
