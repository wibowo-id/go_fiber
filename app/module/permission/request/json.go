package request

import (
	"github.com/google/uuid"
	"go_fiber_wibowo/app/database/schema"
	"go_fiber_wibowo/utils/helpers"
	"time"
)

type PermissionRequest struct {
	Id        uuid.UUID `json:"id"`
	Judul     string    `json:"judul"`
	Url       string    `json:"url"`
	Icon      string    `json:"icon"`
	Deskripsi string    `json:"deskripsi"`
	ParentId  string    `json:"parent_id"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy uuid.UUID `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy uuid.UUID `json:"updated_by"`
}

func (req *PermissionRequest) ToDomain() *schema.Permission {
	uuidValid := helpers.IsValidUUID(req.ParentId)
	if uuidValid == true {
		return &schema.Permission{
			Id:        req.Id,
			Judul:     req.Judul,
			Url:       req.Url,
			Icon:      req.Icon,
			Deskripsi: req.Deskripsi,
			ParentId:  uuid.Must(uuid.Parse(req.ParentId)),
			IsActive:  req.IsActive,
			Base: schema.Base{
				CreatedBy: req.CreatedBy,
				CreatedAt: req.CreatedAt,
			},
		}
	} else {
		return &schema.Permission{
			Id:        req.Id,
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
	}
}

func (req *PermissionRequest) ToDomainUpdate() *schema.Permission {
	uuidValid := helpers.IsValidUUID(req.ParentId)
	if uuidValid == true {
		return &schema.Permission{
			Judul:     req.Judul,
			Url:       req.Url,
			Icon:      req.Icon,
			Deskripsi: req.Deskripsi,
			ParentId:  uuid.Must(uuid.Parse(req.ParentId)),
			IsActive:  req.IsActive,
			Base: schema.Base{
				UpdatedAt: req.UpdatedAt,
				UpdatedBy: req.UpdatedBy,
			},
		}
	} else {
		return &schema.Permission{
			Judul:     req.Judul,
			Url:       req.Url,
			Icon:      req.Icon,
			Deskripsi: req.Deskripsi,
			IsActive:  req.IsActive,
			ParentId:  uuid.Nil,
			Base: schema.Base{
				UpdatedAt: req.UpdatedAt,
				UpdatedBy: req.UpdatedBy,
			},
		}
	}
}

type PermissionDeleteRequest struct {
	IsActive  bool      `json:"is_active"`
	DeletedAt time.Time `json:"deleted_at"`
	DeletedBy uuid.UUID `json:"deleted_by"`
}
