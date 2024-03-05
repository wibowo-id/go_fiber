package options

import (
	"dlh_oss_be/app/module/options/controller"
	"dlh_oss_be/app/module/options/repository"
	"dlh_oss_be/app/module/options/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

// struct of OptionRouter
type OptionRouter struct {
	App        fiber.Router
	Controller *controller.Controller
}

// register bulky of badanUsaha module
var NewOptionModule = fx.Options(
	// register repository of badanUsaha module
	fx.Provide(repository.NewOptionRepository),

	// register service of badanUsaha module
	fx.Provide(service.NewOptionService),

	// register controller of badanUsaha module
	fx.Provide(controller.NewController),

	// register router of badanUsaha module
	fx.Provide(NewOptionRouter),
)

// init OptionRouter
func NewOptionRouter(fiber *fiber.App, controller *controller.Controller) *OptionRouter {
	return &OptionRouter{
		App:        fiber,
		Controller: controller,
	}
}

// register routes of badanUsaha module
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
