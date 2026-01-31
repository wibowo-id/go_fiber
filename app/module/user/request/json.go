package request

import (
	"go_fiber_wibowo/app/database/schema"
	"go_fiber_wibowo/utils/crypto"
	"time"

	"github.com/google/uuid"
)

type UserRequest struct {
	Id          crypto.EncryptedUUID `json:"id"`
	NamaLengkap string               `json:"nama_lengkap"`
	Email       string               `json:"email"`
	Telepon     string               `json:"telepon"`
	Password    string               `json:"password"`
	RoleId      crypto.EncryptedUUID `json:"role_id"`
	IsActive    bool                 `json:"is_active"`
	CreatedAt   time.Time            `json:"created_at"`
	CreatedBy   uuid.UUID            `json:"created_by"`
	UpdatedAt   time.Time            `json:"updated_at"`
	UpdatedBy   uuid.UUID            `json:"updated_by"`
	DeletedAt   time.Time            `json:"deleted_at"`
	DeletedBy   uuid.UUID            `json:"deleted_by"`
}

func (req *UserRequest) ToDomain() *schema.User {
	return &schema.User{
		Id:          req.Id.UUID,
		NamaLengkap: req.NamaLengkap,
		Email:       req.Email,
		Telepon:     req.Telepon,
		Password:    req.Password,
		RoleId:      req.RoleId.UUID,
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
		RoleId:      req.RoleId.UUID,
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
