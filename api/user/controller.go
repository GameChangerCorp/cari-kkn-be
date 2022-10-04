package user

import (
	userBusiness "github.com/GameChangerCorp/cari-kkn-be/business/user"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service userBusiness.Service
}

func NewController(service userBusiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (Controller *Controller) LoginAuth(c *gin.Context) {
	var auth userBusiness.AuthLogin
	err := c.ShouldBindJSON(&auth)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	res, err := Controller.service.Login(auth)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
}

func (Controller *Controller) RegisterAuth(c *gin.Context) {
	var auth userBusiness.RegUser
	err := c.ShouldBindJSON(&auth)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = Controller.service.Register(auth)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}
