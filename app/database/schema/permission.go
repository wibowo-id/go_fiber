package schema

import "github.com/google/uuid"

type Permission struct {
	Id        uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key"`
	ParentId  uuid.UUID `json:"parent_id" gorm:"type:varchar(36)"`
	Judul     string    `json:"judul" gorm:"type:varchar(255)"`
	Url       string    `json:"url" gorm:"type:varchar(255)"`
	Icon      string    `json:"icon" gorm:"type:varchar(255)"`
	Deskripsi string    `json:"deskripsi" gorm:"type:text"`
	IsActive  bool      `json:"is_active"`
	Base
}
