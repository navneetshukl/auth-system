package user

import "errors"

var (
	ErrInvalidEmail       error = errors.New("invalid email")
	ErrInvalidPassword    error = errors.New("invallid password")
	ErrWrongPassword      error = errors.New("wrong password")
	ErrEmptyName          error = errors.New("name is empty")
	ErrSomethingWentWrong error = errors.New("something went wrong")
	ErrInsertingUser      error = errors.New("error inserting user to DB")
	ErrInvalidCredentials error=errors.New("invalid credentials")
	ErrInvalidToken error=errors.New("token is invalid")
)
