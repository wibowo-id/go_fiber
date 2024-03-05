package service

import (
	"errors"
	"github.com/google/uuid"
	"go_fiber_wibowo/app/database/schema"
	"go_fiber_wibowo/app/module/role_permission/repository"
	"go_fiber_wibowo/app/module/role_permission/response"
)

// RolePermissionService
type rolePermissionService struct {
	rolePermission repository.RolePermissionRepository
}

// define interface of IRolePermissionService
//
//go:generate mockgen -destination=role_permission_service_mock.go -package=service . RolePermissionService
type RolePermissionService interface {
	GetMenu(roleId uuid.UUID) (res []*response.MenuPermission, err error)
	Index() (res []*schema.RolePermission, err error)
	Create() (res []*schema.RolePermission, err error)
	FindById() (res []*schema.RolePermission, err error)
	Update() (res []*schema.RolePermission, err error)
	Delete() (res []*schema.RolePermission, err error)
}

// init RolePermissionService
func NewRolePermissionService(rolePermission repository.RolePermissionRepository) RolePermissionService {
	return &rolePermissionService{
		rolePermission: rolePermission,
	}
}

func (_i *rolePermissionService) GetMenu(roleId uuid.UUID) (res []*response.MenuPermission, err error) {
	res, err = _i.rolePermission.GetMenu(roleId)
	if err != nil {
		err = errors.New("menu not found")
		return
	}

	return
}

func (_i *rolePermissionService) Index() (res []*schema.RolePermission, err error) {
	res, err = _i.rolePermission.Index()
	if err != nil {
		err = errors.New("user not found")
		return
	}

	return
}

func (_i *rolePermissionService) Create() (res []*schema.RolePermission, err error) {
	res, err = _i.rolePermission.Index()
	if err != nil {
		err = errors.New("user not found")
		return
	}

	return
}

func (_i *rolePermissionService) FindById() (res []*schema.RolePermission, err error) {
	res, err = _i.rolePermission.Index()
	if err != nil {
		err = errors.New("user not found")
		return
	}

	return
}

func (_i *rolePermissionService) Update() (res []*schema.RolePermission, err error) {
	res, err = _i.rolePermission.Index()
	if err != nil {
		err = errors.New("user not found")
		return
	}

	return
}

func (_i *rolePermissionService) Delete() (res []*schema.RolePermission, err error) {
	res, err = _i.rolePermission.Index()
	if err != nil {
		err = errors.New("user not found")
		return
	}

	return
}
