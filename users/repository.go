package users

import "gorm.io/gorm"

// Repository defines method for users
type Repository interface {
	Save(user Users) (Users, error)
}

// repository implements the Repository interface
type repository struct {
	db *gorm.DB // Gorm database connection
}

// UserRepository creates a new repository instance
func UserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Save saves a user to database.
// Return the saved user or error
func (r *repository) Save(user Users) (Users, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return Users{}, err
	}
	return user, nil
}
