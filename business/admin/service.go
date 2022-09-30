package admin

import (
	"github.com/go-playground/validator/v10"
)

type Repository interface {
}
type Service interface {
}

type service struct {
	repository Repository
	validate   *validator.Validate
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
		validate:   validator.New(),
	}
}
