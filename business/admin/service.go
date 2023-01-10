package admin

import (
	"errors"

	"github.com/GameChangerCorp/cari-kkn-be/repository"
	"github.com/GameChangerCorp/cari-kkn-be/utils"
	"github.com/go-playground/validator/v10"
)

type Repository interface {
	FindAdminByUsername(username string) (*Admin, error)
	CreateAdmin(auth RegAdmin) error
	FindAllDesa() ([]DesaKKN, error)
	GetDesaById(id string) (*DesaKKN, error)
	CreateDesa(desa CreateDesaKKN) error
	ApproveRequestDesa(id, status string) error
	UpdateDesa(id string, desa UpdateDesaKKN) error
	DeleteDesa(id string) error
}
type Service interface {
	LoginAuth(auth AuthLogin) (*ResponseLogin, error)
	RegisterAdmin(auth RegAdmin) error
	GetAllDesa() ([]DesaKKN, error)
	CreateDesa(desa CreateDesaKKN) error
	AcceptRequestDesa(id, status string) error
	GetDesaById(id string) (*DesaKKN, error)
	UpdateDesa(id string, desa UpdateDesaKKN) error
	DeleteDesa(id string) error
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

func (s *service) RegisterAdmin(reg RegAdmin) error {
	err := s.validate.Struct(reg)
	if err != nil {
		return err
	}
	hash, err := utils.Hash(reg.Password)

	reg.Password = string(hash)
	err = s.repository.CreateAdmin(reg)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetAllDesa() ([]DesaKKN, error) {
	res, err := s.repository.FindAllDesa()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *service) CreateDesa(desa CreateDesaKKN) error {
	err := s.validate.Struct(desa)
	if err != nil {
		return err
	}
	err = s.repository.CreateDesa(desa)
	if err != nil {
		return err
	}
	return nil
}

func (s service) AcceptRequestDesa(id, status string) error {
	if id == "" {
		return errors.New("id is empty")
	}
	if status == "approve" {
		status = "APPROVED"
	} else if status == "reject" {
		status = "REJECTED"
	}
	newStatus := repository.SetStatus(status)
	err := s.repository.ApproveRequestDesa(id, newStatus.Status)
	if err != nil {
		return err
	}
	return nil
}

func (s service) GetDesaById(id string) (*DesaKKN, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	res, err := s.repository.GetDesaById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s service) UpdateDesa(id string, input UpdateDesaKKN) error {
	err := s.validate.Struct(input)
	if err != nil {
		return err
	}
	err = s.repository.UpdateDesa(id, input)
	if err != nil {
		return err
	}
	return nil
}

func (s service) DeleteDesa(id string) error {
	if id == "" {
		return errors.New("id is empty")
	}
	err := s.repository.DeleteDesa(id)
	if err != nil {
		return err
	}
	return nil
}
