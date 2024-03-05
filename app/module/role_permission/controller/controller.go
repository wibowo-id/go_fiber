package controller

import "dlh_oss_be/app/module/role_permission/service"

type Controller struct {
	RolePermission RolePermissionController
}

func NewController(rolePermissionService service.RolePermissionService) *Controller {
	return &Controller{
		RolePermission: NewRolePermissionController(rolePermissionService),
	}
}
