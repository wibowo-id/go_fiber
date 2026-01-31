package schema

import (
	"errors"
	"github.com/google/uuid"
	"go_fiber_wibowo/utils/helpers"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id            uuid.UUID `gorm:"type:varchar(36);primary_key" json:"id"`
	NamaLengkap   string    `gorm:"type:varchar(255);not null" json:"nama_lengkap" binding:"required"`
	Email         string    `gorm:"type:varchar(255);unique;not null" json:"email" binding:"required, email"`
	Telepon       string    `gorm:"type:varchar(20);not null" json:"telepon" binding:"required"`
	Password      string    `gorm:"type:varchar(255);not null" json:"password" binding:"required"`
	RoleId        uuid.UUID `gorm:"type:varchar(36);not null" json:"role_id" binding:"required"`
	Otp           string    `gorm:"type:varchar(6);null" json:"otp"`
	PathPhoto     string    `gorm:"type:varchar(255);null" json:"path_photo"`
	FirebaseToken string    `gorm:"type:varchar(255);null" json:"firebase_token"`
	RememberToken string    `gorm:"type:varchar(255);null" json:"remember_token"`
	IsActive      bool      `gorm:"true" json:"is_active" binding:"required"`
	Base
}

// ComparePassword compare password
func (u *User) ComparePassword(password string) bool {
	err := helpers.VerifyPassword(password, u.Password)
	if err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return false
	}
	//return helpers.ValidateHash(password, u.Password)
	return true
}
