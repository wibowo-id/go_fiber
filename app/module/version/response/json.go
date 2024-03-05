package response

import (
	"github.com/google/uuid"
	"go_fiber_wibowo/app/database/schema"
)

type Version struct {
	Id         uuid.UUID `json:"id"`
	Version    string    `json:"version"`
	MinVersion string    `json:"min_version"`
	Url        string    `json:"url"`
}

func FromDomain(version *schema.Version) (res *Version) {
	if version != nil {
		res = &Version{
			Id:         version.Id,
			Version:    version.Version,
			MinVersion: version.MinVersion,
			Url:        version.Url,
		}
	}

	return res
}
