package admin

import "go.mongodb.org/mongo-driver/bson/primitive"

type ResponseLogin struct {
	ID       primitive.ObjectID `json:"id"`
	Username string             `json:"username"`
	Expired  int                `json:"expired"`
	Token    string             `json:"token"`
}
type AuthLogin struct {
	Username string `bson:"username,omitempty" validate:"required"`
	Password string `bson:"password,omitempty" validate:"required"`
}

type Admin struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username,omitempty" binding:"required"`
	Fullname string             `bson:"fullname,omitempty" binding:"required"`
	Password string             `bson:"password,omitempty" binding:"required"`
	Role_id  primitive.ObjectID `bson:"role_id,omitempty" binding:"required" json:"role_id"`
	Roles    Role               `bson:"roles" json:"roles"`
}

type Role struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name        string             `bson:"name,omitempty" binding:"required" json:"name"`
	Label       string             `bson:"label,omitempty" binding:"required" json:"label"`
	Description string             `bson:"description,omitempty" binding:"required" json:"description"`
}
type RegAdmin struct {
	Username string `bson:"username,omitempty" binding:"required"`
	Fullname string `bson:"fullname,omitempty" binding:"required"`
	Password string `bson:"password,omitempty" binding:"required"`
}

type DesaKKN struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
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

type CreateDesaKKN struct {
	UrlPhoto    string   `json:"url_photo" bson:"url_photo,omitempty"`
	NamaTempat  string   `json:"nama_tempat" bson:"nama_tempat,omitempty"`
	NamaKota    string   `json:"nama_kota" bson:"nama_kota,omitempty"`
	JumlahOrang int      `json:"jumlah_orang" bson:"jumlah_orang,omitempty"`
	Fasilitas   []string `json:"fasilitas" bson:"fasilitas,omitempty"`
	Kebutuhan   []string `json:"kebutuhan" bson:"kebutuhan,omitempty"`
	PicJabatan  string   `json:"pic_jabatan" bson:"pic_jabatan,omitempty"`
	PicNama     string   `json:"pic_nama" bson:"pic_nama,omitempty"`
	PicPhone    string   `json:"pic_phone" bson:"pic_phone,omitempty"`
}

type UpdateDesaKKN struct {
	NamaTempat  string   `json:"nama_tempat" bson:"nama_tempat,omitempty"`
	NamaKota    string   `json:"nama_kota" bson:"nama_kota,omitempty"`
	JumlahOrang int      `json:"jumlah_orang" bson:"jumlah_orang,omitempty"`
	Fasilitas   []string `json:"fasilitas" bson:"fasilitas,omitempty"`
	Kebutuhan   []string `json:"kebutuhan" bson:"kebutuhan,omitempty"`
	PicJabatan  string   `json:"pic_jabatan" bson:"pic_jabatan,omitempty"`
	PicNama     string   `json:"pic_nama" bson:"pic_nama,omitempty"`
	PicPhone    string   `json:"pic_phone" bson:"pic_phone,omitempty"`
}
