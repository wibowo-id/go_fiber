package controller

import "dlh_oss_be/app/module/options/service"
import auth_service "dlh_oss_be/app/module/auth/service"

type Controller struct {
	Option OptionController
}

func NewController(optionService service.OptionService, authService auth_service.AuthService) *Controller {
	return &Controller{
		Option: NewOptionsController(optionService, authService),
	}
}
