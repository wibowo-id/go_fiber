package controller

import "dlh_oss_be/app/module/user/service"

type Controller struct {
	User UserController
}

func NewController(userService service.UserService) *Controller {
	return &Controller{
		User: NewUserController(userService),
	}
}
