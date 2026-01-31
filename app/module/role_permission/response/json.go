package response

import (
	"github.com/google/uuid"
)

type CreateResponse struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

type MenuPermission struct {
	Id       uuid.UUID         `json:"id"`
	ParentId uuid.UUID         `json:"parent_id"`
	Name     string            `json:"name"`
	Path     string            `json:"path"`
	Icon     string            `json:"icon"`
	Items    []*MenuPermission `json:"items" gorm:"foreignKey:ParentId;references:Id"`
}
