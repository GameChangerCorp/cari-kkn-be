package admin

import (
	"fmt"

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

func (Controller *Controller) LoginAuth(c *gin.Context) {
	var auth adminBusiness.AuthLogin
	err := c.ShouldBindJSON(&auth)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	res, err := Controller.service.LoginAuth(auth)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
	return
}

func (Controller *Controller) RegisterAdmin(c *gin.Context) {
	var auth adminBusiness.RegAdmin
	err := c.ShouldBindJSON(&auth)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = Controller.service.RegisterAdmin(auth)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
	return
}

func (Controller *Controller) GetAllDesa(c *gin.Context) {
	res, err := Controller.service.GetAllDesa()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, map[string]interface{}{
		"code":    200,
		"message": "success",
		"result":  res,
	})
	return
}

func (Controller *Controller) CreateDesa(c *gin.Context) {
	var desa adminBusiness.CreateDesaKKN
	err := c.ShouldBindJSON(&desa)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = Controller.service.CreateDesa(desa)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, map[string]interface{}{
		"code":    200,
		"message": "success create desa",
	})
	return
}

func (Controller *Controller) ApproveRequestDesa(c *gin.Context) {
	id := c.Param("id")
	status := c.Param("status")
	if status != "approve" && status != "reject" {
		c.JSON(400, map[string]interface{}{
			"code":    400,
			"message": "status must be approved or rejected",
		})
		return
	}
	err := Controller.service.AcceptRequestDesa(id, status)
	if err != nil {
		c.JSON(400, map[string]interface{}{
			"code":    400,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, map[string]interface{}{
		"code":    200,
		"message": fmt.Sprintf("success %s desa", status),
	})
	return
}

func (Controller *Controller) GetDesaById(c *gin.Context) {
	id := c.Param("id")
	res, err := Controller.service.GetDesaById(id)
	if err != nil {
		c.JSON(400, map[string]interface{}{
			"code":    400,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, map[string]interface{}{
		"code":    200,
		"message": "success",
		"result":  res,
	})
	return
}
