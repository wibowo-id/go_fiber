package response

import (
	"github.com/google/uuid"
)

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
