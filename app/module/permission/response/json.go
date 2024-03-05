package response

import "go_fiber_wibowo/app/database/schema"

func FromDomain(role *schema.Permission) (res *schema.Permission) {
	if role != nil {
		res = &schema.Permission{
			Id:        role.Id,
			Judul:     role.Judul,
			Deskripsi: role.Deskripsi,
			Url:       role.Url,
			Icon:      role.Icon,
			IsActive:  role.IsActive,
			Base: schema.Base{
				CreatedAt: role.CreatedAt,
				CreatedBy: role.CreatedBy,
				UpdatedAt: role.UpdatedAt,
				UpdatedBy: role.UpdatedBy,
				DeletedAt: role.DeletedAt,
				DeletedBy: role.DeletedBy,
			},
		}
	}

	return res
}
