package request

import (
	"github.com/google/uuid"
	"time"
)

type LoginRequest struct {
	Email    string `json:"email" example:"john.doe@gmail.com" validate:"required"`
	Password string `json:"password" example:"12345678" validate:"required"`
}

type RegisterRequest struct {
	NamaLengkap     string `json:"nama_lengkap" example:"John Doe" validate:"required,max=255"`
	Telepon         string `json:"telepon" example:"12345678" validate:"required,max=20"`
	Email           string `json:"email" example:"john.doe@gmail.com" validate:"required,email,max=255"`
	Password        string `json:"password" example:"12345678" validate:"required,min=8,max=255"`
	ConfirmPassword string `json:"confirm_password" example:"12345678" validate:"required,min=8,max=255"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" example:"john@gmail.com" validate:"required,max=255"`
}

type ChangePasswordRequest struct {
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type ChangePasswordPublicRequest struct {
	Telepon         string `json:"telepon" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type OtpRequest struct {
	Telepon string `json:"telepon" example:"12345678"`
}

type OtpConfirmRequest struct {
	Telepon string `json:"telepon" example:"081323456789"`
	Otp     string `json:"otp" example:"123456"`
}

type UpdateProfileRequest struct {
	NamaLengkap string `json:"nama_lengkap" example:"John Doe"`
	Telepon     string `json:"telepon"`
	Email       string `json:"email"`
}

type UpdateFirebaseTokenRequest struct {
	FirebaseToken string `json:"firebase_token" example:"firebase_token" validate:"required,max=255"`
}

type PayloadSendMessage struct {
	ApiKey  string `json:"api_key"`
	Sender  string `json:"sender"`
	Number  string `json:"number"`
	Message string `json:"message"`
}

type CreatedAtBy struct {
	CreatedAt time.Time `json:"created_at"`
	CreatedBy uuid.UUID `json:"created_by"`
}

type UpdatedAtBy struct {
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy uuid.UUID `json:"updated_by"`
}

type DeletedAtBy struct {
	DeletedAt time.Time `json:"deleted_at"`
	DeletedBy uuid.UUID `json:"deleted_by"`
}
