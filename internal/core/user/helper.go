package user

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret = []byte(os.Getenv("JWT_SECRET"))

// generateJWT creates a new JWT token for the given email
func GenerateJWT(email string) (string, error) {
	expirationTime := time.Now().Add(30 * time.Minute) // Token expires in 30 minutes
	claims := &Claims{
		Username: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "your-app-name",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(JWTSecret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
