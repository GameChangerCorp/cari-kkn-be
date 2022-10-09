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
	Phone       string             `bson:"phone,omitempty" binding:"required"`
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
	Phone       string             `bson:"phone,omitempty" binding:"required"`
	Universitas string             `bson:"universitas,omitempty" binding:"required"`
	RoleId      primitive.ObjectID `bson:"role_id,omitempty"`
}

type Desa struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UrlPhoto    string             `json:"url_photo" bson:"url_photo,omitempty"`
	NamaTempat  string             `json:"nama_tempat" bson:"nama_tempat,omitempty"`
	NamaKota    string             `json:"nama_kota" bson:"nama_kota,omitempty"`
	JumlahOrang int                `json:"jumlah_orang" bson:"jumlah_orang,omitempty"`
	Fasilitas   []string           `json:"fasilitas" bson:"fasilitas,omitempty"`
	Kebutuhan   []string           `json:"kebutuhan" bson:"kebutuhan,omitempty"`
	PicJabatan  string             `json:"pic_jabatan" bson:"pic_jabatan,omitempty"`
	PicNama     string             `json:"pic_nama" bson:"pic_nama,omitempty"`
	PicPhone    string             `json:"pic_phone" bson:"pic_phone,omitempty"`
}

type Reservation struct {
	Id          string `json:"id" bson:"_id,omitempty"`
	UserId      string `json:"user_id" bson:"user_id,omitempty" binding:"required"`
	VillageId   string `json:"village_id" bson:"village_id,omitempty" binding:"required"`
	NamaTempat  string `json:"nama_tempat" bson:"nama_tempat,omitempty"`
	NamaKota    string `json:"nama_kota" bson:"nama_kota,omitempty"`
	JumlahOrang int    `json:"jumlah_orang" bson:"jumlah_orang,omitempty"`
	Status      string `json:"status" bson:"status,omitempty"`
}

type GetReservation struct {
	UserId string `json:"user_id" bson:"user_id,omitempty" binding:"required"`
}
type DataReservation struct {
	User   User   `json:"user" bson:"user,omitempty"`
	Desa   Desa   `json:"desa" bson:"desa,omitempty"`
	Status string `json:"status" bson:"status,omitempty"`
}
