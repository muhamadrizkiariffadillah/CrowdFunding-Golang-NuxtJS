package users

import "time"

// Users entity for database
type Users struct {
	Id             uint
	FullName       string
	Occupation     string
	AvatarFileName string
	Email          string
	HashPassword   string
	Role           string
	Token          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
