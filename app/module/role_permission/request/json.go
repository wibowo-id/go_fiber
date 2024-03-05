package request

import (
	"github.com/google/uuid"
	"time"
)

type CreateRequest struct {
	Id           uuid.UUID `json:"id"`
	RoleId       string    `json:"role_id"`
	PermissionId string    `json:"permission_id"`
	OrderNumber  int       `json:"order_number"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	CreatedBy    uuid.UUID `json:"created_by"`
	UpdatedAt    time.Time `json:"updated_at"`
	UpdatedBy    uuid.UUID `json:"updated_by"`
}
