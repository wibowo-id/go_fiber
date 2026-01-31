package role_permission

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go_fiber_wibowo/app/middleware"
	"go_fiber_wibowo/app/module/role_permission/controller"
	"go_fiber_wibowo/app/module/role_permission/repository"
	"go_fiber_wibowo/app/module/role_permission/service"
)

// struct of RolePermissionRouter
type RolePermissionRouter struct {
	App        fiber.Router
	Controller *controller.Controller
}

// register bulky of roleRolePermission module
var NewRolePermissionModule = fx.Options(
	// register repository of roleRolePermission module
	fx.Provide(repository.NewRolePermissionRepository),

	// register service of roleRolePermission module
	fx.Provide(service.NewRolePermissionService),

	// register controller of roleRolePermission module
	fx.Provide(controller.NewController),

	// register router of roleRolePermission module
	fx.Provide(NewRolePermissionRouter),
)

// init RolePermissionRouter
func NewRolePermissionRouter(fiber *fiber.App, controller *controller.Controller) *RolePermissionRouter {
	return &RolePermissionRouter{
		App:        fiber,
		Controller: controller,
	}
}

// register routes of roleRolePermission module
func (_i *RolePermissionRouter) RegisterRolePermissionRoutes(prefix fiber.Router) {
	// define controllers
	Controller := _i.Controller.RolePermission

	// define routes
	prefix.Route("/role-permission", func(router fiber.Router) {
		router.Use(middleware.Protected())
		router.Get("/", Controller.Index)
		router.Post("/", Controller.Index)
		router.Get("/:id", Controller.Index)
		router.Put("/:id", Controller.Index)
		router.Delete("/:id", Controller.Index)
	})
}
