package user

import (
	"auth/internal/auth"
	"context"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo      *Repository
	jwtSecret string
}

func NewService(repo *Repository, jwtSecret string) *Service {
	return &Service{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (s *Service) CreateUser(ctx context.Context,userCreation UserCreationRequest)(UserCreationResponse,error){
	email:=strings.ToLower(userCreation.Email);
	user,err:=s.repo.FindByEmail(ctx,email)
	if user.Email=="" && err!=pgx.ErrNoRows{
		return UserCreationResponse{},err
	}
	if user.Email!=""{
		return UserCreationResponse{},errors.New("User Already Existed")
	}

	hashedPassword,err:=bcrypt.GenerateFromPassword([]byte(userCreation.Password),bcrypt.DefaultCost)
	if err!=nil{
		return UserCreationResponse{},errors.New("Failed while hashing Password")
	}

	userObj := User{
		Email:        email,
		PasswordHash: string(hashedPassword),
		Role:         "user",
	}
	
	err=s.repo.Create(ctx,&userObj)
	if err!=nil{
		return UserCreationResponse{},err;
	}

	token, err := auth.CreateToken(s.jwtSecret,email,userObj.Role)
	if err != nil {
		return UserCreationResponse{}, err
	}

	return UserCreationResponse{
		Token: token,
		Email: email,
	}, nil

}
