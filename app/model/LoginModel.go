// app/model/LoginModel.go
package model

// User represents a user model in the database
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
