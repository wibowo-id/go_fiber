package response

import (
	"github.com/google/uuid"
	"time"
)

type UserResponse struct {
	Id          uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key"`
	NamaLengkap string    `json:"nama_lengkap"`
	Email       string    `json:"email"`
	Telepon     string    `json:"telepon"`
	Password    string    `json:"password"`
	RoleId      uuid.UUID `json:"role_id"`
	NamaRole    string    `json:"nama_role"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   uuid.UUID `json:"created_by"`
}

func FromDomain(req *UserResponse) (res *UserResponse) {
	if req != nil {
		res = &UserResponse{
			Id:          req.Id,
			NamaLengkap: req.NamaLengkap,
			Email:       req.Email,
			Telepon:     req.Telepon,
			Password:    req.Password,
			RoleId:      req.RoleId,
			NamaRole:    req.NamaRole,
			IsActive:    req.IsActive,
			CreatedAt:   req.CreatedAt,
			CreatedBy:   req.CreatedBy,
		}
	}

	return res
}
