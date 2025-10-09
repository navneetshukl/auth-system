package user

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret = []byte(os.Getenv("JWT_SECRET"))

// generateJWT creates a new JWT token for the given email
func GenerateJWT(email string, ttl int) (string, error) {
	expirationTime := time.Now().Add(time.Duration(ttl) * time.Minute) // Token expires in 30 minutes
	claims := &Claims{
		Username: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "auth-service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(JWTSecret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ExtractEmailFromToken(tokenString string) (string, error) {
	// Parse the token with the same Claims structure used for generation
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})

	if err != nil {
		return "", err
	}

	// Validate token and extract the email
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return "", ErrInvalidToken
	}

	return claims.Username, nil
}
