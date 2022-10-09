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

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (controller *Controller) LoginAuth(c *gin.Context) {
	var auth userBusiness.AuthLogin
	err := c.ShouldBindJSON(&auth)

	if err != nil {
		c.JSON(400, response{Message: err.Error(), Data: nil})
		return
	}
	res, err := controller.service.Login(auth)
	if err != nil {
		c.JSON(400, response{Message: err.Error(), Data: nil})
		return
	}
	c.JSON(200, res)
}

func (controller *Controller) RegisterAuth(c *gin.Context) {
	var auth userBusiness.RegUser
	err := c.ShouldBindJSON(&auth)
	if err != nil {
		c.JSON(400, response{Message: err.Error(), Data: nil})
		return
	}
	err = controller.service.Register(auth)
	if err != nil {
		c.JSON(400, response{Message: err.Error(), Data: nil})
		return
	}
	c.JSON(201, response{Message: "success", Data: nil})
}

func (controller *Controller) GetAllData(c *gin.Context) {
	data, err := controller.service.GetAllDesa()

	if err != nil {
		c.JSON(400, response{Message: err.Error(), Data: nil})
		return
	}

	c.JSON(200, response{Message: "data found", Data: data})
}

func (controller *Controller) Reservation(c *gin.Context) {
	var dataReservation userBusiness.ReqReservation

	err := c.ShouldBindJSON(&dataReservation)

	if err != nil {
		c.JSON(400, response{Message: err.Error(), Data: nil})
		return
	}

	err = controller.service.Reservation(dataReservation.UserId, dataReservation.VillageId)

	if err != nil {
		c.JSON(500, response{Message: err.Error(), Data: nil})
		return
	}

	c.JSON(201, response{Message: "success", Data: nil})
}

func (controller *Controller) GetReservation(c *gin.Context) {
	userId := c.Param("id")

	if userId == "" {
		c.JSON(400, response{Message: "user id can't be empty", Data: nil})
		return
	}

	// userId userBusiness.GetReservation

	// err := c.ShouldBindJSON(&userId)

	data, err := controller.service.GetReservation(userId)

	if err != nil {
		c.JSON(500, response{Message: err.Error(), Data: nil})
		return
	}

	c.JSON(200, response{Message: "success", Data: data})
}
