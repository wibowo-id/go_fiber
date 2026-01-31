package controller

import "go_fiber_wibowo/app/module/permission/service"

type Controller struct {
	Permission PermissionController
}

func NewController(permissionService service.PermissionService) *Controller {
	return &Controller{
		Permission: NewPermissionController(permissionService),
	}
}
