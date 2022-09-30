package admin

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBRepository struct {
	col     *mongo.Collection
	colRole *mongo.Collection
	colCoop *mongo.Collection
	colProd *mongo.Collection
}

func NewMongoRepository(col *mongo.Database) *MongoDBRepository {
	return &MongoDBRepository{
		col:     col.Collection("admin"),
		colRole: col.Collection("roles_admin"),
		colCoop: col.Collection("cooperation"),
		colProd: col.Collection("products"),
	}
}
