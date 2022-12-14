package api

import (
	"github.com/GameChangerCorp/cari-kkn-be/api/admin"
	"github.com/GameChangerCorp/cari-kkn-be/api/user"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	AdminController *admin.Controller
	UserController  *user.Controller
}

func RegistrationPath(e *gin.Engine, controller Controller) {
	e.GET("/admin/desa", controller.AdminController.GetAllDesa)
	e.POST("/admin/desa", controller.AdminController.CreateDesa)
	e.GET("/admin/desa/:id", controller.AdminController.GetDesaById)
	e.POST("/admin/login", controller.AdminController.LoginAuth)
	e.POST("admin/register", controller.AdminController.RegisterAdmin)
	e.POST("/admin/desa/:status/:id", controller.AdminController.ApproveRequestDesa)
	e.PUT("/admin/desa/:id", controller.AdminController.UpdateDesa)
	e.DELETE("/admin/desa/:id", controller.AdminController.DeleteDesa)

	e.POST("user/login", controller.UserController.LoginAuth)
	e.POST("user/register", controller.UserController.RegisterAuth)
	e.GET("user/desa", controller.UserController.GetAllData)
	e.POST("user/reservation", controller.UserController.Reservation)
	e.GET("user/reservation/:id", controller.UserController.GetReservation)
}
