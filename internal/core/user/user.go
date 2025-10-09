package user

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Password   string     `json:"password"`
	Mobile     string     `json:"mobile,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	IsVerified bool       `json:"isVerified,omitempty"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserService interface {
	RegisterUser(ctx context.Context, data *User) error
	LoginUser(ctx context.Context, email, password string) (*User, string, error)
}
