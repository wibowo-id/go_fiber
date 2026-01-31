package schema

import "github.com/google/uuid"

type Version struct {
	Id         uuid.UUID `gorm:"primary_key" json:"id"`
	Version    string    `gorm:"type:varchar(20);not null" json:"version"`
	MinVersion string    `gorm:"type:varchar(20);not null" json:"min_version"`
	Url        string    `gorm:"type:varchar(255);not null" json:"url"`
	Base
}
