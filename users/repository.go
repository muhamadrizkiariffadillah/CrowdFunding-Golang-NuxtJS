package users

import "gorm.io/gorm"

// Repository defines method for users
type Repository interface {
	Save(user Users) (Users, error)
	FindByEmail(email string) (Users, error)
	FindById(id int) (Users, error)
	Update(user Users) (Users, error)
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

// FindByEmail finds a user by email
// return user or error
func (r *repository) FindByEmail(email string) (Users, error) {
	var user Users

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return Users{}, err
	}
	return user, nil
}

// FindById FindById finds a user by ID
// return user or error
func (r *repository) FindById(id int) (Users, error) {
	var user Users
	err := r.db.Where("id = ?").Find(&user).Error
	if err != nil {
		return Users{}, err
	}
	return user, nil
}

func (r *repository) Update(user Users) (Users, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return Users{}, err
	}
	return user, nil
}
