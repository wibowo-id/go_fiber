package schema

import "github.com/google/uuid"

type RolePermission struct {
	Id           uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key"`
	RoleId       uuid.UUID `json:"role_id" gorm:"type:varchar(36)"`
	PermissionId uuid.UUID `json:"permission_id" gorm:"type:varchar(36)"`
	OrderNumber  int       `json:"order_number"`
	IsActive     bool      `json:"is_active"`
	Base
}
