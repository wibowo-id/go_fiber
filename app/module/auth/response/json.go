package response

import (
	"dlh_oss_be/app/database/schema"
	"github.com/google/uuid"
)

type BadanUsahaWithName struct {
	Id             uuid.UUID `json:"id"`
	NamaBadanUsaha string    `json:"nama_badan_usaha"`
	Nib            string    `json:"nib"`
	Telepon        string    `json:"telepon"`
	UserId         uuid.UUID `json:"user_id"`
	UserNama       string    `json:"user_nama"`
	ProvinsiId     string    `json:"provinsi_id"`
	ProvinsiNama   string    `json:"provinsi_nama"`
	KotaId         string    `json:"kota_id"`
	KotaNama       string    `json:"kota_nama"`
	KecamatanId    string    `json:"kecamatan_id"`
	KecamatanNama  string    `json:"kecamatan_nama"`
	KelurahanId    string    `json:"kelurahan_id"`
	KelurahanNama  string    `json:"kelurahan_nama"`
	Rt             string    `json:"rt"`
	Rw             string    `json:"rw"`
	KodePos        string    `json:"kode_pos"`
	AlamatLengkap  string    `json:"alamat_lengkap"`
	IsActive       bool      `json:"is_active"`
	schema.Base
}

type UserWithBadanUsaha struct {
	schema.User
	BadanUsaha []*BadanUsahaWithName `json:"badan_usaha" gorm:"foreignKey:CreatedBy;references:Id"`
}

type LoginResponse struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

type RegisterResponse struct {
	Id          uuid.UUID `json:"id"`
	NamaLengkap string    `json:"nama_lengkap"`
	Telepon     string    `json:"telepon"`
	Email       string    `json:"email"`
}

type UpdateFirebaseResponse struct {
	FirebaseToken string `json:"firebase_token"`
}

type UpdateProfileResponse struct {
	NamaLengkap string `json:"nama_lengkap"`
	Telepon     string `json:"telepon"`
	Email       string `json:"email"`
}

type ProfileResponse struct {
	NamaLengkap string `json:"nama_lengkap"`
	Telepon     string `json:"telepon"`
	Email       string `json:"email"`
	PathPhoto   string `json:"path_photo"`
}

type OtpResponse struct {
	Otp int `json:"otp"`
}

type Dashboard struct {
	UserClient int `json:"user_client"`
	UserAdmin  int `json:"user_admin"`
}
