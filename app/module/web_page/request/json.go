package request

import (
	"github.com/google/uuid"
	"go_fiber_wibowo/app/database/schema"
	"time"
)

type WebPageRequest struct {
	Id        uuid.UUID `json:"id"`
	Judul     string    `json:"judul"`
	Url       string    `json:"url"`
	Konten    string    `json:"konten"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy uuid.UUID `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy uuid.UUID `json:"updated_by"`
}

func (req *WebPageRequest) ToDomain() *schema.WebPage {
	return &schema.WebPage{
		Id:       req.Id,
		Judul:    req.Judul,
		Url:      req.Url,
		Konten:   req.Konten,
		IsActive: req.IsActive,
		Base: schema.Base{
			CreatedAt: req.CreatedAt,
			CreatedBy: req.CreatedBy,
		},
	}
}

func (req *WebPageRequest) ToDomainUpdate() *schema.WebPage {
	return &schema.WebPage{
		Judul:    req.Judul,
		Url:      req.Url,
		Konten:   req.Konten,
		IsActive: req.IsActive,
		Base: schema.Base{
			UpdatedAt: req.UpdatedAt,
			UpdatedBy: req.UpdatedBy,
		},
	}
}

type WebPageDeleteRequest struct {
	IsActive  bool      `json:"is_active"`
	DeletedAt time.Time `json:"deleted_at"`
	DeletedBy uuid.UUID `json:"deleted_by"`
}
