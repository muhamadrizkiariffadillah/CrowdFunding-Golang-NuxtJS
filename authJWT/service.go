package authJWT

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GenerateToken(userId int) (string, error)
	ValidateToke(token string)(*jwt.Token,error)
}

type jwtService struct {
}

func NewJwtService() *jwtService {
	return &jwtService{}
}

var SECRET_KEY = []byte("random_token")

func (s *jwtService) GenerateToken(userId int) (string, error) {
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


func (s *jwtService)ValidateToke(jwtToken string)(*jwt.Token,error){

	validateToken,err := jwt.Parse(jwtToken,func(t *jwt.Token) (interface{}, error) {

		_,ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok{

			return nil,errors.New("invalid token")

		}

		return []byte(SECRET_KEY),nil
	})

	if err != nil {

		return &jwt.Token{},nil

	}	

	return validateToken,nil
}
