package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type ResponseLogin struct {
	ID      primitive.ObjectID `json:"id"`
	Email   string             `json:"email"`
	Expired int                `json:"expired"`
	Token   string             `json:"token"`
}

type AuthLogin struct {
	Email    string `bson:"email, omitempty" validate:"required"`
	Password string `bson:"password, omitempty" validate:"required"`
}

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Email       string             `bson:"email,omitempty" binding:"required"`
	Password    string             `bson:"password,omitempty" binding:"required"`
	Fullname    string             `bson:"fullname,omitempty" binding:"required"`
	Phone       int                `bson:"phone,omitempty" binding:"required"`
	Universitas string             `bson:"universitas,omitempty" binding:"required"`
	Role_id     primitive.ObjectID `bson:"role_id,omitempty" binding:"required" json:"role_id"`
	Roles       Role               `bson:"roles" json:"roles"`
}

type Role struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name        string             `bson:"name,omitempty" binding:"required" json:"name"`
	Label       string             `bson:"label,omitempty" json:"label"`
	Description string             `bson:"description,omitempty" json:"description"`
}

type RegUser struct {
	Fullname    string             `bson:"fullname,omitempty" binding:"required"`
	Email       string             `bson:"email,omitempty" binding:"required"`
	Password    string             `bson:"password,omitempty" binding:"required"`
	Phone       int                `bson:"phone,omitempty" binding:"required"`
	Universitas string             `bson:"universitas,omitempty" binding:"required"`
	RoleId      primitive.ObjectID `bson:"role_id,omitempty"`
}
