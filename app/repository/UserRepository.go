// app/repository/UserRepository.go
package repository

import (
	"LoginStudy/app/database"
	"database/sql"
	"fmt"
)

// UserRepository handles data access for users
type UserRepository struct{}

// GetPasswordByUsername fetches the password for a given username
func (r *UserRepository) GetPasswordByUsername(username string) (string, error) {
	var password string

	// Query the database for the password
	err := database.DB.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&password)
	if err == sql.ErrNoRows {
		return "", fmt.Errorf("username not found")
	} else if err != nil {
		return "", fmt.Errorf("error querying database: %v", err)
	}

	return password, nil
}
