package schema

import (
	"github.com/google/uuid"
	"time"

	"gorm.io/gorm"
)

type Base struct {
	CreatedBy uuid.UUID      `json:"created_by" gorm:"type:varchar(36);not null"`
	UpdatedBy uuid.UUID      `json:"updated_by" gorm:"type:varchar(36);null"`
	DeletedBy uuid.UUID      `json:"deleted_by" gorm:"type:varchar(36);null"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
