package permission

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go_fiber_wibowo/app/module/permission/controller"
	"go_fiber_wibowo/app/module/permission/repository"
	"go_fiber_wibowo/app/module/permission/service"
)

// struct of PermissionRouter
type PermissionRouter struct {
	App        fiber.Router
	Controller *controller.Controller
}

// register bulky of permission module
var NewPermissionModule = fx.Options(
	// register repository of permission module
	fx.Provide(repository.NewPermissionRepository),

	// register service of permission module
	fx.Provide(service.NewPermissionService),

	// register controller of permission module
	fx.Provide(controller.NewController),

	// register router of permission module
	fx.Provide(NewPermissionRouter),
)

// init PermissionRouter
func NewPermissionRouter(fiber *fiber.App, controller *controller.Controller) *PermissionRouter {
	return &PermissionRouter{
		App:        fiber,
		Controller: controller,
	}
}

// register routes of permission module
func (_i *PermissionRouter) RegisterPermissionRoutes(prefix fiber.Router) {
	// define controllers
	Controller := _i.Controller.Permission

	// define routes
	prefix.Route("/admin/permission", func(router fiber.Router) {
		router.Get("/", Controller.Index)
		router.Post("/", Controller.Store)
		router.Get("/:id", Controller.Show)
		router.Put("/:id", Controller.Update)
		router.Delete("/:id", Controller.Delete)
	})
}
