package users

// RegisterUserInput catches input from signup endpoint.
type RegisterUserInput struct {
	FullName   string `json:"full_name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

// LoginUserInput catches input from login endpoint
type LoginUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
