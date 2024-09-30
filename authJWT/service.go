package authJWT

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/helper"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/users"
)

type Service interface {
	GenerateToken(userId int) (string, error)
	ValidateToke(token string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewJwtService() *jwtService {
	return &jwtService{}
}

var SECRET_KEY = []byte("random_token")

// GenerateToken creates a new JWT token for a given user ID
func (s *jwtService) GenerateToken(userId int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"created": time.Now().Unix(),
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 24-hour expiration
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// ValidateToke validates the JWT token and returns the parsed token
func (s *jwtService) ValidateToke(jwtToken string) (*jwt.Token, error) {
	// Parse the token and validate it
	validateToken, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		// Ensure the token method is what you expect
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		return SECRET_KEY, nil
	})

	// Handle errors in parsing
	if err != nil {
		return nil, err // Return nil and the error encountered
	}

	// Optionally check if token is valid (e.g., not expired)
	if claims, ok := validateToken.Claims.(jwt.MapClaims); ok && validateToken.Valid {
		// Check for expiration
		if exp, ok := claims["exp"].(float64); ok {
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				return nil, errors.New("token is expired")
			}
		}
	} else {
		return nil, errors.New("invalid token")
	}

	return validateToken, nil
}

func MiddleWare(authService Service, userService users.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		// Check if Authorization header is present and properly formatted
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			response := helper.APIResponse(http.StatusUnauthorized, "failed", "unauthorized", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Extract the token string from the Authorization header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token
		token, err := authService.ValidateToke(tokenString)
		if err != nil || !token.Valid {
			response := helper.APIResponse(http.StatusUnauthorized, "failed", "unauthorized", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Extract claims from the validated token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			response := helper.APIResponse(http.StatusUnauthorized, "failed", "unauthorized", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Retrieve the user ID from claims
		userId, ok := claims["user_id"].(float64)
		if !ok {
			response := helper.APIResponse(http.StatusUnauthorized, "failed", "unauthorized", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Convert userId to int
		user, err := userService.GetUserByID(int(userId))
		if err != nil {
			response := helper.APIResponse(http.StatusUnauthorized, "failed", "unauthorized", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Set the user in the context for further handlers
		c.Set("currentUser", user)
		c.Next()
	}
}
