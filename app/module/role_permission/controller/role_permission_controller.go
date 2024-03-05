package controller

import (
	"dlh_oss_be/app/module/role_permission/service"
	"dlh_oss_be/utils/response"
	"github.com/gofiber/fiber/v2"
)

type rolePermissionController struct {
	rolePermissionService service.RolePermissionService
}

type RolePermissionController interface {
	Index(c *fiber.Ctx) error
}

func NewRolePermissionController(rolePermissionService service.RolePermissionService) RolePermissionController {
	return &rolePermissionController{
		rolePermissionService: rolePermissionService,
	}
}

func (_i *rolePermissionController) Index(c *fiber.Ctx) error {
	res, err := _i.rolePermissionService.Index()
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Data:     res,
		Messages: response.Messages{"Success"},
		Code:     fiber.StatusOK,
	})
}
