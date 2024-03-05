package schema

import "github.com/google/uuid"

type Role struct {
	Id       uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key"`
	Role     string    `json:"role" gorm:"type:varchar(255)"`
	Beranda  string    `json:"beranda" gorm:"type:varchar(255)"`
	IsActive bool      `json:"is_active"`
	Base
}
