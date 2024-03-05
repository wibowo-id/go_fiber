package schema

import "github.com/google/uuid"

type WebPage struct {
	Id       uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key"`
	Judul    string    `json:"judul" gorm:"type:varchar(255)"`
	Url      string    `json:"url" gorm:"type:varchar(255)"`
	Konten   string    `json:"konten" gorm:"type:text"`
	IsActive bool      `json:"is_active"`
	Base
}
