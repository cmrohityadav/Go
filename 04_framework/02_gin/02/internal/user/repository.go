package user

import (
	"context"
	"errors"

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

func (r *Repository) Create(ctx context.Context,user *User)error{
	row:=r.Pool.QueryRow(
		ctx,
		`SELECT id FROM users WHERE email=$1`,
		user.Email,
	)

	err:=row.Scan(&user.Id)
	//User already exist
	if err==nil{
		return ErrUserAlreadyExists;
	}

	//DB related issue
	if !errors.Is(err,pgx.ErrNoRows){
		return err
	}
	

	err = r.Pool.QueryRow(
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

