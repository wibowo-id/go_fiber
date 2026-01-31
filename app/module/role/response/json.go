package response

import (
	"github.com/google/uuid"
	"time"
)

type Role struct {
	Id        uuid.UUID `json:"id"`
	Role      string    `json:"role"`
	Beranda   string    `json:"beranda"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy uuid.UUID `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy uuid.UUID `json:"updated_by"`
	DeletedAt time.Time `json:"deleted_at"`
	DeletedBy uuid.UUID `json:"deleted_by"`
}

func FromDomain(role *Role) (res *Role) {
	if role != nil {
		res = &Role{
			Id:        role.Id,
			Role:      role.Role,
			Beranda:   role.Beranda,
			IsActive:  role.IsActive,
			CreatedAt: role.CreatedAt,
			CreatedBy: role.CreatedBy,
			UpdatedAt: role.UpdatedAt,
			UpdatedBy: role.UpdatedBy,
			DeletedAt: role.DeletedAt,
			DeletedBy: role.DeletedBy,
		}
	}

	return res
}
