package request

import (
	"dlh_oss_be/app/database/schema"
	"dlh_oss_be/utils/paginator"
	"github.com/google/uuid"
)

type VersionRequest struct {
	Id         uuid.UUID `json:"id" validate:"required,min=3,max=255"`
	Version    string    `json:"version" validate:"required,min=3"`
	MinVersion string    `json:"min_version" validate:"required,min=3"`
	Url        string    `json:"url" validate:"required,min=3"`
}

type VersionsRequest struct {
	Pagination *paginator.Pagination `json:"pagination"`
}

func (req *VersionRequest) ToDomain() *schema.Version {
	return &schema.Version{
		Id:         req.Id,
		Version:    req.Version,
		MinVersion: req.MinVersion,
		Url:        req.Url,
	}
}
