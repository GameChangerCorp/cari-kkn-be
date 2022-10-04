package user

import (
	"errors"
	"fmt"

	"github.com/GameChangerCorp/cari-kkn-be/utils"
)

func (s *service) Login(auth AuthLogin) (*ResponseLogin, error) {
	err := s.Validate.Struct(auth)
	if err != nil {
		return nil, err
	}
	data, err := s.Repository.FindAccount(auth.Email)
	if err != nil {
		return nil, errors.New("account not found")
	}
	err = utils.VerifyPassword(data.Password, auth.Password)
	if err != nil {
		return nil, errors.New("wrong password")
	}
	token, expired, err := utils.CreateToken(data.Email, data.ID, data.Roles.Name)

	if err != nil {
		fmt.Printf("error create token: %v", err)
	}

	res := &ResponseLogin{
		ID:      data.ID,
		Email:   data.Email,
		Expired: *expired,
		Token:   *token,
	}
	return res, nil
}

func (s *service) Register(reg RegUser) error {
	err := s.Validate.Struct(reg)
	if err != nil {
		return err
	}

	hash, err := utils.Hash(reg.Password)

	if err != nil {
		fmt.Printf("error create token %v", err)
	}

	reg.Password = string(hash)
	err = s.Repository.CreateAccount(reg)
	if err != nil {
		return err
	}
	return nil
}
