package request

import (
	"github.com/google/uuid"
	"go_fiber_wibowo/app/database/schema"
	"time"
)

type UserRequest struct {
	Id          uuid.UUID `json:"id"`
	NamaLengkap string    `json:"nama_lengkap"`
	Email       string    `json:"email"`
	Telepon     string    `json:"telepon"`
	Password    string    `json:"password"`
	RoleId      uuid.UUID `json:"role_id"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   uuid.UUID `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   uuid.UUID `json:"updated_by"`
	DeletedAt   time.Time `json:"deleted_at"`
	DeletedBy   uuid.UUID `json:"deleted_by"`
}

func (req *UserRequest) ToDomain() *schema.User {
	return &schema.User{
		Id:          req.Id,
		NamaLengkap: req.NamaLengkap,
		Email:       req.Email,
		Telepon:     req.Telepon,
		Password:    req.Password,
		RoleId:      req.RoleId,
		IsActive:    req.IsActive,
		Base: schema.Base{
			CreatedBy: req.CreatedBy,
			CreatedAt: req.CreatedAt,
		},
	}
}

func (req *UserRequest) ToDomainUpdate() *schema.User {
	return &schema.User{
		NamaLengkap: req.NamaLengkap,
		Email:       req.Email,
		Telepon:     req.Telepon,
		Password:    req.Password,
		RoleId:      req.RoleId,
		IsActive:    req.IsActive,
		Base: schema.Base{
			UpdatedAt: req.UpdatedAt,
			UpdatedBy: req.UpdatedBy,
		},
	}
}

type UserDeleteRequest struct {
	IsActive  bool      `json:"is_active"`
	DeletedAt time.Time `json:"deleted_at"`
	DeletedBy uuid.UUID `json:"deleted_by"`
}
