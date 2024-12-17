// app/service/LoginService.go
package service

import (
	"LoginStudy/app/repository"
	"fmt"
)

// LoginService handles the business logic for user authentication
type LoginService struct {
	UserRepo *repository.UserRepository
}

// Authenticate checks the user's credentials
func (s *LoginService) Authenticate(username, password string) error {
	// Get the stored password from the repository
	dbPassword, err := s.UserRepo.GetPasswordByUsername(username)
	if err != nil {
		return err
	}

	// In a real-world app, you should hash the password and compare it (bcrypt)
	if dbPassword != password {
		return fmt.Errorf("incorrect password")
	}

	// Successful authentication
	return nil
}
