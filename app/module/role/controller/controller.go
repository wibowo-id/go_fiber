package controller

import "go_fiber_wibowo/app/module/role/service"

type Controller struct {
	Role RoleController
}

func NewController(roleService service.RoleService) *Controller {
	return &Controller{
		Role: NewRoleController(roleService),
	}
}
