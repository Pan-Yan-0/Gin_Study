// app/service/LoginService.go
package service

import (
	"LoginStudy/app/database"
	"database/sql"
	"fmt"
)

// Authenticate checks the user's credentials
func Authenticate(username, password string) error {
	var dbPassword string

	// Query the database to get the stored password for the given username
	err := database.DB.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&dbPassword)
	if err == sql.ErrNoRows {
		// User not found
		return fmt.Errorf("username not found")
	} else if err != nil {
		// Other database errors
		return fmt.Errorf("internal server error")
	}

	// In a real-world app, you should hash the password and compare it (bcrypt)
	if dbPassword != password {
		return fmt.Errorf("incorrect password")
	}

	// Successful authentication
	return nil
}
