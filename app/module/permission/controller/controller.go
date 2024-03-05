package controller

import "dlh_oss_be/app/module/permission/service"

type Controller struct {
	Permission PermissionController
}

func NewController(permissionService service.PermissionService) *Controller {
	return &Controller{
		Permission: NewPermissionController(permissionService),
	}
}
