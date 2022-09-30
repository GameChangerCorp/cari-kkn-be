package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AdminClaims struct {
	Username string
	Email    string
	Role     string
	jwt.StandardClaims
}

func CreateToken(username string, id primitive.ObjectID, role string) (*string, *int, error) {
	TimeSec := 7200
	expirationTime := time.Now().Add(time.Duration(TimeSec) * time.Second)

	claims := &AdminClaims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	SECRET_KEY := os.Getenv("SECRET_JWT")
	token_jwt, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return nil, nil, err
	}
	return &token_jwt, &TimeSec, err
}
