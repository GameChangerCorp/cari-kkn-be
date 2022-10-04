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
	e.POST("/admin/login", controller.AdminController.LoginAuth)
	e.POST("admin/register", controller.AdminController.RegisterAdmin)

	e.POST("/user/register", controller.UserController.RegisterAdmin)
}
