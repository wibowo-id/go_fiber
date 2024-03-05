package response

import (
	"dlh_oss_be/app/database/schema"
	"github.com/google/uuid"
	"time"
)

type WebPageResponse struct {
	Id        uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key"`
	Judul     string    `json:"judul"`
	Url       string    `json:"url"`
	Konten    string    `json:"konten"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy uuid.UUID `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy uuid.UUID `json:"updated_by"`
}

func FromDomain(webPage *schema.WebPage) (res *WebPageResponse) {
	if webPage != nil {
		res = &WebPageResponse{
			Id:        webPage.Id,
			Judul:     webPage.Judul,
			Url:       webPage.Url,
			Konten:    webPage.Konten,
			IsActive:  webPage.IsActive,
			CreatedAt: webPage.CreatedAt,
			CreatedBy: webPage.CreatedBy,
		}
	}

	return res
}

func FromDomainUpdate(webPage *schema.WebPage) (res *WebPageResponse) {
	if webPage != nil {
		res = &WebPageResponse{
			Id:        webPage.Id,
			Judul:     webPage.Judul,
			Url:       webPage.Url,
			Konten:    webPage.Konten,
			IsActive:  webPage.IsActive,
			UpdatedAt: webPage.UpdatedAt,
			UpdatedBy: webPage.UpdatedBy,
		}
	}

	return res
}
