package schema

import (
	"github.com/google/uuid"
	"time"
)

type ForgotPassword struct {
	Id          uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key"`
	Email       string    `json:"email" gorm:"type:varchar(255)"`
	Key         string    `json:"key" gorm:"type:varchar(255)"`
	ExpiredDate time.Time `json:"expired_date"`
	CreatedBy   uuid.UUID `gorm:"created_by" gorm:"type:varchar(36);not null"`
	UpdatedBy   uuid.UUID `gorm:"created_by" gorm:"type:varchar(36);null"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
