package modules

import (
	"github.com/GameChangerCorp/cari-kkn-be/api"
	adminApi "github.com/GameChangerCorp/cari-kkn-be/api/admin"
	userApi "github.com/GameChangerCorp/cari-kkn-be/api/user"
	adminBusiness "github.com/GameChangerCorp/cari-kkn-be/business/admin"
	userBusiness "github.com/GameChangerCorp/cari-kkn-be/business/user"
	"github.com/GameChangerCorp/cari-kkn-be/config"
	adminRepo "github.com/GameChangerCorp/cari-kkn-be/repository/admin"
	userRepo "github.com/GameChangerCorp/cari-kkn-be/repository/user"
	"github.com/GameChangerCorp/cari-kkn-be/utils"
)

func RegistrationModules(dbCon *utils.DatabaseConnection, _ *config.AppConfig) api.Controller {
	adminPermitRepository := adminRepo.RepositoryFactory(dbCon)
	adminPermitService := adminBusiness.NewService(adminPermitRepository)
	adminPermitController := adminApi.NewController(adminPermitService)

	usersRepo := userRepo.RepositoryFactory(dbCon)
	usersService := userBusiness.NewService(usersRepo)
	usersController := userApi.NewController(usersService)

	controller := api.Controller{
		AdminController: adminPermitController,
		UserController:  usersController,
	}
	return controller
}
