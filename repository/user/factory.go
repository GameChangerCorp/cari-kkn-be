package user

import (
	"github.com/GameChangerCorp/cari-kkn-be/business/user"
	"github.com/GameChangerCorp/cari-kkn-be/utils"
)

func RepositoryFactory(dbCon *utils.DatabaseConnection) user.Repository {
	userRepo := NewMongoRepository(dbCon.MongoDB)
	return userRepo
}
