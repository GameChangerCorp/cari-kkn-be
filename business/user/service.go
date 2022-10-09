package user

import "github.com/go-playground/validator/v10"

type Repository interface {
	FindAccount(email string) (*User, error)
	CheckAccount(email string) bool
	CreateAccount(user RegUser) error
	FetchAllDesa() ([]Desa, error)
	InsertReservation(userId, villageId string) error
	FetchReservation(userId string) ([]DataReservation, error)
}

type Service interface {
	Login(auth AuthLogin) (*ResponseLogin, error)
	Register(auth RegUser) error
	GetAllDesa() ([]Desa, error)
	Reservation(userId, villageId string) error
	GetReservation(userId string) ([]DataReservation, error)
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
