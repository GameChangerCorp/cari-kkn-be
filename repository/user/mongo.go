package user

import (
	"context"
	"errors"

	"github.com/GameChangerCorp/cari-kkn-be/business/user"
	"github.com/GameChangerCorp/cari-kkn-be/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBRepository struct {
	colUser        *mongo.Collection
	colRoleUser    *mongo.Collection
	colDesa        *mongo.Collection
	colReservation *mongo.Collection
}

func NewMongoRepository(col *mongo.Database) *MongoDBRepository {
	return &MongoDBRepository{
		colUser:        col.Collection("user"),
		colRoleUser:    col.Collection("roles_user"),
		colDesa:        col.Collection("desa-kkn"),
		colReservation: col.Collection("reservation"),
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

func (repo *MongoDBRepository) FetchAllDesa() ([]user.Desa, error) {
	var data []user.Desa

	cursor, err := repo.colDesa.Find(context.TODO(), bson.D{})

	if err != nil {
		return nil, err
	}

	err = cursor.All(context.Background(), &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *MongoDBRepository) InsertReservation(userId, villageId string) error {
	newStatus := repository.SetStatus("ON_PROCESS")

	newData := user.Reservation{UserId: userId, VillageId: villageId, Status: newStatus.Status}
	_, err := repo.colReservation.InsertOne(context.Background(), newData)

	if err != nil {
		return err
	}

	return nil
}

func (repo *MongoDBRepository) FetchReservation(userId string) ([]user.DataReservation, error) {
	// newStatus := repository.SetStatus("ON_PROCESS")

	// newData := user.DataReservation{User: }
	// _, err := repo.colReservation.InsertOne(context.Background(), newData)

	qry := []bson.M{
		{
			"$match": bson.M{
				"user_id": userId,
			},
		},
		{
			"$lookup": bson.M{
				// Define the tags collection for the join.
				"from": "user",
				// Specify the variable to use in the pipeline stage.
				// "let": bson.M{
				// 	"tags": "$tags",
				// },
				"pipeline": []bson.M{
					// Select only the relevant tags from the tags collection.
					// Otherwise all the tags are selected.
					{
						"$match": bson.M{
							"$expr": bson.M{
								"$in": []interface{}{
									"$uuid",
									"$$tags",
								},
							},
						},
					},
					// Sort tags by their name field in asc. -1 = desc
					{
						"$sort": bson.M{
							"_id": userId,
						},
					},
				},
				// Use tags as the field name to match struct field.
				"as": "user_id",
			},
		},
	}

	cur, err := repo.colReservation.Aggregate(context.Background(), qry)

	if err != nil {
		return nil, err
	}

	var reservation []user.DataReservation

	if err := cur.All(context.Background(), &reservation); err != nil {
		return nil, err
	}

	if len(reservation) == 0 {
		return nil, errors.New("data not found")
	}

	return reservation, nil
}
