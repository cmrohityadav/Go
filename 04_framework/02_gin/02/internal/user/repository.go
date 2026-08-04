package user

import (
	"context"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	Pool *pgxpool.Pool
}

var ErrUserAlreadyExists = errors.New("user already exists")

func NewRepository(pool *pgxpool.Pool)*Repository{
	return &Repository{Pool: pool}
}

func(r *Repository)FindByEmail(ctx context.Context,email string)(User,error){
	emailLower:=strings.ToLower(email);
	row:=r.Pool.QueryRow(
		ctx,
		`SELECT id, email, passwordhash, role FROM users WHERE email=$1`,
		emailLower,
	)

	var user User;

	err:=row.Scan(&user.Id,&user.Email,&user.PasswordHash,&user.Role)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return User{}, pgx.ErrNoRows
		}
		return User{}, err
	}

	return  user,nil

}
func (r *Repository) Create(ctx context.Context,user *User)error{


	err := r.Pool.QueryRow(
		ctx,
		`INSERT INTO users (email, passwordHash, role)
		VALUES ($1, $2, $3)
		RETURNING id`,
		user.Email,
		user.PasswordHash,
		user.Role,
	).Scan(&user.Id)

	if err != nil {
		return err
	}

	return nil;
}

