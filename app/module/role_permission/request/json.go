package request

import (
	"go_fiber_wibowo/utils/crypto"
	"time"

	"github.com/google/uuid"
)

type CreateRequest struct {
	Id           crypto.EncryptedUUID `json:"id"`
	RoleId       crypto.EncryptedUUID `json:"role_id"`
	PermissionId crypto.EncryptedUUID `json:"permission_id"`
	OrderNumber  int                  `json:"order_number"`
	IsActive     bool                 `json:"is_active"`
	CreatedAt    time.Time            `json:"created_at"`
	CreatedBy    uuid.UUID            `json:"created_by"`
	UpdatedAt    time.Time            `json:"updated_at"`
	UpdatedBy    uuid.UUID            `json:"updated_by"`
}
