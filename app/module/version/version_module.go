package version

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go_fiber_wibowo/app/module/version/controller"
	"go_fiber_wibowo/app/module/version/repository"
	"go_fiber_wibowo/app/module/version/service"
)

// VersionRouter struct of VersionRouter
type VersionRouter struct {
	App        fiber.Router
	Controller *controller.Controller
}

// register bulky of version module
var NewVersionModule = fx.Options(
	// register repository of version module
	fx.Provide(repository.NewVersionRepository),

	// register service of version module
	fx.Provide(service.NewVersionService),

	// register controller of version module
	fx.Provide(controller.NewController),

	// register router of version module
	fx.Provide(NewVersionRouter),
)

// init VersionRouter
func NewVersionRouter(fiber *fiber.App, controller *controller.Controller) *VersionRouter {
	return &VersionRouter{
		App:        fiber,
		Controller: controller,
	}
}

// register routes of version module
func (_i *VersionRouter) RegisterVersionRoutes(prefix fiber.Router) {
	// define controllers
	versionController := _i.Controller.Version

	// define routes
	prefix.Route("/versions", func(router fiber.Router) {
		router.Get("/", versionController.Index)
		router.Get("/:id", versionController.Show)
		router.Post("/", versionController.Store)
		router.Put("/:id", versionController.Update)
		router.Delete("/:id", versionController.Delete)
	})
}
