package admin

import (
	adminBusiness "github.com/GameChangerCorp/cari-kkn-be/business/admin"
)

type Controller struct {
	service adminBusiness.Service
}

func NewController(service adminBusiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}
