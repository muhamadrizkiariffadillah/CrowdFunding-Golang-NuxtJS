package users

// RegisterUserInput catches input
type RegisterUserInput struct {
	FullName   string `json:"full_name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}
