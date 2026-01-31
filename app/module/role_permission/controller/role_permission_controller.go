package controller

import (
	"github.com/gofiber/fiber/v2"
	"go_fiber_wibowo/app/module/role_permission/service"
	"go_fiber_wibowo/utils/response"
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
