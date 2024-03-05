package controller

import "dlh_oss_be/app/module/role/service"

type Controller struct {
	Role RoleController
}

func NewController(badanUsahaService service.RoleService) *Controller {
	return &Controller{
		Role: NewRoleController(badanUsahaService),
	}
}
