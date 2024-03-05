package controller

import (
	"dlh_oss_be/app/module/auth/service"
	role_permission_service "dlh_oss_be/app/module/role_permission/service"
	user_repo "dlh_oss_be/app/module/user/repository"
	wa_service "dlh_oss_be/app/module/wa_sender/service"
)

type Controller struct {
	Auth AuthController
}

func NewController(authService service.AuthService, rolePermissionService role_permission_service.RolePermissionService, waSenderService wa_service.WaSenderService, userRepo user_repo.UserRepository) *Controller {
	return &Controller{
		Auth: NewAuthController(authService, rolePermissionService, waSenderService, userRepo),
	}
}
