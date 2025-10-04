package port

import (
	"auth-system/internal/core/user"
	"context"
)

type UserRepository interface {
	RegisterUser(ctx context.Context, user *user.User) error
	FindUserByEmail(ctx context.Context, email string) (*user.User, error)
}
