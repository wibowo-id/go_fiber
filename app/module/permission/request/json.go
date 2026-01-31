package request

import (
	"go_fiber_wibowo/app/database/schema"
	"go_fiber_wibowo/utils/crypto"
	"time"

	"github.com/google/uuid"
)

type PermissionRequest struct {
	Id        crypto.EncryptedUUID `json:"id"`
	Judul     string               `json:"judul"`
	Url       string               `json:"url"`
	Icon      string               `json:"icon"`
	Deskripsi string               `json:"deskripsi"`
	ParentId  crypto.EncryptedUUID `json:"parent_id"`
	IsActive  bool                 `json:"is_active"`
	CreatedAt time.Time            `json:"created_at"`
	CreatedBy uuid.UUID            `json:"created_by"`
	UpdatedAt time.Time            `json:"updated_at"`
	UpdatedBy uuid.UUID            `json:"updated_by"`
}

func (req *PermissionRequest) ToDomain() *schema.Permission {
	perm := &schema.Permission{
		Id:        req.Id.UUID,
		Judul:     req.Judul,
		Url:       req.Url,
		Icon:      req.Icon,
		Deskripsi: req.Deskripsi,
		IsActive:  req.IsActive,
		Base: schema.Base{
			CreatedBy: req.CreatedBy,
			CreatedAt: req.CreatedAt,
		},
	}
	if req.ParentId.UUID != uuid.Nil {
		perm.ParentId = req.ParentId.UUID
	}
	return perm
}

func (req *PermissionRequest) ToDomainUpdate() *schema.Permission {
	perm := &schema.Permission{
		Judul:     req.Judul,
		Url:       req.Url,
		Icon:      req.Icon,
		Deskripsi: req.Deskripsi,
		IsActive:  req.IsActive,
		Base: schema.Base{
			UpdatedAt: req.UpdatedAt,
			UpdatedBy: req.UpdatedBy,
		},
	}
	if req.ParentId.UUID != uuid.Nil {
		perm.ParentId = req.ParentId.UUID
	} else {
		perm.ParentId = uuid.Nil
	}
	return perm
}

type PermissionDeleteRequest struct {
	IsActive  bool      `json:"is_active"`
	DeletedAt time.Time `json:"deleted_at"`
	DeletedBy uuid.UUID `json:"deleted_by"`
}
