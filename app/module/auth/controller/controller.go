package controller

import (
	"go_fiber_wibowo/app/module/auth/service"
	role_permission_service "go_fiber_wibowo/app/module/role_permission/service"
	user_repo "go_fiber_wibowo/app/module/user/repository"
	wa_service "go_fiber_wibowo/app/module/wa_sender/service"
)

type Controller struct {
	Auth AuthController
}

func NewController(authService service.AuthService, rolePermissionService role_permission_service.RolePermissionService, waSenderService wa_service.WaSenderService, userRepo user_repo.UserRepository) *Controller {
	return &Controller{
		Auth: NewAuthController(authService, rolePermissionService, waSenderService, userRepo),
	}
}
