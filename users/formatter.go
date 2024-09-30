package users

// UserFormatter formats the user data for API responses.
type UserFormatter struct {
	FullName   string `json:"full_name"`  // User's full name.
	Occupation string `json:"occupation"` // User's occupation.
	Email      string `json:"email"`      // User's email address.
	Password   string `json:"password"`   // User's hashed password.
	Token      string `json:"token"`
}

// UserFormatter formats the user data along with a token.
// Takes a Users struct and a token string, and returns a formatted UserFormatter.
func APIUserFormatter(user Users, token string) UserFormatter {
	formatter := UserFormatter{
		FullName:   user.FullName,
		Occupation: user.Occupation,
		Email:      user.Email,
		Password:   user.HashPassword, // Return the hashed password.
		Token:      token,
	}
	return formatter
}
