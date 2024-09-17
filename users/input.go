package users

// RegisterUserInput catches input
type RegisterUserInput struct {
	FullName   string `json:"full_name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}
