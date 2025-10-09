package db

import (
	"auth-system/internal/core/user"
	"context"
	"database/sql"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryImpl {
	return UserRepositoryImpl{
		db: db,
	}
}

func (u *UserRepositoryImpl) RegisterUser(ctx context.Context, user *user.User) error {
	sql := `insert into users (name,email,password,mobile) values($1,$2,$3,$4);`
	_, err := u.db.ExecContext(ctx, sql, user.Name, user.Email, user.Password, user.Mobile)
	if err != nil {
		return err
	}

	return nil
}
func (u *UserRepositoryImpl) FindUserByEmail(ctx context.Context, email string) (*user.User, error) {
	query := `SELECT  name, email, password, mobile FROM users WHERE email = $1;`
	var usr user.User
	err := u.db.QueryRowContext(ctx, query, email).Scan(
		&usr.Name,
		&usr.Email,
		&usr.Mobile,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			// No user found with the given email
			return nil, sql.ErrNoRows
		}
		// Other database errors
		return nil, err
	}
	return &usr, nil
}

func (u *UserRepositoryImpl) UpdateUser(ctx context.Context, email string) error {
	query := `UPDATE users SET is_verified = TRUE WHERE email = $1;`

	result, err := u.db.ExecContext(ctx, query, email)
	if err != nil {
		return err 
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
