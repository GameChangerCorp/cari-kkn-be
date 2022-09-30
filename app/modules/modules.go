package modules

import (
	"github.com/GameChangerCorp/cari-kkn-be/api"
	adminApi "github.com/GameChangerCorp/cari-kkn-be/api/admin"
	adminBusiness "github.com/GameChangerCorp/cari-kkn-be/business/admin"
	"github.com/GameChangerCorp/cari-kkn-be/config"
	adminRepo "github.com/GameChangerCorp/cari-kkn-be/repository/admin"
	"github.com/GameChangerCorp/cari-kkn-be/utils"
)

func RegistrationModules(dbCon *utils.DatabaseConnection, _ *config.AppConfig) api.Controller {
	adminPermitRepository := adminRepo.RepositoryFactory(dbCon)
	adminPermitService := adminBusiness.NewService(adminPermitRepository)
	adminPermitController := adminApi.NewController(adminPermitService)
	controller := api.Controller{
		AdminController: adminPermitController,
	}
	return controller
}
