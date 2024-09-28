package users

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Service defines the available methods for user-related operations.
type Service interface {
	// RegisterUser registers a new user based on the provided input.
	// Returns the created user or an error if the process fails.
	RegisterUser(input RegisterUserInput) (Users, error)
	// LoginUser is a login endpoint based on the email users
	// Returns the user or an error email or password do not match
	LoginUser(input LoginUserInput) (Users, error)

	IsEmailAvailable(input CheckEmailInput) (bool, error)

	UploadAvatar(Id int, fileLocation string) (Users, error)

	GetUserByID(userId int) (Users, error)

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

// LoginUser handles user login
// it finds a user by email and returns user or error
func (s *service) LoginUser(input LoginUserInput) (Users, error) {
	userEmail := input.Email
	userPassword := input.Password

	user, err := s.repo.FindByEmail(userEmail)
	if err != nil {
		return Users{}, err
	}

	if user.Id == 0 {
		return Users{}, errors.New("the user is not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(userPassword))
	if err != nil {
		return Users{}, errors.New("wrong password")
	}

	return user, nil
}

// IsEmailAvailable handles check user email
// return true or false
func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repo.FindByEmail(email)

	if err != nil || user.Id != 0 {
		return false, nil
	}

	return true, nil
}

func (s *service) UploadAvatar(Id int, fileLocation string) (Users, error) {
	user, err := s.repo.FindById(Id)

	if err != nil {
		return Users{}, err
	}

	user.AvatarFileName = fileLocation

	updatedUser, err := s.repo.Save(user)

	if err != nil {
		return Users{}, err
	}

	return updatedUser, nil
}

func (s *service) GetUserByID(userId int)(Users,error)  {

	user, err := s.repo.FindById(userId)

	if err != nil {

		return Users{}, err

	}

	if user.Id == 0{
		return Users{},errors.New("user not found")
	}

	return user,nil
}