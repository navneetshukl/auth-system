package port

import (
	"auth-system/internal/core/user"
	"context"
)

type UserRepository interface {
	RegisterUser(ctx context.Context, user *user.User) error
	LoginUser(ctx context.Context, email, password string) (*user.User, error)
}
