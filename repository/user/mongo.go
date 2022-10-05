package user

import (
	"context"

	"github.com/GameChangerCorp/cari-kkn-be/business/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBRepository struct {
	colUser     *mongo.Collection
	colRoleUser *mongo.Collection
}

func NewMongoRepository(col *mongo.Database) *MongoDBRepository {
	return &MongoDBRepository{
		colUser:     col.Collection("user"),
		colRoleUser: col.Collection("roles_user"),
	}
}

func (repo *MongoDBRepository) FindAccount(email string) (*user.User, error) {
	var data user.User
	err := repo.colUser.FindOne(context.Background(), bson.M{"email": email}).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (repo *MongoDBRepository) CreateAccount(auth user.RegUser) error {
	var role user.Role
	err := repo.colRoleUser.FindOne(context.Background(), bson.M{"name": "user"}).Decode(&role)

	if err != nil {
		return err
	}

	auth.RoleId = role.ID
	_, err = repo.colUser.InsertOne(context.Background(), auth)
	if err != nil {
		return err
	}
	return nil
}
