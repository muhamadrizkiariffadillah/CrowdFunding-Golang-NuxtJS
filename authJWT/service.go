package authJWT

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type Service interface {
	GenerateToken(userId int) (string, error)
}

type jwtService struct {
}

func NewJwtService() *jwtService {
	return &jwtService{}
}

var SECRET_KEY = []byte("random_token")

func (s jwtService) GenerateToken(userId int) (string, error) {
	claim := jwt.MapClaims{
		"user_id": userId,
		"created": time.Now(),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
