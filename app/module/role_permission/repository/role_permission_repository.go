package repository

import (
	"github.com/google/uuid"
	"go_fiber_wibowo/app/database/schema"
	"go_fiber_wibowo/app/module/role_permission/response"
	"go_fiber_wibowo/internal/bootstrap/database"
	"go_fiber_wibowo/utils/config"
)

type rolePermissionRepository struct {
	DB  *database.Database
	Cfg *config.Config
}

//go:generate mockgen -destination=role_repository_mock.go -package=repository . RolePermissionRepository
type RolePermissionRepository interface {
	GetMenu(roleId uuid.UUID) (user []*response.MenuPermission, err error)
	Index() (user []*schema.RolePermission, err error)
}

func NewRolePermissionRepository(db *database.Database) RolePermissionRepository {
	return &rolePermissionRepository{
		DB: db,
	}
}

func (_i *rolePermissionRepository) GetMenu(roleId uuid.UUID) (res []*response.MenuPermission, err error) {
	qState := "SELECT permissions.id, permissions.judul as name, permissions.url as path, permissions.icon, permissions.parent_id FROM role_permissions "
	qState = qState + "JOIN roles on roles.id = role_permissions.role_id "
	qState = qState + "JOIN permissions on permissions.id = role_permissions.permission_id "
	qState = qState + "WHERE roles.id = '" + roleId.String() + "' "
	qState = qState + "AND permissions.parent_id = '" + uuid.Nil.String() + "' "
	qState = qState + "ORDER BY role_permissions.order_number ASC"

	err = _i.DB.DB.Raw(qState).Scan(&res).Error
	if err != nil {
		return nil, err
	}

	for index, permission := range res {
		var item []*response.MenuPermission

		qStateItem := "SELECT permissions.id, permissions.judul as name, permissions.url as path, permissions.icon, permissions.parent_id FROM role_permissions "
		qStateItem = qStateItem + "JOIN roles on roles.id = role_permissions.role_id "
		qStateItem = qStateItem + "JOIN permissions on permissions.id = role_permissions.permission_id "
		qStateItem = qStateItem + "WHERE roles.id = '" + roleId.String() + "' "
		qStateItem = qStateItem + "AND permissions.parent_id = '" + permission.Id.String() + "' "
		qStateItem = qStateItem + "ORDER BY role_permissions.order_number ASC"

		err = _i.DB.DB.Raw(qStateItem).Scan(&item).Error
		if err != nil {
			return nil, err
		}

		res[index].Items = item
	}

	return res, nil
}

func (_i *rolePermissionRepository) Index() (res []*schema.RolePermission, err error) {
	return
}

func (_i *rolePermissionRepository) Create() (res []*schema.RolePermission, err error) {
	return
}

func (_i *rolePermissionRepository) FindById() (res []*schema.RolePermission, err error) {
	return
}

func (_i *rolePermissionRepository) Update() (res []*schema.RolePermission, err error) {
	return
}

func (_i *rolePermissionRepository) Delete() (res []*schema.RolePermission, err error) {
	return
}
