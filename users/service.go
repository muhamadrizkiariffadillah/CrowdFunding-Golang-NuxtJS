package users

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

// Service defines the available methods for user-related operations.
type Service interface {
	// RegisterUser registers a new user based on the provided input.
	// Returns the created user or an error if the process fails.
	RegisterUser(input RegisterUserInput) (Users, error)
}

// service implements the Service interface.
type service struct {
	repo Repository // Repository to interact with the database.
}

// UserServices creates a new instance of the service.
func UserServices(repo Repository) *service {
	return &service{repo}
}

// RegisterUser handles user registration.
// It hashes the password, creates a new user record, and saves it to the repository.
func (s *service) RegisterUser(input RegisterUserInput) (Users, error) {
	// Create a new Users struct and populate its fields from the input.
	user := Users{
		FullName:   input.FullName,
		Occupation: input.Occupation,
		Email:      input.Email,
		Token:      "",     // TODO: Should be set later
		Role:       "user", // Default role
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Hash the user's password using bcrypt.
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return Users{}, err
	}
	user.HashPassword = string(passwordHash)

	// Save the new user in the repository.
	newUser, err := s.repo.Save(user)
	if err != nil {
		return Users{}, err
	}
	return newUser, nil
}
