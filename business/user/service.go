package user

import "github.com/go-playground/validator/v10"

type Repository interface {
	FindAccount(email string) (*User, error)
	CreateAccount(user RegUser) error
}

type Service interface {
	Login(auth AuthLogin) (*ResponseLogin, error)
	Register(auth RegUser) error
}

type service struct {
	Repository Repository
	Validate   *validator.Validate
}

func NewService(repository Repository) Service {
	return &service{
		Repository: repository,
		Validate:   validator.New(),
	}
}
