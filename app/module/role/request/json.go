package request

import (
	"dlh_oss_be/app/database/schema"
	"github.com/google/uuid"
	"time"
)

type RoleRequest struct {
	Id        uuid.UUID `json:"id"`
	Role      string    `json:"role"`
	Beranda   string    `json:"beranda"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy uuid.UUID `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy uuid.UUID `json:"updated_by"`
}

func (req *RoleRequest) ToDomain() *schema.Role {
	return &schema.Role{
		Id:       req.Id,
		Role:     req.Role,
		Beranda:  req.Beranda,
		IsActive: req.IsActive,
		Base: schema.Base{
			CreatedBy: req.CreatedBy,
			CreatedAt: req.CreatedAt,
		},
	}
}

func (req *RoleRequest) ToDomainUpdate() *schema.Role {
	return &schema.Role{
		Role:     req.Role,
		Beranda:  req.Beranda,
		IsActive: req.IsActive,
		Base: schema.Base{
			UpdatedAt: req.UpdatedAt,
			UpdatedBy: req.UpdatedBy,
		},
	}
}

type RoleDeleteRequest struct {
	IsActive  bool      `json:"is_active"`
	DeletedAt time.Time `json:"deleted_at"`
	DeletedBy uuid.UUID `json:"deleted_by"`
}
