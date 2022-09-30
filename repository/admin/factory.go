package admin

import (
	"github.com/GameChangerCorp/cari-kkn-be/business/admin"
	"github.com/GameChangerCorp/cari-kkn-be/utils"
)

func RepositoryFactory(dbCon *utils.DatabaseConnection) admin.Repository {
	adminRepo := NewMongoRepository(dbCon.MongoDB)
	return adminRepo
}
