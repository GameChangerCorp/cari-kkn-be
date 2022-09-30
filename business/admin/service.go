package admin

import (
	"errors"

	"github.com/GameChangerCorp/cari-kkn-be/utils"
	"github.com/go-playground/validator/v10"
)

type Repository interface {
	FindAdminByUsername(username string) (*Admin, error)
}
type Service interface {
	LoginAuth(auth AuthLogin) (*ResponseLogin, error)
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

func (s *service) LoginAuth(auth AuthLogin) (*ResponseLogin, error) {
	err := s.validate.Struct(auth)
	if err != nil {
		return nil, err
	}
	data, err := s.repository.FindAdminByUsername(auth.Username)
	if err != nil {
		return nil, errors.New("wrong username")
	}
	err = utils.VerifyPassword(data.Password, auth.Password)
	if err != nil {
		return nil, errors.New("wrong password")
	}
	token, expired, err := utils.CreateToken(data.Username, data.ID, data.Roles.Name)

	res := &ResponseLogin{
		ID:       data.ID,
		Username: data.Username,
		Expired:  *expired,
		Token:    *token,
	}
	return res, nil
}
