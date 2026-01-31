package request

import (
	"go_fiber_wibowo/app/database/schema"
	"go_fiber_wibowo/utils/crypto"
	"go_fiber_wibowo/utils/paginator"
)

type VersionRequest struct {
	Id         crypto.EncryptedUUID `json:"id" validate:"required"`
	Version    string               `json:"version" validate:"required,min=3"`
	MinVersion string               `json:"min_version" validate:"required,min=3"`
	Url        string               `json:"url" validate:"required,min=3"`
}

type VersionsRequest struct {
	Pagination *paginator.Pagination `json:"pagination"`
}

func (req *VersionRequest) ToDomain() *schema.Version {
	return &schema.Version{
		Id:         req.Id.UUID,
		Version:    req.Version,
		MinVersion: req.MinVersion,
		Url:        req.Url,
	}
}
