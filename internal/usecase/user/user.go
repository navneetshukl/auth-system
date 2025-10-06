package user

import (
	"auth-system/internal/adapter/port"
	userSvc "auth-system/internal/core/user"
	"auth-system/pkg/utils"
	"context"
	"database/sql"
)

type UserServiceImpl struct {
	userRepo port.UserRepository
}

func NewUserService(userRepo port.UserRepository) userSvc.UserService {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}
func (u *UserServiceImpl) RegisterUser(ctx context.Context, data *userSvc.User) error {
	if data.Email == "" {
		return userSvc.ErrInvalidEmail
	}
	if data.Password == "" {
		return userSvc.ErrInvalidPassword
	}
	if data.Name == "" {
		return userSvc.ErrEmptyName
	}

	// hash the password
	hashedPassword, err := utils.HashPassword(data.Password)
	if err != nil {
		return userSvc.ErrSomethingWentWrong
	}
	user := &userSvc.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: hashedPassword,
		Mobile:   data.Mobile,
	}

	err = u.userRepo.RegisterUser(ctx, user)
	if err != nil {
		return userSvc.ErrInsertingUser
	}
	return nil
}
func (u *UserServiceImpl) LoginUser(ctx context.Context, email, password string) (*userSvc.User, string, error) {

	user, err := u.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, "", userSvc.ErrInvalidCredentials
		}
		return nil, "", userSvc.ErrSomethingWentWrong
	}

	if user == nil {
		return nil, "", userSvc.ErrInvalidCredentials
	}
	err = utils.ComparePassword(user.Password, password)
	if err != nil {
		return nil, "", userSvc.ErrInvalidCredentials
	}

	token, err := userSvc.GenerateJWT(user.Email)
	if err != nil {
		return nil, "", userSvc.ErrSomethingWentWrong
	}
	return user, token, nil

}
