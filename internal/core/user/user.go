package user

import (
	"context"
	"time"
)

type User struct {
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Mobile    string     `json:"mobile"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UserService interface {
	RegisterUser(ctx context.Context, data *User) error
}
