package role

import (
	"dlh_oss_be/app/module/role/controller"
	"dlh_oss_be/app/module/role/repository"
	"dlh_oss_be/app/module/role/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

// struct of RoleRouter
type RoleRouter struct {
	App        fiber.Router
	Controller *controller.Controller
}

// register bulky of role module
var NewRoleModule = fx.Options(
	// register repository of role module
	fx.Provide(repository.NewRoleRepository),

	// register service of role module
	fx.Provide(service.NewRoleService),

	// register controller of role module
	fx.Provide(controller.NewController),

	// register router of role module
	fx.Provide(NewRoleRouter),
)

// init RoleRouter
func NewRoleRouter(fiber *fiber.App, controller *controller.Controller) *RoleRouter {
	return &RoleRouter{
		App:        fiber,
		Controller: controller,
	}
}

// register routes of role module
func (_i *RoleRouter) RegisterRoleRoutes(prefix fiber.Router) {
	// define controllers
	Controller := _i.Controller.Role

	// define routes
	prefix.Route("/admin/role", func(router fiber.Router) {
		router.Get("/", Controller.Index)
		router.Post("/", Controller.Store)
		router.Get("/:id", Controller.Show)
		router.Put("/:id", Controller.Update)
		router.Delete("/:id", Controller.Delete)
	})
}
