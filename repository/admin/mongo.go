package admin

import (
	"context"
	"errors"
	"fmt"

	"github.com/GameChangerCorp/cari-kkn-be/business/admin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBRepository struct {
	colAdmin       *mongo.Collection
	colDesa        *mongo.Collection
	colReservation *mongo.Collection
}

func NewMongoRepository(col *mongo.Database) *MongoDBRepository {
	return &MongoDBRepository{
		colAdmin:       col.Collection("admin"),
		colDesa:        col.Collection("desa-kkn"),
		colReservation: col.Collection("reservation"),
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

func (repo *MongoDBRepository) CreateAdmin(auth admin.RegAdmin) error {
	_, err := repo.colAdmin.InsertOne(context.Background(), auth)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MongoDBRepository) FindAllDesa() ([]admin.DesaKKN, error) {
	var data []admin.DesaKKN
	cursor, err := repo.colDesa.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &data)
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	return data, nil
}

func (repo *MongoDBRepository) GetDesaById(id string) (*admin.DesaKKN, error) {
	var data admin.DesaKKN
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}
	err = repo.colDesa.FindOne(context.Background(), bson.M{"_id": objId}).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (repo *MongoDBRepository) CreateDesa(desa admin.CreateDesaKKN) error {
	_, err := repo.colDesa.InsertOne(context.Background(), desa)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MongoDBRepository) ApproveRequestDesa(id, status string) error {
	objId, _ := primitive.ObjectIDFromHex(id)
	fmt.Println(objId)
	err := repo.colReservation.FindOneAndUpdate(context.Background(), bson.M{"_id": objId, "status": "ON PROCESS"}, bson.M{"$set": bson.M{"status": status}}).Err()
	if err != nil {
		fmt.Println(err)
		return errors.New("wrong id desa")
	}
	return nil
}

func (repo *MongoDBRepository) UpdateDesa(id string, input admin.UpdateDesaKKN) error {
	objId, _ := primitive.ObjectIDFromHex(id)
	err := repo.colDesa.FindOneAndUpdate(context.Background(), bson.M{"_id": objId}, input).Err()
	if err != nil {
		return err
	}

	return nil
}

func (repo *MongoDBRepository) DeleteDesa(id string) error {
	objId, _ := primitive.ObjectIDFromHex(id)
	err := repo.colDesa.FindOneAndDelete(context.Background(), bson.M{"_id": objId}).Err()
	if err != nil {
		return err
	}
	return nil
}
