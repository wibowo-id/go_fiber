package controller

import "go_fiber_wibowo/app/module/role_permission/service"

type Controller struct {
	RolePermission RolePermissionController
}

func NewController(rolePermissionService service.RolePermissionService) *Controller {
	return &Controller{
		RolePermission: NewRolePermissionController(rolePermissionService),
	}
}
