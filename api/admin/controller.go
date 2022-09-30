package admin

import (
	adminBusiness "github.com/GameChangerCorp/cari-kkn-be/business/admin"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service adminBusiness.Service
}

func NewController(service adminBusiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (Controller *Controller) LoginAuth(c *gin.Context) error {
	var auth adminBusiness.AuthLogin
	err := c.ShouldBindJSON(&auth)
	if err != nil {
		return err
	}
	res, err := Controller.service.LoginAuth(auth)
	if err != nil {
		return err
	}
	c.JSON(200, res)
	return nil
}
