package controller

import "go_fiber_wibowo/app/module/user/service"

type Controller struct {
	User UserController
}

func NewController(userService service.UserService) *Controller {
	return &Controller{
		User: NewUserController(userService),
	}
}
