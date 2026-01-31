package options

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go_fiber_wibowo/app/module/options/controller"
	"go_fiber_wibowo/app/module/options/repository"
	"go_fiber_wibowo/app/module/options/service"
)

// struct of OptionRouter
type OptionRouter struct {
	App        fiber.Router
	Controller *controller.Controller
}

// register bulky of options module
var NewOptionModule = fx.Options(
	// register repository of options module
	fx.Provide(repository.NewOptionRepository),

	// register service of options module
	fx.Provide(service.NewOptionService),

	// register controller of options module
	fx.Provide(controller.NewController),

	// register router of options module
	fx.Provide(NewOptionRouter),
)

// init OptionRouter
func NewOptionRouter(fiber *fiber.App, controller *controller.Controller) *OptionRouter {
	return &OptionRouter{
		App:        fiber,
		Controller: controller,
	}
}

// register routes of options module
func (_i *OptionRouter) RegisterOptionRoutes(prefix fiber.Router) {
	// define controllers
	Controller := _i.Controller.Option

	// define routes
	prefix.Route("/options", func(router fiber.Router) {
		router.Get("/wilayah", Controller.GetWilayah)
		router.Get("/role", Controller.GetRole)
		router.Get("/permission", Controller.GetPermission)
	})
}
