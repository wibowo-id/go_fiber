package controller

import "go_fiber_wibowo/app/module/options/service"
import auth_service "go_fiber_wibowo/app/module/auth/service"

type Controller struct {
	Option OptionController
}

func NewController(optionService service.OptionService, authService auth_service.AuthService) *Controller {
	return &Controller{
		Option: NewOptionsController(optionService, authService),
	}
}
