package admin

import (
	"context"

	"github.com/GameChangerCorp/cari-kkn-be/business/admin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBRepository struct {
	colAdmin *mongo.Collection
}

func NewMongoRepository(col *mongo.Database) *MongoDBRepository {
	return &MongoDBRepository{
		colAdmin: col.Collection("admin"),
	}
}

func (repo *MongoDBRepository) FindAdminByUsername(username string) (*admin.Admin, error) {
	var data admin.Admin
	err := repo.colAdmin.FindOne(context.Background(), bson.M{"username": username}).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
