package api

import (
	"github.com/GameChangerCorp/cari-kkn-be/api/admin"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	AdminController *admin.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {

}
